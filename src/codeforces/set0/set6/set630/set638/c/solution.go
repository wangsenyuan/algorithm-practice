package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, len(res))
	for _, r := range res {
		fmt.Fprint(writer, len(r))
		for _, i := range r {
			fmt.Fprint(writer, " ", i)
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) (n int, edges [][]int, res [][]int) {
	fmt.Fscan(reader, &n)
	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) [][]int {
	adj := make([][]int, n)

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}

	var k int
	var root int
	for u := range n {
		if len(adj[u]) > k {
			k = len(adj[u])
			root = u
		}
	}

	ans := make([][]int, k)

	var dfs func(p int, u int, x int)

	dfs = func(p int, u int, x int) {
		var mex int
		for _, i := range adj[u] {
			v := (edges[i][0] - 1) ^ (edges[i][1] - 1) ^ u
			if p == v {
				continue
			}
			if mex == x {
				mex++
			}
			ans[mex] = append(ans[mex], i+1)
			dfs(u, v, mex)
			mex++
		}
	}

	dfs(-1, root, -1)

	return ans
}
