package main

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int) {

	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
	}

	for u := 1; u <= n; u++ {
		slices.Sort(adj[u])
	}

	// 先需要生成所有的path
	var paths [][]int

	var dfs func(u int, cur []int)
	dfs = func(u int, cur []int) {
		cur = append(cur, u)
		paths = append(paths, slices.Clone(cur))
		for _, v := range adj[u] {
			dfs(v, cur)
		}
		cur = cur[:len(cur)-1]
	}

	for u := 1; u <= n; u++ {
		dfs(u, nil)
	}

	var cnt int
	ask := func(k int) []int {
		if k == 0 {
			t.Fatal("k cannot be 0")
		}
		cnt++
		// 没有32的倍数
		if cnt > (n + len(edges)) {
			t.Fatalf("Sample asked too much times %d", cnt)
		}
		if k > len(paths) {
			return nil
		}
		return slices.Clone(paths[k-1])
	}

	res := solve(n, ask)

	slices.SortFunc(edges, func(a []int, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	slices.SortFunc(res, func(a []int, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	if !reflect.DeepEqual(res, edges) {
		t.Fatalf("Sample expect %v, but got %v", edges, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 3},
		{1, 2},
		{2, 4},
		{3, 4},
		{2, 5},
		{3, 5},
	}
	runSample(t, n, edges)
}

func TestSample2(t *testing.T) {
	n := 2
	runSample(t, n, nil)
}

func TestSample3(t *testing.T) {
	n := 3
	edges := [][]int{
		{2, 1},
	}
	runSample(t, n, edges)
}
