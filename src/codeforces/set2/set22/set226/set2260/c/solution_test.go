package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expectXor, expectOps int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	gotXor, gotOps := drive(reader)
	if gotXor != expectXor || gotOps != expectOps {
		t.Fatalf("Sample expect %d %d, but got %d %d", expectXor, expectOps, gotXor, gotOps)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 1
`, 4, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `0 5
`, 5, 0)
}

func TestSample3(t *testing.T) {
	// 1010
	// 0100
	// 
	runSample(t, `6 4
`, 10, 4)
}
