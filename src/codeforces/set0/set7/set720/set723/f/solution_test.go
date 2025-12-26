package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, x, y, ds, dt, ok, res := drive(reader)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !expect {
		return
	}

	if len(res) != n-1 {
		t.Fatalf("Sample result %v, not a valid tree", res)
	}

	deg := make([]int, n+1)

	set := NewDSU(n + 1)

	for _, e := range res {
		u, v := e[0], e[1]
		deg[u]++
		deg[v]++
		if !set.Union(u, v) {
			t.Fatalf("Sample result %v, not a valid tree", res)
		}
	}

	if deg[x] > ds || deg[y] > dt {
		t.Fatalf("Sample result %v, not meeting the condition deg of x = %d, y = %d", res, deg[x], deg[y])
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 2
2 3
3 1
1 2 1 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 8
7 4
1 3
5 4
5 7
3 2
2 4
6 1
1 2
6 4 1 4
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 15
4 1
5 10
2 1
5 7
9 2
4 6
6 7
9 1
6 9
8 4
8 3
9 8
3 9
2 3
7 10
10 1 2 3
`
	expect := true
	runSample(t, s, expect)
}
