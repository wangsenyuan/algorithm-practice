package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 2
2 1
100 100
100 100`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1 2 1 2
3 2 1 2
1 2 1 1
1 3 1 2
1 2 3 4
5 6 7 8`
	expect := 14
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1 2 2
2 2 1
2 1 1
100 100 100
100 100 100`
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
8 7 2 8 4 8
7 7 9 7 1 1
8 3 1 1 8 5
6 8 3 1 1 4
1 4 5 1 9 6
7 1 1 6 8 2
11 23 20 79 30 15
15 83 73 57 34 63`
	expect := 183
	runSample(t, s, expect)
}
