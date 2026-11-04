package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 2 2
2 4
6 8
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 2 7
6 7
`, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 2 1
5 7
1 2
5 100
`, 104)
}

func TestSample4(t *testing.T) {
	runSample(t, `7 4 5
	7 7 7 12
	7 12 12 7
	7 7 7 7
	7 7 12 7
	7 7 12 12
	12 12 7 12
	7 7 7 7
`, 9)
}
