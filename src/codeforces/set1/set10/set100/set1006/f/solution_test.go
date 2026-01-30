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
	s := `3 3 11
2 1 5
7 10 0
12 6 4
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4 2
1 3 3 3
0 3 3 2
3 0 1 1
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 4 1000000000000000000
1 3 3 3
0 3 3 2
3 0 1 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 1 1000000000000000000
1000000000000000000
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1 2 3
1 2
`
	expect := 1
	runSample(t, s, expect)
}

