package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(c, edges, queries)
}

func solve(c []int, edges [][]int, queries [][]int) []int {
	n := len(c)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	qry := make([][]int, n)
	for i, cur := range queries {
		v := cur[0] - 1
		qry[v] = append(qry[v], i)
	}

	sz := make([]int, n)
	big := make([]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		sz[u] = 1
		big[u] = -1
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
				if big[u] < 0 || sz[v] > sz[big[u]] {
					big[u] = v
				}
			}
		}
	}
	dfs(-1, 0)

	mc := slices.Max(c)

	freqOfColor := make([]int, mc+1)
	cntOfFreq := make([]int, n+1)

	begin := make([]int, n)
	end := make([]int, n)
	var timer int
	at := make([]int, n)

	rem := func(u int) {
		for i := begin[u]; i < end[u]; i++ {
			cntOfFreq[freqOfColor[at[i]]]--
			freqOfColor[at[i]]--
		}
	}

	ins := func(u int) {
		for i := begin[u]; i < end[u]; i++ {
			freqOfColor[at[i]]++
			cntOfFreq[freqOfColor[at[i]]]++
		}
	}

	ans := make([]int, len(queries))

	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		begin[u] = timer
		at[timer] = c[u]
		timer++

		for _, v := range adj[u] {
			if v != p && v != big[u] {
				dfs2(u, v)
				rem(v)
			}
		}
		if big[u] >= 0 {
			dfs2(u, big[u])
		}
		for _, v := range adj[u] {
			if v != p && v != big[u] {
				ins(v)
			}
		}
		freqOfColor[c[u]]++
		cntOfFreq[freqOfColor[c[u]]]++
		end[u] = timer

		for _, i := range qry[u] {
			k := queries[i][1]
			if k <= n {
				ans[i] = cntOfFreq[queries[i][1]]
			}
		}
	}

	dfs2(-1, 0)

	return ans
}
