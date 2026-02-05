package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	p, res := drive(reader)

	n := len(p)

	if len(res) == n-1 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	pos := make([]int, n+1)
	for i, v := range p {
		pos[v] = i
	}

	fa := make([]int, n+1)
	for i := range n + 1 {
		fa[i] = i
	}

	find := func(x int) int {
		y := x
		for fa[y] != y {
			y = fa[y]
		}

		for x != y {
			fa[x], x = y, fa[x]
		}

		return y
	}

	for _, e := range res {
		u, v := e[0], e[1]
		if u > v {
			u, v = v, u
		}
		if pos[u] > pos[v] {
			t.Fatalf("Sample result %v, not correct, u < v, but pos[u] = %d, pos[v] = %d", res, u, v)
		}
		u = find(u)
		v = find(v)
		if u != v {
			fa[v] = u
		}
	}

	for i := 1; i <= n; i++ {
		if find(i) != find(1) {
			t.Fatalf("Sample result %v, not correct, i = %d, find(i) = %d, find(1) = %d", res, i, find(i), find(1))
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6
1 3 4 5 2 6
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
3 4 1 2
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
4 3 5 1 2
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
1 2 3 4
`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `7
4 3 5 7 6 2 1
`
	expect := false
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `6
2 4 6 1 3 5
`
	expect := true
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `3
2 1 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `4
2 4 1 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `6
4 2 6 5 1 3
`
	expect := true
	runSample(t, s, expect)
}
