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
	runSample(t, `3 2 3
1 2 3
1 2
2 3
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 2 2
1 1 2
1 2
2 1
0`)
}

func TestSample3(t *testing.T) {
	runSample(t, `10 3 2
2 1 1 2 1 1 2 1 2 2
4 10
9 3
5 7
2`)
}
