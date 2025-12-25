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
	_, _, best, res := drive(reader)
	fmt.Fprintln(writer, best)
	for _, e := range res {
		fmt.Fprintln(writer, e[0], e[1])
	}
}

func drive(reader *bufio.Reader) (n int32, edges [][]int32, best int32, res [][]int32) {
	var m int32
	fmt.Fscan(reader, &n, &m)
	edges = make([][]int32, m)
	for i := range m {
		edges[i] = make([]int32, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	best, res = solve(n, edges)
	return
}

func solve(n int32, edges [][]int32) (best int32, res [][]int32) {
	adj := make([][]int, n)

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}

	dfn := make([]int32, n)
	low := make([]int32, n)
	var timer int32
	stack := make([]int32, n)
	var top int
	vis := make([]bool, n)

	m := len(edges)

	comp := make([]int32, n)
	var cnt []int32

	var dfs func(p int32, u int32)
	dfs = func(p int32, u int32) {
		timer++
		dfn[u] = timer
		low[u] = timer
		vis[u] = true
		stack[top] = u
		top++

		for _, eid := range adj[u] {
			x, y := edges[eid][0]-1, edges[eid][1]-1
			v := x ^ y ^ u
			if p == v {
				continue
			}
			if dfn[v] == 0 {
				dfs(u, v)
				low[u] = min(low[u], low[v])
			} else if vis[v] {
				low[u] = min(low[u], dfn[v])
			}
		}

		if low[u] == dfn[u] {
			var w int32
			for top > 0 {
				v := stack[top-1]
				top--
				vis[v] = false
				comp[v] = int32(len(cnt))
				w++
				if u == v {
					break
				}
			}
			cnt = append(cnt, w)
		}
	}

	// 它们是连通
	dfs(-1, 0)

	var xid int
	for i := range cnt {
		if cnt[i] > cnt[xid] {
			xid = i
		}
	}

	best = cnt[xid]

	res = make([][]int32, m)

	var dfs2 func(u int32)

	res = make([][]int32, m)

	dfs2 = func(u int32) {
		vis[u] = true
		for _, eid := range adj[u] {
			x, y := edges[eid][0]-1, edges[eid][1]-1
			v := x ^ y ^ u
			if comp[v] == int32(xid) {
				if !vis[v] {
					res[eid] = []int32{u + 1, v + 1}
					dfs2(v)
				} else {
					// 这个是个回边
					res[eid] = []int32{v + 1, u + 1}
				}
			}
		}
	}

	for u := range n {
		if comp[u] == int32(xid) {
			dfs2(u)
			break
		}
	}

	que := make([]int32, n)
	dist := make([]int32, n)
	var head, tail int
	for i := range n {
		dist[i] = -1
		if comp[i] == int32(xid) {
			que[head] = i
			head++
			dist[i] = 0
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		for _, eid := range adj[u] {
			x, y := edges[eid][0]-1, edges[eid][1]-1
			v := x ^ y ^ u
			if comp[v] != int32(xid) && dist[v] < 0 {
				dist[v] = dist[u] + 1
				que[head] = v
				head++
			}
		}
	}

	for i, e := range edges {
		if res[i] == nil {
			u, v := e[0]-1, e[1]-1
			if dist[u] < dist[v] {
				u, v = v, u
			}
			res[i] = []int32{u + 1, v + 1}
		}
	}

	return
}
