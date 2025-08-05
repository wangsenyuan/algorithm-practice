package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7 12
############
#S........T#
#.########.#
#..........#
#..........#
#..#..#....#
############
`
	expect := 15
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 8
########
#......#
#.####.#
#...T#S#
########
`
	expect := 14
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 8
########
#.#S...#
#.####.#
#...T#.#
########
`
	expect := -1
	runSample(t, s, expect)
}
