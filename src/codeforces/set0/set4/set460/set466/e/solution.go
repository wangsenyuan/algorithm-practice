package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for _, v := range res {
		if v {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func drive(reader *bufio.Reader) []bool {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	events := make([][]int, m)
	for i := range m {
		var t int
		fmt.Fscan(reader, &t)
		switch t {
		case 1:
			var x, y int
			fmt.Fscan(reader, &x, &y)
			events[i] = []int{1, x, y}
		case 2:
			var x int
			fmt.Fscan(reader, &x)
			events[i] = []int{2, x}
		case 3:
			var y, j int
			fmt.Fscan(reader, &y, &j)
			events[i] = []int{3, y, j}
		}
	}
	return solve(n, events)
}

func solve(n int, events [][]int) []bool {
	m := len(events)
	g := NewGraph(n, m)

	send := make([][]int, n)
	var docs [][]int

	event_id := make([]int, n)

	for i := range n {
		event_id[i] = -1
	}

	deg := make([]int, n)
	for i, cur := range events {
		switch cur[0] {
		case 1:
			// y becomes boss of x
			x, y := cur[1]-1, cur[2]-1
			g.AddEdge(y, x)
			deg[x]++
			event_id[x] = i
		case 2:
			x := cur[1] - 1
			send[x] = append(send[x], len(docs))
			docs = append(docs, []int{i, x})
		}
	}
	dep := make([]int, n)
	sz := make([]int, n)
	pos := make([]int, n)
	var id int
	var dfs func(u int)
	dfs = func(u int) {
		sz[u] = 1
		pos[u] = id
		id++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dep[v] = dep[u] + 1
			dfs(v)
			sz[u] += sz[v]
		}
	}

	for u := range n {
		if deg[u] == 0 {
			dfs(u)
		}
	}

	isAnc := func(u int, v int) bool {
		return pos[u] <= pos[v] && pos[v] < pos[u]+sz[u]
	}

	// 这个是关于高度的树
	tr := NewTree(n)

	reach_at := make([]int, len(docs))

	var dfs2 func(u int)
	dfs2 = func(u int) {
		if event_id[u] >= 0 {
			// 这个不是root
			tr.Update(dep[u], event_id[u])
		}

		for _, i := range send[u] {
			// i 是文件的序号, docs[i]是i的操作序列好
			if tr[0] < docs[i][0] {
				// 在发送文件前，整条路径已经完整了
				reach_at[i] = 0
			} else {
				reach_at[i] = tr.GetRightMostPosGe(docs[i][0])
			}
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs2(v)
		}

		if event_id[u] >= 0 {
			tr.Update(dep[u], -1)
		}
	}

	for u := range n {
		if deg[u] == 0 {
			dfs2(u)
		}
	}

	var ans []bool

	for _, cur := range events {
		if cur[0] == 3 {
			y, i := cur[1]-1, cur[2]-1
			x := docs[i][1]
			if !isAnc(y, x) {
				ans = append(ans, false)
				continue
			}
			ans = append(ans, reach_at[i] <= dep[y])
		}
	}

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

type Tree []int

func NewTree(n int) Tree {
	tr := make(Tree, 4*n)
	for i := range 4 * n {
		tr[i] = -1
	}
	return tr
}

func (tr Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr[i] = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr[i] = max(tr[i*2+1], tr[i*2+2])
	}
	n := len(tr) / 4
	f(0, 0, n-1)
}

func (tr Tree) GetRightMostPosGe(v int) int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		mid := (l + r) >> 1
		if tr[2*i+2] >= v {
			return f(i*2+2, mid+1, r)
		}
		return f(i*2+1, l, mid)
	}
	n := len(tr) / 4
	return f(0, 0, n-1)
}
