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
	if len(res) != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
0 0 0 0 1
0 1 2 3 4
5`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5
0 0 1 0 1
0 1 2 2 4
2`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `4
1 0 0 0
2 3 4 2
1`
	runSample(t, s)
}
