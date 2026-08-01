package main

import (
	"os/exec"
	"strings"
	"testing"
)

func run(t *testing.T, stdin string, args ...string) (string, int) {
	t.Helper()
	cmd := exec.Command("go", append([]string{"run", "."}, args...)...)
	cmd.Stdin = strings.NewReader(stdin)
	out, err := cmd.CombinedOutput()
	exit := 0
	if ee, ok := err.(*exec.ExitError); ok {
		exit = ee.ExitCode()
	} else if err != nil {
		t.Fatalf("go run: %v\n%s", err, out)
	}
	return string(out), exit
}

func TestGoodAnchor(t *testing.T) {
	md := "# Hello World\n\nText [go](#hello-world)\n"
	out, exit := run(t, md)
	if exit != 0 {
		t.Fatalf("exit %d, out=%s", exit, out)
	}
	if !strings.Contains(out, "all anchors resolve") {
		t.Fatal(out)
	}
}

func TestBrokenAnchor(t *testing.T) {
	md := "# Hello\n\nText [go](#missing)\n"
	out, exit := run(t, md)
	if exit != 1 {
		t.Fatalf("exit %d, out=%s", exit, out)
	}
	if !strings.Contains(out, "broken") {
		t.Fatal(out)
	}
}

func TestSlugify(t *testing.T) {
	cases := map[string]string{
		"Hello World":            "hello-world",
		"Café & Beer":            "caf-beer",
		"Multiple  spaces---together": "multiple-spaces-together",
		"Under_score":            "under_score",
	}
	for in, want := range cases {
		if got := slugify(in); got != want {
			t.Errorf("slugify(%q) = %q, want %q", in, got, want)
		}
	}
}
