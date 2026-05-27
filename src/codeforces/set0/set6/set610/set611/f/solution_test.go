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
	runSample(t, `1 10 2
R
`, 30)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 4 6
RUL
`, 134)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 1 500000
RLRL
`, -1)
}
