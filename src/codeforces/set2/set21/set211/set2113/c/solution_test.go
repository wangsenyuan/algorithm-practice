package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNum(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 3 1
#.#
g.g
2`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2 3 2
#.#
g.g
0`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3 4 2
.gg.
g..#
g##.
4`
	runSample(t, s)
}
