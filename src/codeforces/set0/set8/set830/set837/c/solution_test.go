package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2 2
1 2
2 1
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 10 9
2 3
1 1
5 10
9 11
`
	expect := 56
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 10 10
6 6
7 7
20 5
`
	expect := 0
	runSample(t, s, expect)
}
