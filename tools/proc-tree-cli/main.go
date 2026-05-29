// proc-tree-cli displays the process tree from /proc.
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Process struct {
	PID      int
	PPID     int
	Name     string
	User     string
	State    string
	CPU      float64
	Mem      float64
	Children []*Process
}

func main() {
	filterUser := ""
	showAll := false
	format := "tree"
	pidFilter := 0

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-u", "--user":
			i++
			if i < len(os.Args) {
				filterUser = os.Args[i]
			}
		case "-a", "--all":
			showAll = true
		case "-f", "--format":
			i++
			if i < len(os.Args) {
				format = os.Args[i]
			}
		case "-p", "--pid":
			i++
			if i < len(os.Args) {
				fmt.Sscanf(os.Args[i], "%d", &pidFilter)
			}
		case "-h", "--help":
			usage()
			return
		}
	}

	procs, err := readProcs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	tree := buildTree(procs, filterUser)

	if format == "flat" {
		printFlat(tree, showAll)
	} else {
		if pidFilter > 0 {
			printSubtree(findPID(tree, pidFilter), 0, true, showAll)
		} else {
			for _, root := range tree {
				printSubtree(root, 0, true, showAll)
			}
		}
	}
}

func readProcs() ([]*Process, error) {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	var procs []*Process
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		pid, err := strconv.Atoi(e.Name())
		if err != nil {
			continue
		}

		p := &Process{PID: pid}
		p.readStatus()
		p.readStat()
		procs = append(procs, p)
	}
	return procs, nil
}

func (p *Process) readStatus() {
	path := fmt.Sprintf("/proc/%d/status", p.PID)
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "Name":
			p.Name = val
		case "Uid":
			uid, _ := strconv.Atoi(strings.TrimSpace(strings.Split(val, "\t")[0]))
			p.User = lookupUser(uid)
		case "State":
			p.State = stateChar(val[0])
		case "VmRSS":
			// kB
			if n, err := strconv.Atoi(strings.Fields(val)[0]); err == nil {
				p.Mem = float64(n) / 1024.0 // MB
			}
		}
	}
}

func (p *Process) readStat() {
	path := fmt.Sprintf("/proc/%d/stat", p.PID)
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	// Format: pid (comm) state ppid ...
	// comm can contain spaces and parens, so find last ')'
	lastParen := strings.LastIndex(string(data), ")")
	if lastParen == -1 {
		return
	}

	fields := strings.Fields(string(data[lastParen+2:]))
	// fields[0] = state, fields[1] = ppid, fields[11] = utime, fields[12] = stime
	if len(fields) < 2 {
		return
	}
	p.PPID, _ = strconv.Atoi(fields[1])

	if len(fields) >= 13 {
		utime, _ := strconv.Atoi(fields[11])
		stime, _ := strconv.Atoi(fields[12])
		p.CPU = float64(utime+stime) / 100.0 // rough CPU seconds
	}
}

func buildTree(procs []*Process, filterUser string) []*Process {
	childMap := make(map[int][]*Process)
	var roots []*Process

	for _, p := range procs {
		if filterUser != "" && p.User != filterUser {
			continue
		}
		if p.PPID == 0 {
			roots = append(roots, p)
		} else {
			childMap[p.PPID] = append(childMap[p.PPID], p)
		}
	}

	for _, p := range procs {
		if filterUser != "" && p.User != filterUser {
			continue
		}
		if children, ok := childMap[p.PID]; ok {
			sort.Slice(children, func(i, j int) bool {
				return children[i].PID < children[j].PID
			})
			p.Children = children
		}
	}

	sort.Slice(roots, func(i, j int) bool {
		return roots[i].PID < roots[j].PID
	})

	return roots
}

func findPID(roots []*Process, pid int) *Process {
	for _, p := range roots {
		if p.PID == pid {
			return p
		}
		if child := findPID(p.Children, pid); child != nil {
			return child
		}
	}
	return nil
}

func printSubtree(p *Process, depth int, isLast bool, showAll bool) {
	if p == nil {
		return
	}

	prefix := ""
	if depth > 0 {
		prefix += "│   "
	}
	if isLast {
		prefix += "└── "
	} else {
		prefix += "├── "
	}

	stateColor := stateColorCode(p.State)
	name := fmt.Sprintf("\033[1;37m%s\033[0m", p.Name)
	pidStr := fmt.Sprintf("\033[90m(%d)\033[0m", p.PID)

	line := fmt.Sprintf("%s%s %s", prefix, name, pidStr)
	if showAll {
		line += fmt.Sprintf(" \033[90muser=%s\033[0m", p.User)
		line += fmt.Sprintf(" \033[90mmem=%.1fMB\033[0m", p.Mem)
	}
	fmt.Println(line)

	for i, child := range p.Children {
		printSubtree(child, depth+1, i == len(p.Children)-1, showAll)
	}
}

func printFlat(roots []*Process, showAll bool) {
	printFlatRec(roots, showAll)
}

func printFlatRec(procs []*Process, showAll bool) {
	for _, p := range procs {
		line := fmt.Sprintf("%7d %7d \033[1;37m%-20s\033[0m %s", p.PID, p.PPID, p.Name, p.User)
		if showAll {
			line += fmt.Sprintf(" \033[90m%.1fMB\033[0m", p.Mem)
		}
		fmt.Println(line)
		printFlatRec(p.Children, showAll)
	}
}

func stateChar(c byte) string {
	switch c {
	case 'R':
		return "running"
	case 'S':
		return "sleeping"
	case 'D':
		return "disk_sleep"
	case 'T':
		return "stopped"
	case 'Z':
		return "zombie"
	case 'X':
		return "dead"
	default:
		return string(c)
	}
}

func stateColorCode(state string) string {
	switch state {
	case "running":
		return "\033[92m"
	case "sleeping":
		return "\033[90m"
	case "zombie":
		return "\033[91m"
	case "stopped":
		return "\033[93m"
	default:
		return "\033[90m"
	}
}

func lookupUser(uid int) string {
	// Read /etc/passwd
	data, err := os.ReadFile("/etc/passwd")
	if err != nil {
		return strconv.Itoa(uid)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, ":", 4)
		if len(parts) < 4 {
			continue
		}
		id, _ := strconv.Atoi(parts[2])
		if id == uid {
			return parts[0]
		}
	}
	return strconv.Itoa(uid)
}

func usage() {
	fmt.Println(`proc-tree-cli - Display the process tree from /proc

Usage:
  proc-tree-cli [OPTIONS]

Options:
  -u, --user USER   Filter by username or UID
  -p, --pid PID     Show subtree for a specific PID
  -a, --all         Show extra info (user, memory)
  -f, --format FMT  Output format: tree (default) or flat
  -h, --help        Show this help message

Examples:
  proc-tree-cli
  proc-tree-cli -a
  proc-tree-cli -u root
  proc-tree-cli -p 1 -a
  proc-tree-cli -f flat`)
}

// Ensure filepath is used
var _ = filepath.Join
