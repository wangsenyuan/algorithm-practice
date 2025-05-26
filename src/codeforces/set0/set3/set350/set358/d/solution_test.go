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
	s := `4
1 2 3 4
4 3 2 1
0 1 1 0
13`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3
1 1 1
1 2 1
1 1 1
4`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `7
8 5 7 6 1 8 9
2 7 9 5 4 3 1
2 3 3 4 1 1 3
44`
	runSample(t, s)
}
