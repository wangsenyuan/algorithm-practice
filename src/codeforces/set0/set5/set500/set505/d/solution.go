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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	routes := make([][]int, m)
	for i := range m {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		routes[i] = []int{u, v}
	}
	return solve(n, routes)
}

func solve(n int, routes [][]int) int {
	m := len(routes)
	g := NewGraph(n, m)
	deg := make([]int, n)

	set := NewDSU(n)

	for i, route := range routes {
		u, v := route[0]-1, route[1]-1
		g.AddEdge(u, v, i)
		deg[v]++
		set.Union(u, v)
	}

	que := make([]int, n)
	var head, tail int

	for i := range n {
		if deg[i] == 0 {
			que[head] = i
			head++
		}
	}
	res := n

	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			deg[v]--
			if deg[v] == 0 {
				que[head] = v
				head++
			}
		}
	}

	cycles := make([]bool, n)

	for u := range n {
		if deg[u] > 0 {
			cycles[set.Find(u)] = true
		}
	}

	var roots []int

	for u := range n {
		u = set.Find(u)
		if !cycles[u] {
			roots = append(roots, u)
		}
	}

	slices.Sort(roots)
	roots = slices.Compact(roots)

	return res - len(roots)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.val = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

type DSU struct {
	arr  []int
	cnt  []int
	size int
}

func NewDSU(n int) *DSU {
	set := new(DSU)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = n
	return set
}

func (set *DSU) Find(u int) int {
	if set.arr[u] != u {
		set.arr[u] = set.Find(set.arr[u])
	}
	return set.arr[u]
}

func (set *DSU) Union(a, b int) bool {
	a = set.Find(a)
	b = set.Find(b)
	if a == b {
		return false
	}
	if set.cnt[a] < set.cnt[b] {
		a, b = b, a
	}
	set.cnt[a] += set.cnt[b]
	set.arr[b] = a
	set.size--
	return true
}
