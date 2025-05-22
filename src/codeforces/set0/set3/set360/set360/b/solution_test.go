package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	expect := readNum(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `5 2
4 7 4 7 4
0`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 1
-100 0 100
100`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `6 3
1 2 3 7 8 9
1`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `20 17
-5 -9 11 -7 -17 -8 0 -14 -20 -15 7 -13 0 -3 -14 0 9 -10 6 -19
0`
	runSample(t, s)
}
