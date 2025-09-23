package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, v := range res {
		buf.WriteString(fmt.Sprintf("%d\n", v))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	relations := make([][]int, m)
	for i := range m {
		var f, p int
		fmt.Fscan(reader, &f, &p)
		relations[i] = []int{f, p}
	}
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, relations, a)
}

func solve(n int, relations [][]int, a []int) []int {
	// 有很多课树
	g := NewGraph(n, n)
	r := NewGraph(n, n)
	deg := make([]int, n)
	for _, cur := range relations {
		f, p := cur[0]-1, cur[1]-1
		g.AddEdge(f, p)
		deg[f]++
		r.AddEdge(p, f)
	}

	level := make([]int, n)
	que := make([]int, n)
	var head, tail int
	for u := range n {
		if deg[u] == 0 {
			que[head] = u
			head++
		}
	}

	var roots []int

	for tail < head {
		u := que[tail]
		tail++
		ok := true
		for i := r.nodes[u]; i > 0; i = r.next[i] {
			v := r.to[i]
			ok = false
			deg[v]--
			level[v] = max(level[v], level[u]+1)
			if deg[v] == 0 {
				que[head] = v
				head++
			}
		}
		if ok {
			roots = append(roots, u)
		}
	}

	set := slices.Clone(a)
	slices.Sort(set)
	set = slices.Compact(set)

	// 只要保证在一棵树中，叶子节点比祖先节点，更早出现就可以了
	slices.SortFunc(set, func(x int, y int) int {
		return level[x-1] - level[y-1]
	})

	marked := make([]bool, n)
	for _, v := range set {
		marked[v-1] = true
	}

	var dfs2 func(u int, w int) bool
	dfs2 = func(u int, w int) bool {
		if marked[u] {
			w = u
		}
		if a[u]-1 != w {
			return false
		}
		// a[u] = w
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs2(v, w) {
				return false
			}
		}
		return true
	}

	for _, u := range roots {
		if deg[u] == 0 && !dfs2(u, u) {
			return nil
		}
	}

	return set
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
