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
	runSample(t, `3 1
7 2 3
1 3 2
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 2
7 2 3
1 3 2
`, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 1
2 1
1 2
`, 0)
}
