package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, change, res := drive(reader)
	if expect != change {
		t.Fatalf("Sample expect %d, but got %d", expect, change)
	}
	var cnt int
	for i := range a {
		if a[i] != res[i] {
			cnt++
		}
	}
	if cnt != expect {
		t.Fatalf("Sample result %v, not correct", res)
	}
	n := len(a)
	root := -1
	adj := make([][]int, n)
	for i := range n {
		res[i]--
		if res[i] != i {
			adj[res[i]] = append(adj[res[i]], i)
		} else {
			if root < 0 {
				root = i
			} else {
				t.Fatalf("Sample result %v, have multiple roots", res)
			}
		}
	}
	if root < 0 {
		t.Fatalf("Sample result %v, have no root", res)
	}
	marked := make([]bool, n)
	var dfs func(u int) int
	dfs = func(u int) int {
		if marked[u] {
			t.Fatalf("Sample result %v, is not a valid tree", res)
		}
		marked[u] = true
		sz := 1
		for _, v := range adj[u] {
			sz += dfs(v)
		}
		return sz
	}
	sz := dfs(root)
	if sz != n {
		t.Fatalf("Sample result %v, is not a valid tree", res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
2 3 3 4
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
3 2 2 5 3
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8
2 3 5 4 1 6 6 7
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7
4 3 2 6 3 5 2
`
	expect := 1
	runSample(t, s, expect)
}
