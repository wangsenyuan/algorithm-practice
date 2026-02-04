package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3
.###.
..#..
#.#.#
#...#
##..#`
	expect := "10111"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
##
..`
	expect := "11"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 3
##.##.##.#
.####..#..
...#.#..#.
.#.#.#.#..
...####...
#.#.##....
.##...#...
#.##.....#
#....###.#
.#..#.#...`
	expect := "0011010010"
	runSample(t, s, expect)
}

