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
	s := `2 2
0010
0100`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4
001000
000010
000010`
	expect := 12
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 3
01110
01110
01110
01110`
	expect := 18
	runSample(t, s, expect)
}
