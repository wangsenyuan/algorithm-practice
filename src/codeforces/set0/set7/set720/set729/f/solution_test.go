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
	runSample(t, `3
1 3 1
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
-1 -2 -1 -2 -1
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
-4 -2 4 5
`, -13)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
0
`, 0)
}
