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
	s := `4 4
XX..
XX..
XXXX
XXXX
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4
....
.XXX
.XXX
....
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 5
XXXX.
XXXX.
.XX..
.XX..
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 6
.XXX..
.XXXX.
.XXXX.
..XXX.
..XXX.
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1 1
X
`
	expect := 1
	runSample(t, s, expect)
}
