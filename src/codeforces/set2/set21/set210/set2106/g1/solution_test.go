package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, edges [][]int, value []int, root int) {
	n := len(value)
	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	fa := make([]int, n+1)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = p
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}
	}

	dfs(0, root)

	var get func(u int) int

	get = func(u int) int {
		if u == 0 {
			return 0
		}
		return value[u-1] + get(fa[u])
	}

	ask := func(arr []int) int {
		var res int
		for _, u := range arr {
			res += get(u)
		}
		return res
	}

	toggle := func(u int) {
		u--
		value[u] *= -1
	}

	ans := solve(len(value), edges, ask, toggle)

	if !reflect.DeepEqual(ans, value) {
		t.Fatalf("Sample expect %v, but got %v", value, ans)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	values := []int{-1, 1, -1, 1}
	root := 2
	runSample(t, edges, values, root)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	values := []int{-1, 1, -1, 1}
	root := 1
	runSample(t, edges, values, root)
}

func TestSample3(t *testing.T) {
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	values := []int{-1, 1, -1, 1}
	root := 3
	runSample(t, edges, values, root)
}
