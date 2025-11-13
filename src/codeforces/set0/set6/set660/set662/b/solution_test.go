package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, best, res := drive(reader)
	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}
	if expect == -1 {
		return
	}
	pressed := make([]bool, n)
	for _, u := range res {
		pressed[u-1] = true
	}

	color := -1
	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		if pressed[u] {
			w ^= 1
		}
		if pressed[v] {
			w ^= 1
		}
		if color == -1 {
			color = w
		} else if color != w {
			t.Fatalf("Sample result %v, not correct, it leads different edge colors", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 2 B
3 1 R
3 2 B
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 5
1 3 R
2 3 R
3 4 B
4 5 R
4 6 R
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 5
1 2 R
1 3 R
2 3 B
3 4 B
1 4 B
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 6
1 2 R
1 3 R
2 3 R
4 5 B
4 6 B
5 6 B
`
	expect := -1
	runSample(t, s, expect)
}
