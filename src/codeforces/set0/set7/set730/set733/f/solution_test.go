package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, S, W, C, edges, best, res := drive(reader)

	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	set := NewDSU(n)

	var sum int
	var use int
	for _, cur := range res {
		x, vx := cur[0]-1, cur[1]
		u, v := edges[x][0]-1, edges[x][1]-1
		if !set.Union(u, v) {
			t.Fatalf("Sample result %v, not correct, it has cycle at %v", res, cur)
		}
		sum += vx
		use += (W[x] - vx) * C[x]
	}

	if sum != best || use > S {
		t.Fatalf("Sample result %v, not correct, sum = %d, use = %d, best = %d, S = %d", res, sum, use, best, S)
	}
}

func TestSample1(t *testing.T) {
	s := `6 9
1 3 1 1 3 1 2 2 2
4 1 4 2 2 5 3 1 6
1 2
1 3
2 3
2 4
2 5
3 5
3 6
4 5
5 6
7
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
9 5 1
7 7 2
2 1
3 1
3 2
2
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 6
8 10 4 8 4 4
45 51 13 13 37 26
2 7
7 1
6 3
3 1
5 4
1 4
4
`
	expect := 38
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `9 8
2 8 2 10 2 2 5 8
31 29 17 16 31 22 10 13
5 2
8 2
7 3
3 1
6 8
4 7
1 2
9 2
69
`
	expect := 33
	runSample(t, s, expect)
}
