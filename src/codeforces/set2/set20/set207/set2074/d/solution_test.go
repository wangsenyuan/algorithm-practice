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
	runSample(t, `2 3
0 0
1 2`, 13)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 3
0 2
1 2`, 16)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 3
0 2 5
1 1 1`, 14)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 8
0 5 10 15
2 2 2 2`, 52)
}
