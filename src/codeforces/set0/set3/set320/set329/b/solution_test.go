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
	s := `5 7
000E0T3
T0TT0T0
010T0T0
2T0T0T0
0T0S000
3`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `1 4
SE23
2`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `1 10
9T9TSET9T9
0`
	runSample(t, s)
}
