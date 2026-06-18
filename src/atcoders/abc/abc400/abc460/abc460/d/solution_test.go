package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
#.#.
.#..
#...
`
	expect := []string{
		"#.#.",
		".#..",
		"#..#",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
###
###
###
`
	expect := []string{
		"...",
		"...",
		"...",
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 7
.#.....
.......
..#....
.......
....#..
`
	expect := []string{
		".#.##.#",
		"....#..",
		"#.#.###",
		"#.....#",
		"###.#.#",
	}
	runSample(t, s, expect)
}
