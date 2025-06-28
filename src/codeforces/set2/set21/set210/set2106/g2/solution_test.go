package main

import (
	"math/rand/v2"
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

	var cnt int

	ask := func(arr []int) int {
		cnt++
		var res int
		for _, u := range arr {
			res += get(u)
		}
		return res
	}

	toggle := func(u int) {
		cnt++
		u--
		value[u] *= -1
	}

	ans := solve(len(value), edges, ask, toggle)

	if cnt > n+200 {
		t.Fatalf("Sample asked too much times %d", cnt)
	}

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

func TestSample4(t *testing.T) {
	edges := [][]int{
		{1, 4},
		{4, 2},
		{2, 3},
	}
	values := []int{-1, 1, -1, 1}
	root := 3
	runSample(t, edges, values, root)
}

func TestSample5(t *testing.T) {
	edges := [][]int{
		{1, 4},
		{4, 2},
		{2, 3},
	}
	values := []int{-1, 1, -1, 1}
	root := 1
	runSample(t, edges, values, root)
}

func TestSample6(t *testing.T) {
	edges := [][]int{
		{1, 4},
		{4, 2},
		{2, 3},
	}
	values := []int{-1, 1, -1, 1}
	root := 2
	runSample(t, edges, values, root)
}

func TestSample7(t *testing.T) {
	edges := [][]int{
		{1, 2},
		{2, 7},
		{7, 3},
		{7, 4},
		{7, 5},
		{7, 6},
	}
	values := []int{-1, 1, -1, 1, -1, 1, -1}
	root := 2
	runSample(t, edges, values, root)
}

func TestSample8(t *testing.T) {
	edges := [][]int{
		{1, 2},
		{2, 7},
		{7, 3},
		{7, 4},
		{7, 5},
		{7, 6},
	}
	values := []int{-1, 1, -1, 1, -1, 1, -1}
	root := 7
	runSample(t, edges, values, root)
}

func TestSample9(t *testing.T) {
	edges := [][]int{
		{26, 1},
		{26, 2},
		{26, 3},
		{26, 4},
		{26, 5},
		{26, 6},
		{26, 7},
		{26, 8},
		{26, 9},
		{26, 10},
		{26, 11},
		{26, 12},
		{26, 13},
		{26, 14},
		{26, 15},
		{26, 16},
		{26, 17},
		{26, 18},
		{26, 19},
		{26, 20},
		{26, 21},
		{26, 22},
		{26, 23},
		{26, 24},
		{26, 25},
		{26, 27},
		{26, 28},
		{26, 29},
		{26, 30},
		{26, 31},
	}
	values := []int{-1, 1, -1, 1, 1, -1, 1, 1, -1, 1, -1, 1, -1, 1, 1, -1, -1, -1, 1, 1, 1, -1, 1, -1, 1, 1, 1, 1, -1, 1, -1}
	root := 26
	runSample(t, edges, values, root)
}

func TestSample10(t *testing.T) {
	n := 1000
	values := make([]int, n)
	for i := range n {
		values[i] = rand.IntN(2)
		if values[i] == 0 {
			values[i] = -1
		}
	}
	edges := make([][]int, n-1)
	for i := 1; i < n; i++ {
		edges[i-1] = []int{i + 1, 1}
	}
	root := rand.IntN(n) + 1

	runSample(t, edges, values, root)
}
