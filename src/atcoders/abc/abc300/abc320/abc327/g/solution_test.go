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
`, 36)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3
`, 168)
}

func TestSample3(t *testing.T) {
	runSample(t, `12 34
`, 539029838)
}

func TestSample4(t *testing.T) {
	runSample(t, `20 231104
`, 966200489)
}

func TestSmallCases(t *testing.T) {
	runSample(t, `1 1
`, 0)
	runSample(t, `2 1
`, 2)
	runSample(t, `2 2
`, 4)
	runSample(t, `3 1
`, 6)
	runSample(t, `4 3
`, 1536)
}
