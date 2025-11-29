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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
15 -1 -1
10 1 3
5 -1 -1`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8
6 2 3
3 4 5
12 6 7
1 -1 8
4 -1 -1
5 -1 -1
14 -1 -1
2 -1 -1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
5 2 -1
3 3 -1
2 -1 4
4 -1 -1`
	expect := 1
	runSample(t, s, expect)
}
