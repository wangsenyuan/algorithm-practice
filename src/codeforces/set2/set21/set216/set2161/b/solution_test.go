package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1
.
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1
#
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
.##
.##
...
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
#..
.#.
..#
`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3
###
...
...
`
	expect := false
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `3
#.#
...
.#.
`
	expect := false
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `4
####
#..#
#..#
####
`
	expect := false
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `3
..#
...
.#.
`
	expect := true
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `3
..#
#..
...
`
	expect := true
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := `5
#.#.#
.#.#.
#.#.#
.#.#.
#.#.#
`
	expect := false
	runSample(t, s, expect)
}

func TestSample11(t *testing.T) {
	s := `5
...#.
...#.
.....
##...
.....
`
	expect := true
	runSample(t, s, expect)
}