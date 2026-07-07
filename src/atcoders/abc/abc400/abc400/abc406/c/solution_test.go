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
	runSample(t, `6
1 3 6 4 2 5
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
1 2 3 4 5 6
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `12
11 3 8 9 5 2 10 4 1 6 12 7
`, 4)
}
