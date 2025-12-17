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
	s := `6 5 1 2
##.#.
.###.
###..
#...#
.##.#
###..
`
	expect := 11
	runSample(t, s, expect)
}
func TestSample2(t *testing.T) {
	s := `2 5 1 1
#####
.....
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 6 1 3
......
......
......
......
......
......
`
	expect := 6
	runSample(t, s, expect)
}
