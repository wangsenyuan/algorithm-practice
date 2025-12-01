package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}

	return solve(a, edges)
}

func solve(a []int, edges [][]int) []int {
	n := len(a)
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	ans := make([]int, n)

	var dfs1 func(p int, u int)

	dfs1 = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				ans[v] = gcd(ans[u], a[v])
				dfs1(u, v)
			}
		}
	}

	dfs1(-1, 0)

	var fs []int
	for i := 1; i*i <= a[0]; i++ {
		if a[0]%i == 0 {
			fs = append(fs, i)
			if i*i != a[0] {
				fs = append(fs, a[0]/i)
			}
		}
	}
	slices.Sort(fs)

	m := len(fs)
	cnt := make([]int, m)

	var dfs2 func(p int, u int, dist int)

	dfs2 = func(p int, u int, dist int) {
		for i := range m {
			if a[u]%fs[i] == 0 {
				cnt[i]++
			}
			if cnt[i] >= dist {
				ans[u] = max(ans[u], fs[i])
			}
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs2(u, v, dist+1)
			}
		}

		for i := range m {
			if a[u]%fs[i] == 0 {
				cnt[i]--
			}
		}
	}

	dfs2(-1, 0, 0)

	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
