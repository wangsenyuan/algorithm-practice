package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	h, w, coins, best, res := drive(reader)
	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	if len(res) != h-1+w-1 {
		t.Fatalf("Sample expect %d moves, but got %d", h-1+w-1, len(res))
	}

	// no move
	points := make(map[pair]int)
	for _, cur := range coins {
		r, c := cur[0]-1, cur[1]-1
		points[pair{r, c}]++
	}

	var x, y int

	var cnt int
	for _, c := range res {
		if c == 'R' {
			y++
		} else {
			x++
		}
		if x == h || y == w {
			t.Fatalf("move out of bounds at %d %d", x, y)
		}
		cnt += points[pair{x, y}]
	}

	if cnt != best {
		t.Fatalf("Sample expect %d points, but got %d", best, cnt)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4 4
3 3
2 1
2 3
1 4
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2 2
2 1
1 2
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 15 8
2 7
2 9
7 9
10 3
7 11
8 12
9 6
8 1
`
	expect := 5
	runSample(t, s, expect)
}