// ssh-audit audits SSH server configurations by connecting to a remote server
// and enumerating supported ciphers, MACs, key exchange algorithms, and host
// key algorithms for security weaknesses.
//
// It works by opening a TCP connection, reading the SSH banner, sending a
// KEX_INIT packet, and parsing the server's response.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// Rating represents the security rating of an algorithm.
type Rating int

const (
	RatingExcellent Rating = iota
	RatingGood
	RatingFair
	RatingWeak
	RatingBroken
	RatingUnknown
)

func (r Rating) String() string {
	switch r {
	case RatingExcellent:
		return "excellent"
	case RatingGood:
		return "good"
	case RatingFair:
		return "fair"
	case RatingWeak:
		return "weak"
	case RatingBroken:
		return "broken"
	default:
		return "unknown"
	}
}

func (r Rating) Color() string {
	switch r {
	case RatingExcellent:
		return "\033[32m"
	case RatingGood:
		return "\033[36m"
	case RatingFair:
		return "\033[33m"
	case RatingWeak:
		return "\033[31m"
	case RatingBroken:
		return "\033[35m"
	default:
		return "\033[90m"
	}
}

type algoInfo struct {
	Rating Rating
	Notes  string
}

var kexDB = map[string]algoInfo{
	"curve25519-sha256":                       {RatingExcellent, "RFC 8731"},
	"curve25519-sha256@libssh.org":            {RatingExcellent, "RFC 8731"},
	"sntrup761x25519-sha512@openssh.com":      {RatingExcellent, "post-quantum hybrid"},
	"diffie-hellman-group16-sha512":           {RatingGood, "4096-bit MODP"},
	"diffie-hellman-group18-sha512":           {RatingGood, "8192-bit MODP"},
	"diffie-hellman-group-exchange-sha256":    {RatingGood, "2048-bit minimum"},
	"diffie-hellman-group14-sha256":           {RatingGood, "2048-bit MODP"},
	"ecdh-sha2-nistp256":                      {RatingFair, "NIST P-256"},
	"ecdh-sha2-nistp384":                      {RatingFair, "NIST P-384"},
	"ecdh-sha2-nistp521":                      {RatingFair, "NIST P-521"},
	"diffie-hellman-group14-sha1":             {RatingWeak, "SHA-1 deprecated"},
	"diffie-hellman-group-exchange-sha1":      {RatingWeak, "SHA-1 deprecated"},
	"diffie-hellman-group1-sha1":              {RatingBroken, "1024-bit, Logjam"},
}

var cipherDB = map[string]algoInfo{
	"chacha20-poly1305@openssh.com": {RatingExcellent, "AEAD"},
	"aes128-gcm@openssh.com":        {RatingExcellent, "AEAD GCM"},
	"aes256-gcm@openssh.com":        {RatingExcellent, "AEAD GCM"},
	"aes128-ctr":                    {RatingGood, "AES-128 CTR"},
	"aes192-ctr":                    {RatingGood, "AES-192 CTR"},
	"aes256-ctr":                    {RatingGood, "AES-256 CTR"},
	"aes128-cbc":                    {RatingFair, "CBC no AEAD"},
	"aes192-cbc":                    {RatingFair, "CBC no AEAD"},
	"aes256-cbc":                    {RatingFair, "CBC no AEAD"},
	"blowfish-cbc":                  {RatingWeak, "64-bit block, Sweet32"},
	"3des-cbc":                      {RatingBroken, "Sweet32, deprecated"},
	"arcfour128":                    {RatingBroken, "RC4 broken"},
	"arcfour256":                    {RatingBroken, "RC4 broken"},
	"arcfour":                       {RatingBroken, "RC4 broken"},
	"cast128-cbc":                   {RatingBroken, "deprecated"},
}

var macDB = map[string]algoInfo{
	"hmac-sha2-256-etm@openssh.com": {RatingExcellent, "encrypt-then-MAC"},
	"hmac-sha2-512-etm@openssh.com": {RatingExcellent, "encrypt-then-MAC"},
	"hmac-sha2-256":                 {RatingGood, "HMAC-SHA256"},
	"hmac-sha2-512":                 {RatingGood, "HMAC-SHA512"},
	"umac-128-etm@openssh.com":      {RatingGood, "UMAC-128 ETM"},
	"umac-128@openssh.com":          {RatingGood, "UMAC-128"},
	"hmac-sha1-etm@openssh.com":     {RatingFair, "SHA-1 but ETM"},
	"hmac-sha1":                     {RatingWeak, "SHA-1 deprecated"},
	"umac-64@openssh.com":           {RatingWeak, "64-bit tag"},
	"hmac-md5":                      {RatingBroken, "MD5 broken"},
	"hmac-md5-96":                   {RatingBroken, "truncated MD5"},
	"hmac-sha1-96":                  {RatingBroken, "truncated SHA-1"},
}

var hostkeyDB = map[string]algoInfo{
	"ssh-ed25519":       {RatingExcellent, "Ed25519"},
	"rsa-sha2-512":      {RatingGood, "RSA SHA-512"},
	"rsa-sha2-256":      {RatingGood, "RSA SHA-256"},
	"ecdsa-sha2-nistp256": {RatingGood, "ECDSA P-256"},
	"ecdsa-sha2-nistp384": {RatingGood, "ECDSA P-384"},
	"ecdsa-sha2-nistp521": {RatingGood, "ECDSA P-521"},
	"ssh-rsa":           {RatingWeak, "SHA-1 signatures"},
	"ssh-dss":           {RatingBroken, "DSA 1024-bit, deprecated"},
}

type ServerInfo struct {
	Host           string
	Port           int
	KexAlgorithms  []string
	Ciphers        []string
	MACs           []string
	HostKeyAlgos   []string
	ServerVersion  string
	ResponseTime   time.Duration
	HasETM         bool
	HasAEAD        bool
	HasPostQuantum bool
	BrokenCount    int
	WeakCount      int
}

func getRating(algo string, db map[string]algoInfo) (Rating, string) {
	if info, ok := db[algo]; ok {
		return info.Rating, info.Notes
	}
	return RatingUnknown, "unknown algorithm"
}

// buildKexInit builds a minimal SSH KEX_INIT packet (msg type 20).
func buildKexInit() []byte {
	// Cookie: 16 random-ish bytes
	cookie := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10}

	// We list all known algorithms so the server tells us what it supports.
	// But actually, the server sends its OWN kex_init with its preferences.
	// We just need to send any valid kex_init to trigger the response.
	kexList := "curve25519-sha256,curve25519-sha256@libssh.org,diffie-hellman-group16-sha512,diffie-hellman-group14-sha256,diffie-hellman-group-exchange-sha256,ecdh-sha2-nistp256,ecdh-sha2-nistp384,ecdh-sha2-nistp521,diffie-hellman-group14-sha1,diffie-hellman-group1-sha1"
	hostKeyList := "ssh-ed25519,rsa-sha2-512,rsa-sha2-256,ecdsa-sha2-nistp256,ssh-rsa,ssh-dss"
	cipherC := "chacha20-poly1305@openssh.com,aes256-gcm@openssh.com,aes128-gcm@openssh.com,aes256-ctr,aes192-ctr,aes128-ctr,aes256-cbc,aes128-cbc,3des-cbc,blowfish-cbc"
	cipherS := cipherC
	macC := "hmac-sha2-512-etm@openssh.com,hmac-sha2-256-etm@openssh.com,hmac-sha2-512,hmac-sha2-256,hmac-sha1-etm@openssh.com,hmac-sha1,hmac-md5"
	macS := macC
	comprC := "none,zlib@openssh.com,zlib"
	comprS := comprC
	langC := ""
	langS := ""

	var buf bytes.Buffer
	// Length placeholder (4 bytes)
	buf.Write([]byte{0, 0, 0, 0})

	writeString := func(s string) {
		b := []byte(s)
		buf.Write([]byte{byte(len(b) >> 24), byte(len(b) >> 16), byte(len(b) >> 8), byte(len(b))})
		buf.Write(b)
	}

	buf.WriteByte(20) // SSH_MSG_KEXINIT
	buf.Write(cookie)
	writeString(kexList)
	writeString(hostKeyList)
	writeString(cipherC)
	writeString(cipherS)
	writeString(macC)
	writeString(macS)
	writeString(comprC)
	writeString(comprS)
	writeString(langC)
	writeString(langS)
	buf.WriteByte(0) // first_kex_packet_follows = false

	length := uint32(buf.Len() - 4)
	buf.Bytes()[0] = byte(length >> 24)
	buf.Bytes()[1] = byte(length >> 16)
	buf.Bytes()[2] = byte(length >> 8)
	buf.Bytes()[3] = byte(length)

	return buf.Bytes()
}

// readSSHString reads a length-prefixed string from the buffer.
func readSSHString(data []byte, offset int) (string, int, error) {
	if offset+4 > len(data) {
		return "", offset, fmt.Errorf("truncated length")
	}
	strLen := uint32(data[offset])<<24 | uint32(data[offset+1])<<16 | uint32(data[offset+2])<<8 | uint32(data[offset+3])
	offset += 4
	if offset+int(strLen) > len(data) {
		return "", offset, fmt.Errorf("truncated string")
	}
	s := string(data[offset : offset+int(strLen)])
	return s, offset + int(strLen), nil
}

// parseKexInit parses an SSH_MSG_KEXINIT packet and extracts algorithm lists.
func parseKexInit(data []byte) ([]string, []string, []string, []string, []string, []string, error) {
	if len(data) < 5 {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("packet too short")
	}

	// Skip length (4 bytes)
	offset := 4

	// Check message type
	if data[offset] != 20 {
		// Might be SSH_DISCONNECT or SSH_IGNORE; try to find KEXINIT
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("not a KEX_INIT message (type=%d)", data[offset])
	}
	offset++

	// Skip cookie (16 bytes)
	if offset+16 > len(data) {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("truncated cookie")
	}
	offset += 16

	readList := func() ([]string, error) {
		s, nextOff, err := readSSHString(data, offset)
		if err != nil {
			return nil, err
		}
		offset = nextOff
		return strings.Split(s, ","), nil
	}

	kex, err := readList()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	hostKey, err := readList()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	cipherC, err := readList()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	cipherS, err := readList()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	macC, err := readList()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	macS, err := readList()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return kex, hostKey, cipherC, cipherS, macC, macS, nil
}

func readSSHPacket(conn net.Conn) ([]byte, error) {
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	// Read length (4 bytes)
	var lenBuf [4]byte
	if _, err := bufio.NewReader(conn).Read(lenBuf[:]); err != nil {
		// Fallback: read byte by byte
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		return nil, err
	}

	length := uint32(lenBuf[0])<<24 | uint32(lenBuf[1])<<16 | uint32(lenBuf[2])<<8 | uint32(lenBuf[3])
	if length == 0 || length > 35000 {
		return nil, fmt.Errorf("invalid packet length: %d", length)
	}

	// Read the rest
	packet := make([]byte, length)
	packet[0] = lenBuf[0]
	packet[1] = lenBuf[1]
	packet[2] = lenBuf[2]
	packet[3] = lenBuf[3]

	n, err := conn.Read(packet[4:])
	if err != nil && n < int(length)-4 {
		return nil, fmt.Errorf("reading packet: %w", err)
	}

	return packet, nil
}

func probeSSH(host string, port int) (*ServerInfo, error) {
	addr := net.JoinHostPort(host, strconv.Itoa(port))
	start := time.Now()

	conn, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}
	defer conn.Close()

	// Read SSH banner
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	reader := bufio.NewReader(conn)
	banner, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("reading banner: %w", err)
	}
	banner = strings.TrimSpace(banner)
	if !strings.HasPrefix(banner, "SSH-") {
		return nil, fmt.Errorf("not an SSH server: %s", banner[:min(len(banner), 80)])
	}
	// Remove SSH-2.0- or SSH-1.0- prefix to get version string
	serverVersion := strings.TrimPrefix(banner, "SSH-2.0-")
	serverVersion = strings.TrimPrefix(serverVersion, "SSH-1.0-")

	// Send our KEX_INIT
	kexInit := buildKexInit()
	_, err = conn.Write(kexInit)
	if err != nil {
		return nil, fmt.Errorf("writing kex_init: %w", err)
	}

	// Read server's KEX_INIT (may need to skip some bytes from banner read)
	elapsed := time.Since(start)

	info := &ServerInfo{
		Host:          host,
		Port:          port,
		ServerVersion: serverVersion,
		ResponseTime:  elapsed,
	}

	// Try to read packets until we get KEX_INIT
	for i := 0; i < 20; i++ {
		packet, err := readSSHPacket(conn)
		if err != nil {
			// Try one more approach: read remaining data from buffer
			break
		}

		kex, hostKey, cipherC, cipherS, macC, macS, err := parseKexInit(packet)
		if err != nil {
			continue // Not a KEX_INIT, skip
		}

		info.KexAlgorithms = kex
		info.HostKeyAlgos = hostKey
		info.Ciphers = uniqueStrings(append(cipherC, cipherS...))
		info.MACs = uniqueStrings(append(macC, macS...))

		analyzeInfo(info)
		return info, nil
	}

	// If we couldn't parse KEX_INIT but got a connection, return what we have
	if info.KexAlgorithms == nil {
		return nil, fmt.Errorf("could not parse server KEX_INIT (got banner: %s)", info.ServerVersion)
	}

	analyzeInfo(info)
	return info, nil
}

func uniqueStrings(in []string) []string {
	seen := make(map[string]bool)
	var out []string
	for _, s := range in {
		s = strings.TrimSpace(s)
		if s != "" && !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}

func analyzeInfo(info *ServerInfo) {
	for _, c := range info.Ciphers {
		if strings.Contains(c, "gcm") || strings.Contains(c, "chacha20") {
			info.HasAEAD = true
			break
		}
	}
	for _, m := range info.MACs {
		if strings.Contains(m, "etm") {
			info.HasETM = true
			break
		}
	}
	for _, k := range info.KexAlgorithms {
		if strings.Contains(k, "sntrup") || strings.Contains(k, "pqc") {
			info.HasPostQuantum = true
			break
		}
	}

	dbs := []map[string]algoInfo{kexDB, cipherDB, macDB, hostkeyDB}
	lists := [][]string{info.KexAlgorithms, info.Ciphers, info.MACs, info.HostKeyAlgos}
	for i, list := range lists {
		for _, a := range list {
			r, _ := getRating(a, dbs[i])
			if r == RatingBroken {
				info.BrokenCount++
			}
			if r == RatingWeak {
				info.WeakCount++
			}
		}
	}
}

func printAudit(info *ServerInfo) {
	reset := "\033[0m"
	bold := "\033[1m"

	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    SSH SERVER AUDIT                         ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	fmt.Printf("%sTarget:%s %s:%d\n", bold, reset, info.Host, info.Port)
	fmt.Printf("%sServer:%s %s\n", bold, reset, info.ServerVersion)
	fmt.Printf("%sResponse:%s %s\n", bold, reset, info.ResponseTime.Round(time.Millisecond))
	fmt.Println()

	fmt.Printf("%sOverall:%s ", bold, reset)
	if info.BrokenCount > 0 {
		fmt.Printf("\033[31mFAIL\033[0m (%d broken, %d weak)\n", info.BrokenCount, info.WeakCount)
	} else if info.WeakCount > 0 {
		fmt.Printf("\033[33mWARNING\033[0m (%d weak)\n", info.WeakCount)
	} else {
		fmt.Printf("\033[32mPASS\033[0m\n")
	}
	fmt.Println()

	fmt.Printf("  AEAD Ciphers:     %s\n", yesNo(info.HasAEAD))
	fmt.Printf("  Encrypt-then-MAC: %s\n", yesNo(info.HasETM))
	fmt.Printf("  Post-Quantum KEX: %s\n", yesNo(info.HasPostQuantum))
	fmt.Println()

	fmt.Println("── Key Exchange Algorithms ───────────────────────────────────")
	printAlgoList(info.KexAlgorithms, kexDB)
	fmt.Println()

	fmt.Println("── Ciphers ───────────────────────────────────────────────────")
	printAlgoList(info.Ciphers, cipherDB)
	fmt.Println()

	fmt.Println("── MAC Algorithms ────────────────────────────────────────────")
	printAlgoList(info.MACs, macDB)
	fmt.Println()

	fmt.Println("── Host Key Algorithms ───────────────────────────────────────")
	printAlgoList(info.HostKeyAlgos, hostkeyDB)
	fmt.Println()

	if info.BrokenCount > 0 || info.WeakCount > 0 {
		printRecommendations(info)
	}
}

func printAlgoList(algos []string, db map[string]algoInfo) {
	for _, algo := range algos {
		rating, notes := getRating(algo, db)
		fmt.Printf("  %s%-45s %s%-10s\033[0m \033[90m%s\033[0m\n",
			rating.Color(), algo, rating.Color(), strings.ToUpper(rating.String()), notes)
	}
}

func yesNo(v bool) string {
	if v {
		return "\033[32mYES\033[0m"
	}
	return "\033[31mNO\033[0m"
}

func printRecommendations(info *ServerInfo) {
	bold := "\033[1m"
	reset := "\033[0m"

	fmt.Println("── Recommendations ───────────────────────────────────────────")

	collect := func(list []string, db map[string]algoInfo) ([]string, []string) {
		var broken, weak []string
		for _, a := range list {
			r, _ := getRating(a, db)
			if r == RatingBroken {
				broken = append(broken, a)
			}
			if r == RatingWeak {
				weak = append(weak, a)
			}
		}
		return broken, weak
	}

	printRec := func(category string, broken, weak []string) {
		if len(broken) > 0 {
			fmt.Printf("  %sREMOVE (%s):%s  %s\n", bold, category, reset, strings.Join(broken, ", "))
		}
		if len(weak) > 0 {
			fmt.Printf("  %sCONSIDER (%s):%s %s\n", bold, category, reset, strings.Join(weak, ", "))
		}
	}

	bk, wk := collect(info.KexAlgorithms, kexDB)
	printRec("KEX", bk, wk)
	bc, wc := collect(info.Ciphers, cipherDB)
	printRec("Cipher", bc, wc)
	bm, wm := collect(info.MACs, macDB)
	printRec("MAC", bm, wm)
	bh, wh := collect(info.HostKeyAlgos, hostkeyDB)
	printRec("HostKey", bh, wh)

	if !info.HasAEAD {
		fmt.Printf("  %sSUGGEST:%s Add AEAD ciphers (chacha20-poly1305, aes-*-gcm)\n", bold, reset)
	}
	if !info.HasETM {
		fmt.Printf("  %sSUGGEST:%s Add encrypt-then-MAC algorithms\n", bold, reset)
	}

	fmt.Println()
	fmt.Println("  Example sshd_config hardening:")
	fmt.Println("    KexAlgorithms curve25519-sha256,curve25519-sha256@libssh.org,diffie-hellman-group16-sha512")
	fmt.Println("    Ciphers chacha20-poly1305@openssh.com,aes256-gcm@openssh.com,aes128-gcm@openssh.com")
	fmt.Println("    MACs hmac-sha2-512-etm@openssh.com,hmac-sha2-256-etm@openssh.com,hmac-sha2-512,hmac-sha2-256")
	fmt.Println("    HostKeyAlgorithms ssh-ed25519,rsa-sha2-512,rsa-sha2-256")
	fmt.Println()
}

type jsonAlgo struct {
	Name   string `json:"name"`
	Rating string `json:"rating"`
	Notes  string `json:"notes"`
}

type jsonOutput struct {
	Target         string     `json:"target"`
	ServerVersion  string     `json:"server_version"`
	ResponseTime   string     `json:"response_time"`
	KexAlgorithms  []jsonAlgo `json:"kex_algorithms"`
	Ciphers        []jsonAlgo `json:"ciphers"`
	MACs           []jsonAlgo `json:"macs"`
	HostKeyAlgos   []jsonAlgo `json:"host_key_algorithms"`
	HasETM         bool       `json:"has_etm"`
	HasAEAD        bool       `json:"has_aead"`
	HasPostQuantum bool       `json:"has_post_quantum"`
	BrokenCount    int        `json:"broken_count"`
	WeakCount      int        `json:"weak_count"`
}

func toJSONAlgo(list []string, db map[string]algoInfo) []jsonAlgo {
	var out []jsonAlgo
	for _, name := range list {
		r, notes := getRating(name, db)
		out = append(out, jsonAlgo{Name: name, Rating: r.String(), Notes: notes})
	}
	return out
}

func printJSON(info *ServerInfo) {
	output := jsonOutput{
		Target:         fmt.Sprintf("%s:%d", info.Host, info.Port),
		ServerVersion:  info.ServerVersion,
		ResponseTime:   info.ResponseTime.String(),
		KexAlgorithms:  toJSONAlgo(info.KexAlgorithms, kexDB),
		Ciphers:        toJSONAlgo(info.Ciphers, cipherDB),
		MACs:           toJSONAlgo(info.MACs, macDB),
		HostKeyAlgos:   toJSONAlgo(info.HostKeyAlgos, hostkeyDB),
		HasETM:         info.HasETM,
		HasAEAD:        info.HasAEAD,
		HasPostQuantum: info.HasPostQuantum,
		BrokenCount:    info.BrokenCount,
		WeakCount:      info.WeakCount,
	}

	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func usage() {
	fmt.Fprintf(os.Stderr, `ssh-audit - Audit SSH server security configuration

Usage:
  ssh-audit [options] <host>[:port]

Options:
  -json       Output results as JSON
  -p, --port  Specify port (default: 22)
  -h, --help  Show this help message

Examples:
  ssh-audit localhost
  ssh-audit example.com:22
  ssh-audit -p 2222 myserver.com
  ssh-audit -json localhost

The tool connects to the SSH server, performs a KEX_INIT exchange,
and reports on the security of supported algorithms with color-coded ratings.
`)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		os.Exit(1)
	}

	var jsonOutput bool
	port := 22
	var host string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-json":
			jsonOutput = true
		case "-h", "--help":
			usage()
			os.Exit(0)
		case "-p", "--port":
			if i+1 >= len(args) {
				fmt.Fprintln(os.Stderr, "Error: -p requires a port number")
				os.Exit(1)
			}
			i++
			p, err := strconv.Atoi(args[i])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid port: %s\n", args[i])
				os.Exit(1)
			}
			port = p
		default:
			if strings.HasPrefix(arg, "-") {
				fmt.Fprintf(os.Stderr, "Unknown option: %s\n", arg)
				usage()
				os.Exit(1)
			}
			host = arg
		}
	}

	// Parse host:port
	if host != "" {
		parts := strings.Split(host, ":")
		if len(parts) == 2 {
			host = parts[0]
			if p, err := strconv.Atoi(parts[1]); err == nil {
				port = p
			}
		}
	}

	if host == "" {
		fmt.Fprintln(os.Stderr, "Error: no host specified")
		usage()
		os.Exit(1)
	}

	info, err := probeSSH(host, port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if jsonOutput {
		printJSON(info)
	} else {
		printAudit(info)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
