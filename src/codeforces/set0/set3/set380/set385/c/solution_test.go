package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, x := range res {
		y := readNum(reader)
		if x != y {
			t.Fatalf("Sample expect %d, but got %d", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6
5 5 7 10 14 15
3
2 11
3 12
4 4
9
7
0`

	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `7
2 3 5 7 11 4 8
2
8 10
2 123
0
7`

	runSample(t, s)
}
