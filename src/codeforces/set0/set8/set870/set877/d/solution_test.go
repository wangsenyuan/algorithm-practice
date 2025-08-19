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
	s := `3 4 4
....
###.
....
1 1 3 1
`
	expect := 3

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4 1
....
###.
....
1 1 3 1
`
	expect := 8

	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2 1
.#
#.
1 1 2 2
`
	expect := -1

	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 10 3
...##..#..
#.#..#...#
..#.##.#..
##..#..#.#
..#...##..
.#.#.#....
#......#..
.#.####.##
......#...
.#.##...##
4 6 8 1
`
	expect := 7

	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10 10 10
...#......
##.#.#####
...#...###
.#####.###
...#...###
##.#.#####
...#...###
.#####.###
.......###
##########
1 1 1 10
`
	expect := 17

	runSample(t, s, expect)
}
