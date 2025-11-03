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
	s := `2 2 3
000
000

111
111
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3 3
111
111
111

111
111
111

111
111
111
`
	expect := 19
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1 10
0101010101
`
	expect := 0
	runSample(t, s, expect)
}
