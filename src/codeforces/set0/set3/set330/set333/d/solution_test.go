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
	s := `2 2
1 2
3 4
1`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 0 0
0 1 1
1 0 0
0`
	runSample(t, s)
}
