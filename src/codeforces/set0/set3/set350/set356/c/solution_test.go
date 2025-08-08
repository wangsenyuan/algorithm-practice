package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
1 2 2 4 3
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
4 1 1
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
0 3 0 4
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `20
4 1 4 4 2 1 4 3 2 3 1 1 2 2 2 4 4 2 4 2
`, 6)
}

func TestSample5(t *testing.T) {
	runSample(t, `5
4 4 3 3 1
`, 1)
}
