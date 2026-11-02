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
	runSample(t, `3 1
2 2 3 3
`, 11)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 2
1 2 2 3
2 1 3 2
`, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 3
1 1 2 4
3 1 4 2
3 3 4 4
`, 0)
}
