package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 18 2
`, "3")
}

func TestSample2(t *testing.T) {
	runSample(t, `-1 8 3
`, "-2")
}

func TestSample4(t *testing.T) {
	runSample(t, `1 16 5
`, "No solution")
}
