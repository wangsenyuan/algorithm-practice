package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	mainCourses, deps, res := drive(reader)
	if len(res) != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if expect == 0 {
		return
	}

	n := len(deps)
	adj := make([][]int, n)

	deg := make([]int, n)
	for i, cur := range deps {
		for _, v := range cur {
			adj[v-1] = append(adj[v-1], i)
		}
		deg[i] = len(cur)
	}

	marked := make([]bool, n)

	for _, v := range res {
		v--
		if deg[v] > 0 {
			t.Fatalf("Sample result %v, not valid", res)
		}
		marked[v] = true
		for _, u := range adj[v] {
			deg[u]--
		}
	}

	for _, v := range mainCourses {
		if !marked[v-1] {
			t.Fatalf("Sample result %v, not valid", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6 2
5 3
0
0
0
2 2 1
1 4
1 5
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9 3
3 9 5
0
0
3 9 4 5
0
0
1 8
1 6
1 2
2 1 2
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3
1 2 3
1 2
1 3
1 1
`
	expect := 0
	runSample(t, s, expect)
}