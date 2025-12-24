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
	s := `3 4
6 2 8
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 6
9 1 14
`
	expect := 1
	runSample(t, s, expect)
}
