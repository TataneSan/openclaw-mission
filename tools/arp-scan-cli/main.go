// arp-scan-cli views and manages the Linux ARP table.
// It reads /proc/net/arp and supports scanning, watching, and exporting.
package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"sort"
	"strings"
	"syscall"
	"time"
)

var version = "1.0.0"

// ARPEntry represents a single ARP table entry.
type ARPEntry struct {
	IPAddress  string `json:"ip_address"`
	HWType     string `json:"hw_type"`
	Flags      string `json:"flags"`
	HWAddress  string `json:"hw_address"`
	Mask       string `json:"mask"`
	Device     string `json:"device"`
	Vendor     string `json:"vendor"`
	IsComplete bool   `json:"is_complete"`
}

func main() {
	if len(os.Args) < 2 {
		cmdShow(nil)
		return
	}

	switch os.Args[1] {
	case "show", "sh", "list", "ls":
		cmdShow(os.Args[2:])
	case "scan":
		cmdScan(os.Args[2:])
	case "watch":
		cmdWatch(os.Args[2:])
	case "get":
		cmdGet(os.Args[2:])
	case "export":
		cmdExport(os.Args[2:])
	case "add":
		cmdAdd(os.Args[2:])
	case "del", "delete":
		cmdDelete(os.Args[2:])
	case "version":
		fmt.Println("arp-scan-cli", version)
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`arp-scan-cli - Linux ARP table viewer and scanner

Usage:
  arp-scan-cli <command> [arguments]

Commands:
  show, sh [iface]  Show ARP table entries (optionally filtered by interface)
  scan <network>    Scan a network for active hosts via ARP (requires root)
  watch [iface]     Watch ARP table changes in real-time
  get <ip>          Get MAC address for a specific IP
  export [format]   Export ARP table (json, csv, text)
  add <ip> <mac>    Add a static ARP entry (requires root)
  del, delete <ip>  Delete an ARP entry (requires root)
  version           Show version
  help              Show this help message

Examples:
  arp-scan-cli show
  arp-scan-cli show eth0
  arp-scan-cli scan 192.168.1.0/24
  arp-scan-cli watch eth0
  arp-scan-cli get 192.168.1.1
  arp-scan-cli export json
  arp-scan-cli add 192.168.1.100 aa:bb:cc:dd:ee:ff
  arp-scan-cli del 192.168.1.100

Notes:
  - scan, add, and delete require root privileges
  - show, watch, get, and export work without root
  - Reads from /proc/net/arp and ip neigh`)
}

// cmdShow displays the ARP table.
func cmdShow(args []string) {
	iface := ""
	if len(args) > 0 {
		iface = args[0]
	}

	entries := getARPEntries()
	if iface != "" {
		entries = filterByDevice(entries, iface)
	}

	if len(entries) == 0 {
		if iface != "" {
			fmt.Fprintf(os.Stderr, "no ARP entries for interface '%s'\n", iface)
		} else {
			fmt.Println("ARP table is empty")
		}
		return
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].IPAddress < entries[j].IPAddress
	})

	printARPTable(entries)
}

// cmdScan scans a network for active hosts.
func cmdScan(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: arp-scan-cli scan <network>\n")
		fmt.Fprintf(os.Stderr, "example: arp-scan-cli scan 192.168.1.0/24\n")
		os.Exit(1)
	}

	network := args[0]
	_, ipNet, err := net.ParseCIDR(network)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid network: %s\n", network)
		os.Exit(1)
	}

	ones, bits := ipNet.Mask.Size()
	totalHosts := 1 << (bits - ones)

	if totalHosts > 2048 {
		fmt.Fprintf(os.Stderr, "network too large: %d hosts (max 2048)\n", totalHosts)
		os.Exit(1)
	}

	// Get network start
	startIP := ipNet.IP.To4()
	for i := 0; i < 4; i++ {
		startIP[i] &= ipNet.Mask[i]
	}

	fmt.Printf("Scanning %s (%d hosts)...\n", network, totalHosts-2)
	fmt.Println()

	// Use arp-scan if available, otherwise use ping sweep
	if useArpScan(startIP, totalHosts, network) {
		return
	}

	// Fallback: ping sweep
	scanWithPing(startIP, totalHosts, network)
}

// cmdWatch watches ARP table changes.
func cmdWatch(args []string) {
	iface := ""
	interval := 2.0

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-i", "--interval":
			i++
			if i < len(args) {
				fmt.Sscanf(args[i], "%f", &interval)
			}
		default:
			iface = args[i]
		}
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nStopped.")
		os.Exit(0)
	}()

	ticker := time.NewTicker(time.Duration(interval * float64(time.Second)))
	defer ticker.Stop()

	prevEntries := make(map[string]ARPEntry)

	fmt.Printf("Watching ARP table (interval: %.1fs)...\n", interval)
	fmt.Println()

	for range ticker.C {
		entries := getARPEntries()
		if iface != "" {
			entries = filterByDevice(entries, iface)
		}

		currMap := make(map[string]ARPEntry)
		for _, e := range entries {
			currMap[e.IPAddress] = e
		}
		showDiff(prevEntries, currMap)
		prevEntries = currMap
	}
}

// cmdGet gets the MAC address for a specific IP.
func cmdGet(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: arp-scan-cli get <ip>\n")
		os.Exit(1)
	}

	ip := args[0]
	entries := getARPEntries()

	for _, e := range entries {
		if e.IPAddress == ip {
			fmt.Printf("IP Address:  %s\n", e.IPAddress)
			fmt.Printf("MAC Address: %s\n", e.HWAddress)
			fmt.Printf("Interface:   %s\n", e.Device)
			fmt.Printf("Vendor:      %s\n", e.Vendor)
			fmt.Printf("Status:      %s\n", entryStatus(e))
			return
		}
	}

	fmt.Fprintf(os.Stderr, "no ARP entry for %s\n", ip)
	os.Exit(1)
}

// cmdExport exports the ARP table.
func cmdExport(args []string) {
	format := "json"
	if len(args) > 0 {
		format = strings.ToLower(args[0])
	}

	entries := getARPEntries()
	if len(entries) == 0 {
		fmt.Println("ARP table is empty")
		return
	}

	switch format {
	case "json":
		data, err := json.MarshalIndent(entries, "", "  ")
		if err != nil {
			fail("error marshaling JSON: " + err.Error())
		}
		fmt.Println(string(data))
	case "csv":
		write := os.Stdout
		w := csv.NewWriter(write)
		w.Write([]string{"ip_address", "hw_type", "flags", "hw_address", "mask", "device", "vendor"})
		for _, e := range entries {
			w.Write([]string{e.IPAddress, e.HWType, e.Flags, e.HWAddress, e.Mask, e.Device, e.Vendor})
		}
		w.Flush()
	case "text":
		printARPTable(entries)
	default:
		fmt.Fprintf(os.Stderr, "unknown format: %s (use json, csv, or text)\n", format)
		os.Exit(1)
	}
}

// cmdAdd adds a static ARP entry.
func cmdAdd(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: arp-scan-cli add <ip> <mac> [-i <iface>]\n")
		os.Exit(1)
	}

	ip := args[0]
	mac := args[1]
	iface := ""

	i := 2
	for i < len(args) {
		switch args[i] {
		case "-i", "--interface":
			i++
			if i < len(args) {
				iface = args[i]
			}
		}
		i++
	}

	ipArgs := []string{"neigh", "add", ip, "lladdr", mac, "nud", "permanent"}
	if iface != "" {
		ipArgs = append(ipArgs, "dev", iface)
	}

	output, err := exec.Command("ip", ipArgs...).CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error adding ARP entry: %v\n", err)
		if len(output) > 0 {
			fmt.Fprintf(os.Stderr, "%s", output)
		}
		os.Exit(1)
	}
	fmt.Printf("Added static ARP entry: %s -> %s\n", ip, mac)
}

// cmdDelete deletes an ARP entry.
func cmdDelete(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: arp-scan-cli del <ip>\n")
		os.Exit(1)
	}

	ip := args[0]
	output, err := exec.Command("ip", "neigh", "del", ip).CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error deleting ARP entry: %v\n", err)
		if len(output) > 0 {
			fmt.Fprintf(os.Stderr, "%s", output)
		}
		os.Exit(1)
	}
	if len(output) > 0 {
		fmt.Print(string(output))
	}
	fmt.Printf("Deleted ARP entry for %s\n", ip)
}

// getARPEntries reads the ARP table from ip neigh or /proc/net/arp.
func getARPEntries() []ARPEntry {
	if entries := getEntriesFromIP(); len(entries) > 0 {
		return entries
	}
	return getEntriesFromProc()
}

// getEntriesFromIP parses 'ip neigh show' output.
func getEntriesFromIP() []ARPEntry {
	output, err := exec.Command("ip", "neigh", "show").Output()
	if err != nil {
		return nil
	}

	var entries []ARPEntry
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		e := parseIPNeighLine(line)
		entries = append(entries, e)
	}
	return entries
}

// parseIPNeighLine parses a line from 'ip neigh show'.
func parseIPNeighLine(line string) ARPEntry {
	parts := strings.Fields(line)
	e := ARPEntry{}

	for i := 0; i < len(parts); i++ {
		switch parts[i] {
		case "dev":
			if i+1 < len(parts) {
				e.Device = parts[i+1]
				i++
			}
		case "lladdr":
			if i+1 < len(parts) {
				e.HWAddress = parts[i+1]
				e.IsComplete = true
				i++
			}
		case "FAILED", "STALE", "REACHABLE", "DELAY", "PROBE",
			"PERMANENT", "NOARP", "INCOMPLETE", "NONE":
			e.Flags = parts[i]
		default:
			if e.IPAddress == "" && net.ParseIP(parts[i]) != nil {
				e.IPAddress = parts[i]
			}
		}
	}

	if e.HWAddress == "" {
		e.HWAddress = "incomplete"
	}
	e.Vendor = getVendor(e.HWAddress)
	e.HWType = "ether"

	return e
}

// getEntriesFromProc reads /proc/net/arp.
func getEntriesFromProc() []ARPEntry {
	file, err := os.Open("/proc/net/arp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: cannot read ARP table: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var entries []ARPEntry
	scanner := bufio.NewScanner(file)

	// Skip header
	if !scanner.Scan() {
		return entries
	}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 6 {
			continue
		}

		e := ARPEntry{
			IPAddress:  fields[0],
			HWType:     fields[1],
			Flags:      parseFlags(fields[2]),
			HWAddress:  fields[3],
			Mask:       fields[4],
			Device:     fields[5],
			IsComplete: fields[3] != "*",
		}
		e.Vendor = getVendor(e.HWAddress)
		entries = append(entries, e)
	}
	return entries
}

// parseFlags converts hex flags to string.
func parseFlags(hexFlags string) string {
	var flags []string
	val, _ := fmt.Sscanf(hexFlags, "%x", new(uint))
	_ = val
	var f uint
	fmt.Sscanf(hexFlags, "%x", &f)

	if f&(1<<0) != 0 {
		flags = append(flags, "C") // COMPLETE
	}
	if f&(1<<1) != 0 {
		flags = append(flags, "P") // PERMANENT
	}
	if len(flags) == 0 {
		return "?"
	}
	return strings.Join(flags, "")
}

// getVendor returns the vendor name from the OUI (first 3 bytes of MAC).
func getVendor(mac string) string {
	if mac == "" || mac == "incomplete" || mac == "*" {
		return "Unknown"
	}

	oui := strings.ToUpper(strings.Split(mac, ":")[0])
	if v, ok := ouiDB[oui]; ok {
		return v
	}
	return "Unknown"
}

// entryStatus returns a human-readable status for an ARP entry.
func entryStatus(e ARPEntry) string {
	switch {
	case e.Flags == "PERMANENT":
		return "Static"
	case e.IsComplete:
		return "Active"
	case e.Flags == "FAILED":
		return "Failed"
	default:
		return "Incomplete"
	}
}

// filterByDevice filters entries by device/interface.
func filterByDevice(entries []ARPEntry, device string) []ARPEntry {
	var filtered []ARPEntry
	for _, e := range entries {
		if e.Device == device {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

// printARPTable prints entries in a formatted table.
func printARPTable(entries []ARPEntry) {
	fmt.Printf("%-16s %-18s %-12s %-10s %-8s\n", "IP ADDRESS", "MAC ADDRESS", "INTERFACE", "VENDOR", "STATUS")
	fmt.Println(strings.Repeat("-", 70))

	for _, e := range entries {
		status := entryStatus(e)
		fmt.Printf("%-16s %-18s %-12s %-10s %-8s\n",
			e.IPAddress, e.HWAddress, e.Device, e.Vendor, status)
	}
	fmt.Printf("\n%d entry(ies)\n", len(entries))
}

// showDiff shows changes between two ARP snapshots.
func showDiff(prev, curr map[string]ARPEntry) {
	hasChanges := false

	for _, e := range curr {
		if _, ok := prev[e.IPAddress]; !ok {
			fmt.Printf("[+] %s -> %s (%s)\n", e.IPAddress, e.HWAddress, e.Device)
			hasChanges = true
		}
	}

	for ip := range prev {
		if _, ok := curr[ip]; !ok {
			fmt.Printf("[-] %s removed\n", ip)
			hasChanges = true
		}
	}

	if !hasChanges {
		fmt.Printf("[%s] %d entries (no changes)\n", time.Now().Format("15:04:05"), len(curr))
	}
	fmt.Println()
}

// useArpScan tries to use the arp-scan system utility.
func useArpScan(startIP net.IP, totalHosts int, network string) bool {
	_, err := exec.LookPath("arp-scan")
	if err != nil {
		return false
	}

	output, err := exec.Command("arp-scan", "--localnet", "--interface=eth0", "--quiet").CombinedOutput()
	if err != nil {
		return false
	}

	var entries []ARPEntry
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			e := ARPEntry{
				IPAddress:  parts[0],
				HWAddress:  parts[1],
				Device:     "eth0",
				IsComplete: true,
				HWType:     "ether",
			}
			if len(parts) >= 3 {
				e.Vendor = parts[2]
			} else {
				e.Vendor = getVendor(e.HWAddress)
			}
			entries = append(entries, e)
		}
	}

	if len(entries) > 0 {
		printARPTable(entries)
	}
	return true
}

// scanWithPing does a simple ping sweep.
func scanWithPing(startIP net.IP, totalHosts int, network string) {
	active := make(map[string]bool)

	for i := 1; i < totalHosts-1; i++ {
		ip := make(net.IP, 4)
		copy(ip, startIP)
		ip[3] = startIP[3] + byte(i)

		cmd := exec.Command("ping", "-c", "1", "-W", "1", ip.String())
		if err := cmd.Run(); err == nil {
			active[ip.String()] = true
		}
	}

	// Now get ARP entries for active hosts
	entries := getARPEntries()
	var activeEntries []ARPEntry
	for _, e := range entries {
		if active[e.IPAddress] {
			activeEntries = append(activeEntries, e)
		}
	}

	if len(activeEntries) == 0 {
		fmt.Println("No active hosts found")
		return
	}

	fmt.Printf("Found %d active host(s):\n\n", len(activeEntries))
	printARPTable(activeEntries)
}

// Minimal OUI database for common vendors.
var ouiDB = map[string]string{
	"00:00:00": "Reserved",
	"00:0C:29": "VMware",
	"00:1C:14": "VMware",
	"00:50:56": "VMware",
	"00:05:69": "Intel",
	"00:1B:21": "Intel",
	"00:21:5A": "Intel",
	"00:23:54": "D-Link",
	"00:24:8C": "Realtek",
	"00:26:5A": "Realtek",
	"00:E0:4C": "Realtek",
	"00:1A:79": "TP-Link",
	"00:03:7F": "Cisco",
	"00:60:2F": "Cisco",
	"00:1E:C2": "Dell",
	"00:25:90": "Dell",
	"00:26:AB": "Dell",
	"00:24:E8": "Microsoft",
	"00:50:FF": "Microsoft",
	"00:0F:1F": "HP",
	"00:17:C8": "HP",
	"00:21:5C": "HP",
	"00:26:B9": "HP",
	"00:1E:68": "Apple",
	"00:23:69": "Apple",
	"00:26:BB": "Apple",
	"00:1B:44": "ASUS",
	"00:24:B2": "ASUS",
	"00:26:9C": "Netgear",
	"00:03:47": "Linksys",
	"00:23:60": "Linksys",
	"00:0D:67": "Buffalo",
	"00:01:36": "3Com",
	"00:04:AC": "Xerox",
	"00:0D:88": "Brother",
	"00:1B:29": "Brother",
	"00:23:24": "Canon",
	"00:26:5B": "Samsung",
	"00:04:9F": "Synology",
	"00:11:32": "Synology",
	"52:54:00": "QEMU",
	"08:00:27": "VirtualBox",
	"0A:00:27": "VirtualBox",
	"CA:FE:C0": "Npcap",
}

func fail(msg string) {
	fmt.Fprintf(os.Stderr, "error: %s\n", msg)
	os.Exit(1)
}
