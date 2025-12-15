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
	s := `7 2 11
3 4 3 5 1 4 6
7 7 3 6 5 3 9
`
	expect := 12
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 4 20
5 6 4 3 7 5 4 1
10 12 5 12 14 8 5 8
`
	expect := 19
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1 5
6
9
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 1 3
4
7
`
	expect := 0
	runSample(t, s, expect)
}
