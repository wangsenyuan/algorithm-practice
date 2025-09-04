package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

const inf = 1 << 60

type edge struct {
	u int
	v int
	w int
}

func solve(a [][]int) bool {
	n := len(a)

	var edges []edge
	for i := range n {
		if a[i][i] != 0 {
			return false
		}
		for j := range i {
			if a[j][i] <= 0 || a[i][j] != a[j][i] {
				return false
			}
			edges = append(edges, edge{j, i, a[j][i]})
		}
	}

	slices.SortFunc(edges, func(a, b edge) int {
		return a.w - b.w
	})

	set := NewDSU(n)
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		if set.Union(e.u, e.v) {
			g.AddEdge(e.u, e.v, e.w)
			g.AddEdge(e.v, e.u, e.w)
		}
	}

	for i := range n {
		if set.Find(i) != set.Find(0) {
			// not a tree
			return false
		}
	}
	dist := make([]int, n)

	que := make([]int, n)

	bfs := func(u int) {
		for i := range n {
			dist[i] = -1
		}
		dist[u] = 0
		var head, tail int
		que[head] = u
		head++
		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + g.val[i]
					que[head] = v
					head++
				}
			}
		}
	}

	for u := range n {
		bfs(u)
		if !reflect.DeepEqual(dist, a[u]) {
			return false
		}
	}

	return true
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

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
