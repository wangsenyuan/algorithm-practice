package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, d, edges, ok, res := drive(reader)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}

	if !expect {
		return
	}

	deg := make([]int, n)

	for i, e := range edges {
		j := sort.SearchInts(res, i+1)
		if j < len(res) && res[j] == i+1 {
			u, v := e[0]-1, e[1]-1
			deg[u]++
			deg[v]++
		}
	}

	for i := range n {
		if d[i] >= 0 && deg[i]%2 != d[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `1 0
1
`
	expect := false
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 5
0 0 0 -1
1 2
2 3
3 4
1 4
2 4
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 1
1 1
1 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 3
0 -1 1
1 2
2 3
1 3
`
	expect := true
	runSample(t, s, expect)
}
