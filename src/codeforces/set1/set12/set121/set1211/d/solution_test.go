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
	runSample(t, `12 1 2 2
1 1 2 2 2 2 2 3 3 4 6 6
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `14 1 1 3
3 3 1 1 9 9 2 3 6 6 3 18 3 18
`, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, `1 2 3 10
1000000
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `7 2 1 2
1 1 1 2 4 4 4
`, 1)
}
