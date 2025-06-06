package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	expect := readNum(reader)
	if ans != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3 1
.#.
###
.#.
3`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2 5 3
###..
...##
4`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `4 4 2
..##
###.
#..#
####
5`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `2 1 3
#
#
2`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `4 7 2
#.##.##
.##.#..
#####..
##.#.##
6`
	runSample(t, s)
}

func TestSample6(t *testing.T) {
	s := `2 6 5
##...#
.#.###
7`
	runSample(t, s)
}

func TestSample8(t *testing.T) {
	s := `4 10 14
#....#..##
.##.##.###
##....###.
##..#.#..#
21`
	runSample(t, s)
}
