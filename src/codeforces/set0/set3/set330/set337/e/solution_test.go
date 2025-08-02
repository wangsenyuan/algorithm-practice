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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
6 10
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
4 8 72 6
`
	expect := 12
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
10000000000 10000000001 10000000002 10000000003
`
	expect := 33
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
9 6 54
`
	expect := 7
	runSample(t, s, expect)
}
