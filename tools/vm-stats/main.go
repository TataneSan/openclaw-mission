// vm-stats collects and displays system statistics (CPU, RAM, disk, network).
//
// Supports one-shot display, continuous monitoring, and JSON output.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// SystemStats holds all collected system metrics.
type SystemStats struct {
	Timestamp string            `json:"timestamp"`
	Hostname  string            `json:"hostname"`
	Uptime    string            `json:"uptime"`
	CPU       CPUStats          `json:"cpu"`
	Memory    MemoryStats       `json:"memory"`
	Disk      []DiskStats       `json:"disk"`
	Network   []NetworkStats    `json:"network"`
	LoadAvg   LoadAvgStats      `json:"load_avg"`
	Processes ProcessStats      `json:"processes"`
}

type CPUStats struct {
	User    float64 `json:"user"`
	System  float64 `json:"system"`
	Idle    float64 `json:"idle"`
	IOWait  float64 `json:"iowait"`
	Irq     float64 `json:"irq"`
	Nice    float64 `json:"nice"`
	Total   float64 `json:"total"`
	Cores   int     `json:"cores"`
}

type MemoryStats struct {
	Total    int64   `json:"total"`
	Used     int64   `json:"used"`
	Free     int64   `json:"free"`
	Available int64  `json:"available"`
	Buffers  int64   `json:"buffers"`
	Cached   int64   `json:"cached"`
	SwapTotal int64  `json:"swap_total"`
	SwapUsed int64   `json:"swap_used"`
	UsagePct float64 `json:"usage_pct"`
}

type DiskStats struct {
	Filesystem  string  `json:"filesystem"`
	Size        int64   `json:"size"`
	Used        int64   `json:"used"`
	Available   int64   `json:"available"`
	UsagePct    float64 `json:"usage_pct"`
	Mountpoint  string  `json:"mountpoint"`
}

type NetworkStats struct {
	Interface string  `json:"interface"`
	RxBytes   int64   `json:"rx_bytes"`
	TxBytes   int64   `json:"tx_bytes"`
	RxPackets int64   `json:"rx_packets"`
	TxPackets int64   `json:"tx_packets"`
	RxErrors  int64   `json:"rx_errors"`
	TxErrors  int64   `json:"tx_errors"`
}

type LoadAvgStats struct {
	Load1   float64 `json:"load_1m"`
	Load5   float64 `json:"load_5m"`
	Load15  float64 `json:"load_15m"`
}

type ProcessStats struct {
	Total  int `json:"total"`
	Running int `json:"running"`
	Sleeping int `json:"sleeping"`
	Stopped int `json:"stopped"`
	Zombie  int `json:"zombie"`
}

func getHostname() string {
	h, _ := os.Hostname()
	return h
}

func getUptime() string {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "unknown"
	}
	var secs float64
	fmt.Sscanf(string(data), "%f", &secs)
	d := time.Duration(secs * float64(time.Second))
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	mins := int(d.Minutes()) % 60
	return fmt.Sprintf("%dd %dh %dm", days, hours, mins)
}

func getCPUStats() CPUStats {
	// Read /proc/stat twice with a short delay
	readCPU := func() []float64 {
		data, _ := os.ReadFile("/proc/stat")
		var user, nice, system, idle, iowait, irq, softirq, steal float64
		fmt.Sscanf(string(data), "cpu  %f %f %f %f %f %f %f %f",
			&user, &nice, &system, &idle, &iowait, &irq, &softirq, &steal)
		return []float64{user, nice, system, idle, iowait, irq + softirq, steal}
	}

	before := readCPU()
	time.Sleep(250 * time.Millisecond)
	after := readCPU()

	var dUser, dNice, dSystem, dIdle, dIOWait, dIrq float64
	for i := range before {
		switch i {
		case 0:
			dUser = after[0] - before[0]
		case 1:
			dNice = after[1] - before[1]
		case 2:
			dSystem = after[2] - before[2]
		case 3:
			dIdle = after[3] - before[3]
		case 4:
			dIOWait = after[4] - before[4]
		case 5:
			dIrq = after[5] - before[5]
		}
	}

	total := dUser + dNice + dSystem + dIdle + dIOWait + dIrq
	if total == 0 {
		total = 1
	}

	cores := runtime.NumCPU()
	return CPUStats{
		User:    math.Round((dUser / total) * 1000) / 10,
		System:  math.Round((dSystem / total) * 1000) / 10,
		Idle:    math.Round((dIdle / total) * 1000) / 10,
		IOWait:  math.Round((dIOWait / total) * 1000) / 10,
		Irq:     math.Round((dIrq / total) * 1000) / 10,
		Nice:    math.Round((dNice / total) * 1000) / 10,
		Total:   math.Round(((dUser + dNice + dSystem + dIOWait + dIrq) / total) * 1000) / 10,
		Cores:   cores,
	}
}

func getMemoryStats() MemoryStats {
	var memInfo map[string]int64
	memInfo = make(map[string]int64)

	data, _ := os.ReadFile("/proc/meminfo")
	for _, line := range strings.Split(string(data), "\n") {
		var key string
		var val int64
		if _, err := fmt.Sscanf(line, "%s %d", &key, &val); err == nil {
			memInfo[strings.TrimSuffix(key, ":")] = val * 1024 // kB to bytes
		}
	}

	total := memInfo["MemTotal"]
	used := total - memInfo["MemFree"] - memInfo["Buffers"] - memInfo["Cached"] - memInfo["SReclaimable"]
	if used < 0 {
		used = total - memInfo["MemAvailable"]
	}

	pct := 0.0
	if total > 0 {
		pct = math.Round(float64(used)/float64(total)*1000) / 10
	}

	swapUsed := int64(0)
	swapTotal := memInfo["SwapTotal"]
	if memInfo["SwapFree"] < memInfo["SwapTotal"] {
		swapUsed = memInfo["SwapTotal"] - memInfo["SwapFree"]
	}

	return MemoryStats{
		Total:     total,
		Used:      used,
		Free:      memInfo["MemFree"],
		Available: memInfo["MemAvailable"],
		Buffers:   memInfo["Buffers"],
		Cached:    memInfo["Cached"],
		SwapTotal: swapTotal,
		SwapUsed:  swapUsed,
		UsagePct:  pct,
	}
}

func getDiskStats() []DiskStats {
	cmd := exec.Command("df", "-B1", "-P", "-x", "tmpfs", "-x", "devtmpfs", "-x", "overlay")
	output, err := cmd.Output()
	if err != nil {
		cmd = exec.Command("df", "-B1", "-P")
		output, err = cmd.Output()
		if err != nil {
			return nil
		}
	}

	var disks []DiskStats
	lines := strings.Split(string(output), "\n")
	for i, line := range lines {
		if i == 0 {
			continue // header
		}
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}
		fs := fields[0]
		size, _ := strconv.ParseInt(fields[1], 10, 64)
		used, _ := strconv.ParseInt(fields[2], 10, 64)
		avail, _ := strconv.ParseInt(fields[3], 10, 64)
		pctStr := strings.TrimSuffix(fields[4], "%")
		pct, _ := strconv.ParseFloat(pctStr, 64)
		mount := fields[5]

		disks = append(disks, DiskStats{
			Filesystem: fs,
			Size:       size,
			Used:       used,
			Available:  avail,
			UsagePct:   pct,
			Mountpoint: mount,
		})
	}
	return disks
}

func getNetworkStats() []NetworkStats {
	data, err := os.ReadFile("/proc/net/dev")
	if err != nil {
		return nil
	}

	var nets []NetworkStats
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if i < 2 {
			continue // header lines
		}
		if line == "" {
			continue
		}
		fields := strings.Fields(strings.TrimSpace(line))
		if len(fields) < 9 {
			continue
		}
		iface := strings.TrimSuffix(fields[0], ":")
		rxBytes, _ := strconv.ParseInt(fields[1], 10, 64)
		rxPackets, _ := strconv.ParseInt(fields[2], 10, 64)
		rxErrors, _ := strconv.ParseInt(fields[3], 10, 64)
		txBytes, _ := strconv.ParseInt(fields[9], 10, 64)
		txPackets, _ := strconv.ParseInt(fields[10], 10, 64)
		txErrors, _ := strconv.ParseInt(fields[11], 10, 64)

		nets = append(nets, NetworkStats{
			Interface: iface,
			RxBytes:   rxBytes,
			TxBytes:   txBytes,
			RxPackets: rxPackets,
			TxPackets: txPackets,
			RxErrors:  rxErrors,
			TxErrors:  txErrors,
		})
	}
	return nets
}

func getLoadAvg() LoadAvgStats {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return LoadAvgStats{}
	}
	var l1, l5, l15 float64
	fmt.Sscanf(string(data), "%f %f %f", &l1, &l5, &l15)
	return LoadAvgStats{Load1: math.Round(l1*100) / 100, Load5: math.Round(l5*100) / 100, Load15: math.Round(l15*100) / 100}
}

func getProcessStats() ProcessStats {
	data, err := os.ReadFile("/proc/stat")
	if err != nil {
		return ProcessStats{}
	}
	var total, running, blocked int
	fmt.Sscanf(string(data), "procs_running %d procs_blocked %d", &running, &blocked)

	// Count processes
	pids, _ := os.ReadDir("/proc")
	total = 0
	sleeping := 0
	stopped := 0
	zombie := 0
	for _, pid := range pids {
		if !pid.IsDir() {
			continue
		}
		if _, err := strconv.Atoi(pid.Name()); err != nil {
			continue
		}
		total++
		status, _ := os.ReadFile("/proc/" + pid.Name() + "/status")
		for _, line := range strings.Split(string(status), "\n") {
			if strings.HasPrefix(line, "State:") {
				fields := strings.Fields(line)
				if len(fields) >= 2 {
					switch fields[1] {
					case "S", "D":
						sleeping++
					case "T", "t":
						stopped++
					case "Z":
						zombie++
					case "R":
						// already counted in running
					}
				}
				break
			}
		}
	}
	return ProcessStats{Total: total, Running: running, Sleeping: sleeping, Stopped: stopped, Zombie: zombie}
}

func collectStats() SystemStats {
	return SystemStats{
		Timestamp: time.Now().Format(time.RFC3339),
		Hostname:  getHostname(),
		Uptime:    getUptime(),
		CPU:       getCPUStats(),
		Memory:    getMemoryStats(),
		Disk:      getDiskStats(),
		Network:   getNetworkStats(),
		LoadAvg:   getLoadAvg(),
		Processes: getProcessStats(),
	}
}

// ANSI colors
const (
	reset   = "\033[0m"
	bold    = "\033[1m"
	red     = "\033[1;31m"
	yellow  = "\033[1;33m"
	green   = "\033[1;32m"
	blue    = "\033[1;34m"
	cyan    = "\033[1;36m"
	gray    = "\033[90m"
	magenta = "\033[1;35m"
)

func formatBytes(b int64) string {
	if b < 1024 {
		return fmt.Sprintf("%d B", b)
	}
	if b < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(b)/1024)
	}
	if b < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(b)/(1024*1024))
	}
	return fmt.Sprintf("%.2f GB", float64(b)/(1024*1024*1024))
}

func usageColor(pct float64) string {
	switch {
	case pct >= 90:
		return red
	case pct >= 70:
		return yellow
	default:
		return green
	}
}

func progressBar(pct float64, width int) string {
	filled := int(math.Round(pct / 100 * float64(width)))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}
	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	return bar
}

func printStats(s SystemStats) {
	fmt.Printf("\n%s╔══════════════════════════════════════════════════════════════════╗%s\n", bold, reset)
	fmt.Printf("%s║%s %-58s %s║%s\n", bold, gray, "VM STATS", reset, bold)
	fmt.Printf("%s╚══════════════════════════════════════════════════════════════════╝%s\n", bold, reset)

	// Host info
	fmt.Printf("\n  %sHost:%s %-20s %sUptime:%s %s\n", cyan, s.Hostname, gray, cyan, s.Uptime)
	fmt.Printf("  %sTimestamp:%s %s\n", cyan, gray, s.Timestamp)

	// CPU
	fmt.Printf("\n  %s┌─ CPU ──────────────────────────────────────────────────────────┐%s\n", bold, reset)
	fmt.Printf("  %sCores:%s %-4d %sLoad Avg:%s %.2f %.2f %.2f\n", blue, s.CPU.Cores, gray, blue, s.LoadAvg.Load1, s.LoadAvg.Load5, s.LoadAvg.Load15)
	cpuColor := usageColor(s.CPU.Total)
	fmt.Printf("  %sUsage:%s  %s%-6.1f%% %s%s%s\n", blue, reset, cpuColor, s.CPU.Total, progressBar(s.CPU.Total, 40), usageColor(s.CPU.Total)+reset, reset)
	fmt.Printf("  %s├─%s %sUser:%s %5.1f%%  %sSystem:%s %5.1f%%  %sIOWait:%s %5.1f%%  %sIRQ:%s %5.1f%%\n",
		reset, reset, gray, reset, s.CPU.User, gray, reset, s.CPU.System, gray, reset, s.CPU.IOWait, gray, reset, s.CPU.Irq)
	fmt.Printf("  %s└─%s %sIdle:%s %5.1f%%  %sNice:%s %5.1f%%\n", reset, reset, gray, reset, s.CPU.Idle, gray, reset, s.CPU.Nice)

	// Memory
	fmt.Printf("\n  %s┌─ Memory ───────────────────────────────────────────────────────┐%s\n", bold, reset)
	memColor := usageColor(s.Memory.UsagePct)
	fmt.Printf("  %sUsed:%s   %s%-10s / %-10s %s%-6.1f%% %s%s%s\n",
		blue, reset, memColor, formatBytes(s.Memory.Used), formatBytes(s.Memory.Total), memColor, reset, progressBar(s.Memory.UsagePct, 30), reset)
	fmt.Printf("  %s├─%s %sAvailable:%s %-10s %sFree:%s %-10s\n", reset, reset, gray, reset, formatBytes(s.Memory.Available), gray, reset, formatBytes(s.Memory.Free))
	fmt.Printf("  %s├─%s %sBuffers:%s %-10s %sCached:%s %-10s\n", reset, reset, gray, reset, formatBytes(s.Memory.Buffers), gray, reset, formatBytes(s.Memory.Cached))
	if s.Memory.SwapTotal > 0 {
		swapPct := float64(s.Memory.SwapUsed) / float64(s.Memory.SwapTotal) * 100
		swapColor := usageColor(swapPct)
		fmt.Printf("  %s└─%s %sSwap:%s   %s%-10s / %-10s %s%-6.1f%%\n", reset, reset, gray, reset, swapColor, formatBytes(s.Memory.SwapUsed), formatBytes(s.Memory.SwapTotal), swapColor, swapPct)
	} else {
		fmt.Printf("  %s└─%s %sSwap:%s   none configured\n", reset, reset, gray, reset)
	}

	// Disk
	fmt.Printf("\n  %s┌─ Disk ─────────────────────────────────────────────────────────┐%s\n", bold, reset)
	for i, d := range s.Disk {
		dColor := usageColor(d.UsagePct)
		fmt.Printf("  %s%-18s %s%-8.1f%% %s%s%s %s%s / %s\n",
			gray+d.Mountpoint+reset, dColor, d.UsagePct, dColor, progressBar(d.UsagePct, 30), reset, gray, formatBytes(d.Used), formatBytes(d.Size))
		if i < len(s.Disk)-1 {
			fmt.Printf("  %s│%s\n", reset, reset)
		}
	}
	fmt.Printf("  %s└────────────────────────────────────────────────────────────────┘%s\n", reset, reset)

	// Network
	if len(s.Network) > 0 {
		fmt.Printf("\n  %s┌─ Network ──────────────────────────────────────────────────────┐%s\n", bold, reset)
		for i, n := range s.Network {
			fmt.Printf("  %s%-12s %sRX:%s %-12s %sTX:%s %-12s\n",
				cyan, n.Interface, reset, formatBytes(n.RxBytes), reset, formatBytes(n.TxBytes))
			if n.RxErrors > 0 || n.TxErrors > 0 {
				fmt.Printf("  %s              %sErrors:%s RX:%d TX:%d\n", reset, red, reset, n.RxErrors, n.TxErrors)
			}
			if i < len(s.Network)-1 {
				fmt.Printf("  %s│%s\n", reset, reset)
			}
		}
		fmt.Printf("  %s└────────────────────────────────────────────────────────────────┘%s\n", reset, reset)
	}

	// Processes
	fmt.Printf("\n  %s┌─ Processes ────────────────────────────────────────────────────┐%s\n", bold, reset)
	fmt.Printf("  %sTotal:%s %-6d %sRunning:%s %-4d %sSleeping:%s %-4d\n",
		blue, reset, s.Processes.Total, blue, reset, s.Processes.Running, blue, reset, s.Processes.Sleeping)
	if s.Processes.Zombie > 0 {
		fmt.Printf("  %sZombie:%s %d %sStopped:%s %d\n", red, s.Processes.Zombie, reset, gray, s.Processes.Stopped)
	} else {
		fmt.Printf("  %sStopped:%s %d\n", gray, reset, s.Processes.Stopped)
	}
	fmt.Printf("  %s└────────────────────────────────────────────────────────────────┘%s\n", reset, reset)
}

func printJSON(s SystemStats) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(s)
	fmt.Fprintln(os.Stdout)
}

func printCompact(s SystemStats) {
	memColor := usageColor(s.Memory.UsagePct)
	cpuColor := usageColor(s.CPU.Total)

	fmt.Printf("%-10s CPU %s%-5.1f%%%s MEM %s%-5.1f%%%s ",
		s.Timestamp[11:19], cpuColor, s.CPU.Total, reset, memColor, s.Memory.UsagePct, reset)
	fmt.Printf("LOAD %.2f %.2f %.2f  ", s.LoadAvg.Load1, s.LoadAvg.Load5, s.LoadAvg.Load15)
	fmt.Printf("PROC %d  ", s.Processes.Total)

	// Disk usage for /
	for _, d := range s.Disk {
		if d.Mountpoint == "/" {
			dColor := usageColor(d.UsagePct)
			fmt.Printf("DISK %s%-5.1f%%%s ", dColor, d.UsagePct, reset)
			break
		}
	}
	fmt.Println()
}

func main() {
	var (
		format  string
		interval int
		count   int
		jsonOut bool
		compact bool
	)

	flag.StringVar(&format, "format", "full", "Output format: full, compact, json")
	flag.IntVar(&interval, "interval", 1, "Refresh interval in seconds")
	flag.IntVar(&count, "count", 0, "Number of updates (0=infinite)")
	flag.BoolVar(&jsonOut, "json", false, "Output as JSON")
	flag.BoolVar(&compact, "compact", false, "Compact one-line output")
	flag.BoolVar(&jsonOut, "j", false, "Short for --json")
	flag.BoolVar(&compact, "c", false, "Short for --compact")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: vm-stats [options]\n\n")
		fmt.Fprintf(os.Stderr, "Collects and displays system statistics (CPU, RAM, disk, network).\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  vm-stats                    # One-shot full report\n")
		fmt.Fprintf(os.Stderr, "  vm-stats -interval 5        # Refresh every 5 seconds\n")
		fmt.Fprintf(os.Stderr, "  vm-stats -compact           # Compact one-line output\n")
		fmt.Fprintf(os.Stderr, "  vm-stats -json              # JSON output\n")
		fmt.Fprintf(os.Stderr, "  vm-stats -compact -interval 2 -count 10\n")
	}

	flag.Parse()

	if jsonOut {
		format = "json"
	}
	if compact {
		format = "compact"
	}

	if count == 0 {
		count = 1
	}

	i := 0
	for i < count {
		stats := collectStats()

		switch format {
		case "json":
			printJSON(stats)
		case "compact":
			printCompact(stats)
		default:
			if i > 0 {
				fmt.Print("\033[2J\033[H") // Clear screen
			}
			printStats(stats)
		}

		i++
		if i < count {
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}
}
