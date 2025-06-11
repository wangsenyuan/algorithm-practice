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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
4 7
4 7
7 4
0`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5
4 7
7 4
2 11
9 7
1 1
2`
	runSample(t, s)
}
