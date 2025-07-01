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
	runSample(t, `4 0
1 2 3 4
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 2
7 6 6 7 1
2`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 1
6 7 9
4`)
}

func TestSample4(t *testing.T) {
	runSample(t, `6 2
5 1 9 10 13 2
9`)
}
