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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 5
##...
..1#@
432#.
...#.
4	
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 4
#78#
.612
.543
..@.
6	
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3 2
3@
2#
1#
-1
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `15 15
@..............
...............
...............
...............
...............
...............
...............
...............
...............
...............
...............
.............9.
.............8.
...........567.
...........4321
28
`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `10 10
....#.....
........#.
......#...
..........
..........
..........
.........#
..1..#...#
..2...@.#.
..3#...#..
5
`
	runSample(t, s)
}
