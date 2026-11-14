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
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
0 2 3
`, -3)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
0 1000000000
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
0
`, 0)
}
