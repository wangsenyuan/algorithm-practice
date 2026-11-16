package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	got := drive(bufio.NewReader(strings.NewReader(s)))
	if got != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, got)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
1 1 1
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 2 3
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
1 2 2
`, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, `4
1 1 1 1
`, 4)
}

func TestSample5(t *testing.T) {
	runSample(t, `5
1 2 3 4 5
`, 0)
}

func TestSample7(t *testing.T) {
	runSample(t, `10
1 1 1 1 1 1 1 1 1 1
`, 120)
}
