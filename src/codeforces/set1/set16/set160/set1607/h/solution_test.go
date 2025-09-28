package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	dishes, best, ans := drive(reader)

	if expect != best {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	type pair struct {
		first  int
		second int
	}
	freq := make(map[pair]int)

	for i, cur := range dishes {
		x, y := ans[i][0], ans[i][1]
		if x+y != cur[2] || x > cur[0] || y > cur[1] {
			t.Fatalf("Sample result %v, not correct", ans[i])
		}
		x = cur[0] - x
		y = cur[1] - y
		freq[pair{x, y}]++
	}

	if len(freq) != expect {
		t.Fatalf("Sample result %v, not correct", freq)
	}
}

func TestSample1(t *testing.T) {
	s := `3
10 10 2
9 9 0
10 9 1`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
3 4 1
5 1 2`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
7 2 5
6 5 4
5 5 6`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1
13 42 50`
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5
5 7 12
3 1 4
7 3 7
0 0 0
4 1 5`
	expect := 2
	runSample(t, s, expect)
}
