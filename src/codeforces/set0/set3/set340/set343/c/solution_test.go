package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNum(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 4
2 5 6
1 3 6 8
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3
1 2 3
1 2 3
0`)
}

func TestSample3(t *testing.T) {
	runSample(t, `1 2
165
142 200
81`)
}

func TestSample4(t *testing.T) {
	runSample(t, `1 2
5000000000
1 10000000000
14999999998`)
}
