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
1 1 1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8
8 7 6 5 4 3 2 1
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `24
1 8 1 2 8 2 3 8 3 4 8 4 5 8 5 6 8 6 7 8 7 8 8 8
`
	expect := 17
	runSample(t, s, expect)
}
