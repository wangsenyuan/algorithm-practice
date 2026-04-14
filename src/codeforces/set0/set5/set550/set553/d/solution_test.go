package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, roads, F, res := drive(reader)

	if len(res) == 0 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	adj := make([][]int, n)
	for _, e := range roads {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	marked := make([]bool, n)
	for _, v := range F {
		marked[v-1] = true
	}

	S := make([]bool, n)

	get := func(res []int) (d1 int, d2 int) {
		d1 = 1
		d2 = 1
		clear(S)
		for _, u := range res {
			u--
			if marked[u] {
				t.Fatalf("Cant select %d, it is a fort", u+1)
			}
			S[u] = true
		}
		for _, u := range res {
			u--
			var c1 int
			c2 := len(adj[u])
			for _, v := range adj[u] {
				if S[v] {
					c1++
				}
			}
			if d1*c2 > c1*d2 {
				d1 = c1
				d2 = c2
			}
		}
		return
	}

	d1, d2 := get(expect)
	c1, c2 := get(res)

	if d1*c2 != c1*d2 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `9 8 4
3 9 6 8
1 2
1 3
1 4
1 5
2 6
2 7
2 8
2 9
`
	expect := []int{1, 4, 5}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 8 2
2 9
1 3
2 9
4 5
5 6
6 7
7 8
8 10
10 4
`
	expect := []int{1, 5, 4, 8, 10, 6, 3, 7}
	runSample(t, s, expect)
}
