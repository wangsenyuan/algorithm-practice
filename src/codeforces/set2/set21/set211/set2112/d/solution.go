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
		_, _, res := drive(reader)
		if len(res) == 0 {
			fmt.Fprintln(writer, "NO")
			continue
		}
		fmt.Fprintln(writer, "YES")
		for _, e := range res {
			fmt.Fprintln(writer, e[0], e[1])
		}
	}
}

func drive(reader *bufio.Reader) (n int, edges [][]int, res [][]int) {
	fmt.Fscan(reader, &n)
	edges = make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) [][]int {
	if n <= 2 {
		// no way
		return nil
	}
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var res [][]int

	var dfs func(p int, u int, d int)
	dfs = func(p int, u int, d int) {
		if d == 0 {
			res = append(res, []int{p + 1, u + 1})
		} else {
			res = append(res, []int{u + 1, p + 1})
		}
		for _, v := range adj[u] {
			if v != p {
				dfs(u, v, d^1)
			}
		}
	}

	for u := range n {
		if len(adj[u]) == 2 {
			// u is the root
			for i, v := range adj[u] {
				dfs(u, v, i)
			}
			return res
		}
	}

	return nil
}
