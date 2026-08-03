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
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 0
2 3 4
`, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 0
1000000000 999999999
`, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, `6 0
2 5 3 4 1 6
`, 4)
}
