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
	runSample(t, `5 2
2 1 3 1 4
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `9 3
1 1 1 1 1 1 1 1 1
`, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 3
2 8 5 7
`, 0)
}
