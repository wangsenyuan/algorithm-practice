package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, lambs := process(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	n := len(lambs)
	g := NewGraph(n, n)
	res[0]--
	res[1]--
	var root int
	for u, lamb := range lambs {
		if lamb[0] == 0 {
			root = u
		} else if u != res[0] && u != res[1] {
			p := lamb[0] - 1
			g.AddEdge(p, u)
		}
	}

	var dfs func(u int) int

	dfs = func(u int) int {
		res := lambs[u][1]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			res += dfs(g.to[i])
		}
		return res
	}

	x := dfs(root)
	y := dfs(res[0])
	z := dfs(res[1])

	if x != y || y != z {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6
2 4
0 5
4 2
2 1
1 1
4 2`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
2 4
0 6
4 2
2 1
1 1
4 2
`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `49
2 1
43 1
31 1
0 34
14 1
29 1
40 1
40 1
39 1
1 1
28 1
43 1
44 1
44 1
43 1
12 5
12 5
28 1
7 1
15 1
43 1
4 49
15 1
8 1
17 9
5 1
43 1
43 1
40 1
2 1
7 2
24 1
12 1
27 1
2 1
43 1
28 1
5 1
27 1
28 1
37 1
27 1
4 1
28 1
31 1
40 1
21 1
38 1
44 1
`, false)
}

func TestSample4(t *testing.T) {
	runSample(t, `4
0 1
1 -1
2 1
3 -1
`, false)
}
