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
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	var q int
	fmt.Fscan(reader, &q)

	queries := make([]int, q)
	for i := range q {
		fmt.Fscan(reader, &queries[i])
	}

	return solve(n, edges, queries)
}

func solve(n int, edges [][]int, queries []int) []int {
	k := len(queries)

	fa := make([]int, n)
	adj := make([]map[int]bool, n)

	for i := range n {
		fa[i] = i
		adj[i] = make(map[int]bool)
	}

	var find func(x int) int

	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u][v] = true
		adj[v][u] = true
	}

	tot := len(edges)

	play := func(id int) {
		u, v := edges[id][0]-1, edges[id][1]-1
		u, v = find(u), find(v)
		if u == v {
			return
		}

		if len(adj[u]) < len(adj[v]) {
			u, v = v, u
		}

		// 对外的全部删除掉
		tot -= len(adj[u])
		tot -= len(adj[v])
		// 中间那条边多删除了一次
		tot++
		delete(adj[u], v)
		delete(adj[v], u)
		for w := range adj[v] {
			if !adj[w][u] {
				adj[u][w] = true
				adj[w][u] = true
			}
			delete(adj[w], v)
		}

		fa[v] = u

		tot += len(adj[u])
	}

	ans := make([]int, k)

	for i, id := range queries {
		play(id - 1)
		ans[i] = tot
	}

	return ans
}
