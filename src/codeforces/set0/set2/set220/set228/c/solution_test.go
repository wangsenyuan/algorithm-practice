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
	s := `6 11
......*.***
*.*.*....**
.***....*.*
..***.*....
.*.*.....**
......*.*..
3`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 4
..**
..**
....
....
0`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `10 10
*.........
.*........
...**.....
..........
..*.....*.
.*.....*..
....*.*.*.
*.......*.
***...*...
..........
2`
	runSample(t, s)
}
