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
	s := `5 2
0 1 7 2 4
8`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5 3
0 1 7 2 4
9`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `1 1
3
2`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `3 0
2 0 3
3`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `1 100000000000
0
36`
	runSample(t, s)
}
