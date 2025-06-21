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
	runSample(t, `3 6 7
2 13
4 20
7 9
3`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 4 100
10 70
15 76
21 12
30 100
67 85
26`)
}
