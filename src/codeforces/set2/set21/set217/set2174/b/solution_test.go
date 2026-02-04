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
	runSample(t, `3 4
0 0 1
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `6 8
1 2 0 5 1 8
`, 20)
}

func TestSample3(t *testing.T) {
	//
	runSample(t, `3 4
1 0 4
`, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, `5 8
2 4 5 4 3
`, 19)
}
