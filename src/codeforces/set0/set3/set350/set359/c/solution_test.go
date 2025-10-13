package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 2
2 2
`, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3
1 2 3
`, 27)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 2
29 29
`, 73741817)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 5
0 0 0 0
`, 1)
}
