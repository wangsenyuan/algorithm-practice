package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	got := drive(reader)
	if got != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, got)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
3 1 2 2 3
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
3 1 2 1 3
4 1 2 2 3 2 4
`, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, `2
5 1 2 2 3 3 4 3 5
7 3 4 1 2 2 4 4 6 2 7 6 5
`, 7)
}
