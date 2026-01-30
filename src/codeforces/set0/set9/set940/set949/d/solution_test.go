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
	s := `5 1 1
1 0 0 0 4
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 1 2
3 8 0 1 0 0
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `100 66 30
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 25 27 15 53 29 56 30 24 50 39 39 46 4 14 44 16 55 48 15 54 25 4 44 15 29 55 22 49 52 9 2 22 15 14 33 24 38 11 48 27 34 29 8 37 47 36 54 45 24 31 1434
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 1 1
1 1 0 3 0
`
	expect := 0
	runSample(t, s, expect)
}
