package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	relationships := make([][]int, m)
	for i := 0; i < m; i++ {
		relationships[i] = make([]int, 3)
		fmt.Fscan(reader, &relationships[i][0], &relationships[i][1], &relationships[i][2])
	}
	return solve(n, relationships)
}

const mod = 1e9 + 7

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(n int, relationships [][]int) int {
	set := NewDSU(n)

	for _, rel := range relationships {
		a, b, c := rel[0], rel[1], rel[2]
		a--
		b--
		if c == 1 {
			set.Union(a, b)
		}
	}

	g := NewGraph(n, 2*len(relationships))

	for _, rel := range relationships {
		a, b, c := rel[0], rel[1], rel[2]
		a--
		b--
		if c == 0 {
			a = set.Find(a)
			b = set.Find(b)
			if a == b {
				return 0
			}
			g.AddEdge(a, b)
			g.AddEdge(b, a)
		}
	}

	vis := make([]int, n)
	var dfs func(p int, u int, c int) bool
	dfs = func(p int, u int, c int) bool {
		if vis[u] != 0 {
			return vis[u] == c
		}
		vis[u] = c
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && !dfs(u, v, -c) {
				return false
			}
		}
		return true
	}

	var comp int

	for i := range n {
		if set.Find(i) == i && vis[i] == 0 {
			if !dfs(-1, i, 1) {
				return 0
			}
			comp++
		}
	}

	return pow(2, comp-1)
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
