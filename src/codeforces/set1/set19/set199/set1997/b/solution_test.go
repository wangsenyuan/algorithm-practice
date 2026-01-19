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
	s := `8
.......x
.x.xx...
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
..
..
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
xxx
xxx
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `9
..x.x.x.x
x.......x
`
	expect := 2
	runSample(t, s, expect)
}
