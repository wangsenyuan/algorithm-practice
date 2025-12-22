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
	s := `5 2 3
0 12 3 11 7 -2
`
	expect := 11
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 2 2
0 1 -1 1 -1 1 -1
`
	expect := -3
	runSample(t, s, expect)
}
