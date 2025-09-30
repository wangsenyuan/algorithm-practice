package main

import (
	"bufio"
	"fmt"
	"os"
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
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		for j := range 2 {
			fmt.Fscan(reader, &edges[i][j])
		}
	}
	return solve(n, c, edges)
}

type pair struct {
	first  int
	second int
}

func solve(n int, c []int, edges [][]int) []int {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	ans := make([]int, n)
	cnt := make([]int, n)

	merge := func(u int, v int, a map[int]int, b map[int]int) map[int]int {
		if len(a) > len(b) {
			a, b = b, a
			// 这里不交换的，结果就不对了
			// 这里迭代的是b，所以a的记录是可以保留的
		} else {
			cnt[u] = cnt[v]
			ans[u] = ans[v]
		}

		for k, v := range a {
			b[k] += v
			if b[k] > cnt[u] {
				cnt[u] = b[k]
				ans[u] = k
			} else if b[k] == cnt[u] {
				ans[u] += k
			}
		}
		return b
	}

	var dfs func(p int, u int) map[int]int
	dfs = func(p int, u int) map[int]int {
		res := make(map[int]int)
		cnt[u] = 1
		ans[u] = c[u]
		res[c[u]] = 1

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp := dfs(u, v)
				res = merge(u, v, res, tmp)
			}
		}

		return res
	}

	dfs(-1, 0)

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
