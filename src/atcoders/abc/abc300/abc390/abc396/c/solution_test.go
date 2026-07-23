package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 3
8 5 -1 3
3 -2 -4
`, 19)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 3
5 -10 -2 -5
8 1 4
`, 15)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 5
-36 -33 -31
12 12 28 24 27
`, 0)
}
