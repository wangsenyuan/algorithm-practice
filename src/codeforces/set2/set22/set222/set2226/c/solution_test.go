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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
0 1 2 3
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
6 7
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
8 1 7 6 4 3
`, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, `9
9 8 2 4 4 3 5 3 4
`, 6)
}
