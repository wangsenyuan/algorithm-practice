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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 11
2 42
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 7
2 8 16 183
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `12 13
80 68 862370 82217 8 56 5 168 672624 6 286057 11864
`
	expect := 10
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 5
1000000000 1000000000 1000000000 1000000000 1000000000
`
	expect := 25
	runSample(t, s, expect)
}
