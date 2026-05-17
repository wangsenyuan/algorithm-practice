package main

import "testing"

func runSample(t *testing.T, x int, y int, expect bool) {
	ok, res := solve(x, y)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !ok {
		return
	}
	n := x + y
	if len(res) != n-1 {
		t.Fatalf("Sample result %v, not correct", res)
	}
	// 必须检查是否一棵树
	adj := make([][]int, n)
	for _, e := range res {
		u, v := e[0], e[1]
		u--
		v--
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	marked := make([]bool, n)
	var dfs func(p int, u int) (int, bool)
	cnt := make([]int, 2)
	dfs = func(p int, u int) (int, bool) {
		if marked[u] {
			// a loop back
			return 0, false
		}
		marked[u] = true
		sz := 1
		for _, v := range adj[u] {
			if p != v {
				tmp, ok := dfs(u, v)
				if !ok {
					return 0, false
				}
				sz += tmp
			}
		}
		cnt[sz%2]++
		return sz, true
	}
	if _, ok := dfs(-1, 0); !ok {
		t.Fatalf("Sample result %v, not a tree", res)
	}
	for i := range n {
		if !marked[i] {
			t.Fatalf("Sample result %v, not a tree", res)
		}
	}
	if cnt[0] != x || cnt[1] != y {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	x := 1
	y := 1
	expect := true
	runSample(t, x, y, expect)
}

func TestSample2(t *testing.T) {
	x := 2
	y := 1
	expect := false
	runSample(t, x, y, expect)
}

func TestSample3(t *testing.T) {
	x := 0
	y := 3
	expect := true
	runSample(t, x, y, expect)
}

func TestSample4(t *testing.T) {
	x := 3
	y := 4
	expect := true
	runSample(t, x, y, expect)
}

func TestSample5(t *testing.T) {
	x := 0
	y := 2
	expect := false
	runSample(t, x, y, expect)
}

func TestSample6(t *testing.T) {
	x := 1
	y := 0
	expect := false
	runSample(t, x, y, expect)
}

func TestSample7(t *testing.T) {
	x := 4
	y := 7
	expect := true
	runSample(t, x, y, expect)
}
