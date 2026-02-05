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
		res := drive(reader)
		for _, x := range res {
			fmt.Fprintln(writer, x)
		}
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	col := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &col[i])
	}
	adj := make([][]pair, n+1)
	for range n - 1 {
		var u, v, c int
		fmt.Fscan(reader, &u, &v, &c)
		adj[u] = append(adj[u], pair{v, c})
		adj[v] = append(adj[v], pair{u, c})
	}
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(n, col, adj, queries)
}

type pair struct {
	to, cost int
}

func solve(n int, col []int, adj [][]pair, queries [][]int) []int {
	fa := make([]int, n+1)
	zn := make([]int, n+1)
	cnt := make([]map[int]int, n+1)
	for i := range n + 1 {
		cnt[i] = make(map[int]int)
	}

	var costt int

	var dfs func(p int, v int)
	dfs = func(p int, v int) {
		fa[v] = p
		for _, e := range adj[v] {
			u, c := e.to, (e.cost)
			if u == p {
				continue
			}
			if col[v] != col[u] {
				costt += c
			}
			dfs(v, u)
			zn[u] = c
			cnt[v][col[u]] += zn[u]
		}
	}

	dfs(0, 1)

	update := func(v, x int) {
		if v != 1 {
			cnt[fa[v]][col[v]] -= zn[v]
			if col[fa[v]] == col[v] {
				costt += zn[v]
			}
			cnt[fa[v]][x] += zn[v]
			if col[fa[v]] == x {
				costt -= zn[v]
			}
		}
		costt += cnt[v][col[v]]
		costt -= cnt[v][x]
		col[v] = x
	}

	res := make([]int, len(queries))
	for i, qr := range queries {
		pos, x := qr[0], qr[1]
		update(pos, x)
		res[i] = costt
	}

	return res
}
