package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(a, edges)
}

type node struct {
	id int
	x  int
}

func solve(a []int, edges [][]int) float64 {
	n := len(a)
	g := NewGraph(n, 2*len(edges))
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	arr := make([]node, n)
	for i, v := range a {
		arr[i] = node{i, v}
	}

	slices.SortFunc(arr, func(a, b node) int {
		return cmp.Or(b.x-a.x, a.id-b.id)
	})

	pos := make([]int, n)
	for i, cur := range arr {
		pos[cur.id] = i
	}

	set := NewDSU(n)

	var sum int

	for _, cur := range arr {
		u := cur.id
		// 1肯定是对的，因为它不会被它前面的union
		sz := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			// 这个地方不对
			if pos[v] < pos[u] && set.Find(v) != set.Find(u) {
				tmp := set.cnt[set.Find(v)]
				sum += cur.x * sz * tmp
				sz += tmp
				set.Union(u, v)
			}
		}
	}

	return float64(sum) / float64(n*(n-1)/2)
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

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}
