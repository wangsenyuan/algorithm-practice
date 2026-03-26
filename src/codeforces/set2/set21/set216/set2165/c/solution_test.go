package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	for i, x := range res {
		var y int
		fmt.Fscan(reader, &y)
		if x != y {
			t.Fatalf("Sample expect %d, but got %d at %d-th", y, x, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 1
5 7
9
1`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 1
9 9 8
24
7`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `6 4
1 1 4 5 1 4
10
20
30
40
3
11
16
31`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `1 1
0
0
0`
	runSample(t, s)
}
