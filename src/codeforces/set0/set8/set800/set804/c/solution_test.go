package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	w, tot, assign := drive(reader)
	if tot != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, tot)
	}

	occ := make([]bool, tot+1)
	for _, v := range w {
		for _, i := range v {
			i--
			if occ[assign[i]-1] {
				t.Fatalf("Sample result %v, is not valid", assign)
			}
			occ[assign[i]-1] = true
		}
		for _, i := range v {
			i--
			occ[assign[i]-1] = false
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 1
2 2 3
1 2
1 2
2 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 5
0
1 1
1 3
3 2 4 5
2 1
3 2
4 3
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 35
3 17 20 32
4 3 14 24 25
4 4 10 17 26
7 2 9 13 17 23 28 30
9 1 2 7 8 13 16 18 33 35
8 5 6 11 15 17 22 29 34
5 12 19 21 27 31
2 1
3 1
4 3
5 4
6 3
7 4
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 3
0
0
0
1 2
2 3
`
	expect := 1
	runSample(t, s, expect)
}
