package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
0 2 7
2 0 9
7 9 0`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 2 7
2 0 9
7 9 0
`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
0 2 2
7 0 9
7 9 0
`, false)
}

func TestSample4(t *testing.T) {
	runSample(t, `3
0 1 1
1 0 1
1 1 0
`, false)
}

func TestSample5(t *testing.T) {
	runSample(t, `2
0 0
0 0
`, false)
}
