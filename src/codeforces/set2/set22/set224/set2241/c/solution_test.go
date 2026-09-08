package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
0000
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
110
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
110011
`, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, `6
101100
`, 1)
}
