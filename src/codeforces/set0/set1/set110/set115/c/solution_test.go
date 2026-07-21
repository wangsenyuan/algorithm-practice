package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 2
13
..
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 1
1
4
.
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 2
3.
.1
`, 1)
}
