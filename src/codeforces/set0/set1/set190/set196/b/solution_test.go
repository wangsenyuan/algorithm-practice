package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readString(reader)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 4
##.#
##S#
#..#
#.##
#..#
Yes
`)
}
func TestSample2(t *testing.T) {
	runSample(t, `5 4
##.#
##S#
#..#
..#.
#.##
No
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 4
##.#
.#S.
..##
#.#.
#..#
Yes
`)
}

func TestSample4(t *testing.T) {
	runSample(t, `6 9
#.#.#.#.#
S.#.#.#.#
###.#.###
#...#...#
#.#####.#
#.#...#.#
No
`)
}
