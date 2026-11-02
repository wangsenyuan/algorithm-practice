package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6
5 1 4 2 2 -4
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `7
0 0 0 0 0 0 0
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `12
-4 3 0 -4 -4 -1 -4 1 -3 -2 -2 -1
`, 5)
}
