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
	runSample(t, `3 1 8 10
10 8
5 7
11 9
3
10`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 2 10 18
10 4
20 6
5 3
20`)
}
