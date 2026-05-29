// go-build-check verifies a Go project builds cleanly and optionally checks
// cross-platform builds (GOOS/GOARCH) without producing binaries.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Target struct {
	GOOS   string
	GOARCH string
}

var commonTargets = []Target{
	{"linux", "amd64"},
	{"linux", "arm64"},
	{"darwin", "amd64"},
	{"darwin", "arm64"},
	{"windows", "amd64"},
	{"windows", "arm64"},
}

func main() {
	cross := flag.Bool("cross", false, "check cross-platform builds")
	packages := flag.String("pkg", "./...", "packages to build")
	timeout := flag.String("timeout", "10m", "build timeout per target")
	_ = timeout // reserved for future use
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()

	var targets []Target
	if *cross {
		targets = commonTargets
		// Add current platform if not already included
		current := Target{runtime.GOOS, runtime.GOARCH}
		found := false
		for _, t := range targets {
			if t.GOOS == current.GOOS && t.GOARCH == current.GOARCH {
				found = true
				break
			}
		}
		if !found {
			targets = append(targets, current)
		}
	} else {
		targets = []Target{{runtime.GOOS, runtime.GOARCH}}
	}

	failed := false
	for _, target := range targets {
		isCurrent := target.GOOS == runtime.GOOS && target.GOARCH == runtime.GOARCH
		label := fmt.Sprintf("%s/%s", target.GOOS, target.GOARCH)
		if isCurrent {
			label += " (current)"
		}

		fmt.Printf("Checking %s... ", label)

		cmd := exec.Command("go", "build", "-o", "/dev/null", *packages)
		cmd.Env = append(os.Environ(),
			"GOOS="+target.GOOS,
			"GOARCH="+target.GOARCH,
		)

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("FAIL")
			if *verbose {
				fmt.Printf("  Error: %v\n", err)
				fmt.Printf("  Output:\n%s\n", string(output))
			}
			failed = true
		} else {
			fmt.Println("OK")
		}
	}

	fmt.Println()
	if failed {
		fmt.Println("Build check FAILED on one or more targets.")
		os.Exit(1)
	}
	fmt.Println("All targets build successfully.")
}

// CheckBuild verifies a Go project builds for a specific target.
func CheckBuild(target Target, packages string) error {
	cmd := exec.Command("go", "build", "-o", "/dev/null", packages)
	cmd.Env = append(os.Environ(),
		"GOOS="+target.GOOS,
		"GOARCH="+target.GOARCH,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("build failed for %s/%s: %w\n%s", target.GOOS, target.GOARCH, err, string(output))
	}
	return nil
}

// ListTargets returns the list of common cross-platform build targets.
func ListTargets() []Target {
	return commonTargets
}

// FormatTarget returns a formatted string for a build target.
func FormatTarget(target Target) string {
	return fmt.Sprintf("%s/%s", target.GOOS, target.GOARCH)
}

// IsCurrentTarget checks if the given target matches the current platform.
func IsCurrentTarget(target Target) bool {
	return target.GOOS == runtime.GOOS && target.GOARCH == runtime.GOARCH
}

// String returns a string representation of the target.
func (t Target) String() string {
	return FormatTarget(t)
}

// ParseTarget parses a "GOOS/GOARCH" string into a Target.
func ParseTarget(s string) (Target, error) {
	parts := strings.SplitN(s, "/", 2)
	if len(parts) != 2 {
		return Target{}, fmt.Errorf("invalid target format %q, expected GOOS/GOARCH", s)
	}
	return Target{GOOS: parts[0], GOARCH: parts[1]}, nil
}
