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
	runSample(t, `3 2
1 2
2 3
`, 4052)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 6
1 2
2 3
1 4
2 4
1 3
3 4
`, 4052)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 1
1 2
`, 2026)
}

func TestNoEdges(t *testing.T) {
	runSample(t, `2 0
`, 4052)
}
