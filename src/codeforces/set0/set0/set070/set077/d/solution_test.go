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
	s := `3 4
#################
#O..#...#O.O#...#
#.O.#.O.#.O.#...#
#..O#...#O.O#...#
#################
#O.O#OOO#O.O#...#
#.O.#...#...#.O.#
#O.O#OOO#O.O#...#
#################
#O.O#...#O.O#...#
#...#...#...#.O.#
#O.O#...#O.O#...#
#################
3`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2 2
#########
#O.O#O.O#
#.O.#...#
#O.O#O.O#
#########
#...#O.O#
#...#...#
#...#O.O#
#########
2`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `2 2
#########
#..O#O..#
#...#...#
#O..#..O#
#########
#O..#..O#
#...#...#
#..O#O..#
#########
0`
	runSample(t, s)
}
