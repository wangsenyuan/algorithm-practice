package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	best, _ := process(reader)
	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
1 2 3 1 2
	`, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 -2 3 1 -2
	`, 5)
}
