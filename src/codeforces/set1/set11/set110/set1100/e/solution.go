package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_, _, best, res := drive(reader)
	fmt.Fprintln(writer, best, len(res))
	for _, v := range res {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) (n int, edges [][]int, best int, res []int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	edges = make([][]int, m)
	for i := 0; i < m; i++ {
		var u, v, c int
		fmt.Fscan(reader, &u, &v, &c)
		edges[i] = []int{u, v, c}
	}
	best, res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) (best int, res []int) {
	m := len(edges)
	g := NewGraph(n, m)

	marked := make([]int, n)

	var dfs func(u int) bool
	dfs = func(u int) bool {
		marked[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if marked[v] == 1 || marked[v] == 0 && !dfs(v) {
				return false
			}
		}
		marked[u]++
		return true
	}

	check := func(w int) bool {
		g.Reset()

		for i, e := range edges {
			u, v, c := e[0]-1, e[1]-1, e[2]
			if c > w {
				g.AddEdge(u, v, i)
			}
		}

		clear(marked)

		for i := range n {
			if marked[i] == 0 {
				if !dfs(i) {
					return false
				}
			}
		}

		return true
	}

	var arr []int
	for _, e := range edges {
		arr = append(arr, e[2])
	}
	arr = append(arr, 0)
	slices.Sort(arr)
	arr = slices.Compact(arr)
	arr = append(arr, 1<<30)
	at := sort.Search(len(arr), func(i int) bool {
		return check(arr[i])
	})

	best = arr[at]
	g.Reset()

	for i, e := range edges {
		u, v, c := e[0]-1, e[1]-1, e[2]
		if c > best {
			g.AddEdge(u, v, i)
		}
	}

	clear(marked)
	var ord []int

	pos := make([]int, n)

	var dfs1 func(u int)
	dfs1 = func(u int) {
		marked[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if marked[v] == 0 {
				dfs1(v)
			}
		}
		pos[u] = len(ord)
		ord = append(ord, u)
	}

	for i := range n {
		if marked[i] == 0 {
			dfs1(i)
		}
	}

	// 现在ord倒序是topo排序（因为在原图中，没有环）

	for i, e := range edges {
		u, v, c := e[0]-1, e[1]-1, e[2]
		if c <= best && pos[u] < pos[v] {
			res = append(res, i+1)
		}
	}

	return
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

func (g *Graph) Reset() {
	for i := range g.nodes {
		g.nodes[i] = 0
	}
	g.cur = 0
}
