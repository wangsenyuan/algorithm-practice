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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		best, res := drive(reader)

		fmt.Fprintln(writer, best)
		for _, e := range res {
			fmt.Fprintln(writer, e[0], e[1])
		}
	}
}

func drive(reader *bufio.Reader) (best int, res [][]int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	best, res = solve(n, edges)
	return
}

type edge struct {
	to int
	id int
}

func solve(n int, edges [][]int) (best int, res [][]int) {
	m := len(edges)

	adj := make([][]edge, n+1)

	g := make([]map[int]bool, n+1)
	for i := range n + 1 {
		g[i] = make(map[int]bool)
	}

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1

		adj[u] = append(adj[u], edge{v, i})
		adj[v] = append(adj[v], edge{u, i})
		g[u][v] = true
		g[v][u] = true
	}

	var path []int

	marked := make([]bool, m+n+1)

	eid := m

	for u := range n {
		if len(adj[u])%2 == 0 {
			best++
		} else {
			adj[n] = append(adj[n], edge{u, eid})
			adj[u] = append(adj[u], edge{n, eid})
			eid++
		}
	}

	vis := make([]bool, n+1)

	var dfs func(u int)
	dfs = func(u int) {
		vis[u] = true
		for len(adj[u]) > 0 {
			cur := adj[u][0]
			adj[u] = adj[u][1:]
			v := cur.to
			if !marked[cur.id] {
				marked[cur.id] = true
				dfs(v)
			}
		}
		path = append(path, u)
	}

	for u := range n {
		if !vis[u] {
			dfs(u)
			path = append(path, u)
		}
	}

	for i := 0; i+1 < len(path); i++ {
		u, v := path[i], path[i+1]

		if g[u][v] {
			res = append(res, []int{u + 1, v + 1})
		}
	}

	return best, res
}
