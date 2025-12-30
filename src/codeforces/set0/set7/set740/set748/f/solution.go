package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, d, res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, len(d))
	for _, v := range d {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
	for _, v := range res {
		fmt.Fprintln(writer, v[0], v[1], v[2])
	}
}

func drive(reader *bufio.Reader) (n int, edges [][]int, c []int, d []int, res [][]int) {
	var k int
	fmt.Fscan(reader, &n, &k)
	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	c = make([]int, 2*k)
	for i := range 2 * k {
		fmt.Fscan(reader, &c[i])
	}
	d, res = solve(n, edges, c)
	return
}

func solve(n int, edges [][]int, c []int) (d []int, res [][]int) {
	adj := make([][]int, n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	sz := make([]int, n)
	flag := make([]int, n)
	for _, i := range c {
		sz[i-1] = 1
		flag[i-1] = 1
	}

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0)

	k := len(c) / 2

	var dfs2 func(p int, u int) int

	dfs2 = func(p int, u int) int {
		for _, v := range adj[u] {
			if p != v && sz[v] >= k {
				return dfs2(u, v)
			}
		}
		return u
	}

	center := dfs2(-1, 0)

	var arr []int

	var dfs3 func(p int, u int)
	dfs3 = func(p int, u int) {
		if flag[u] == 1 {
			arr = append(arr, u+1)
		}
		for _, v := range adj[u] {
			if p != v {
				dfs3(u, v)
			}
		}
	}

	for _, v := range adj[center] {
		dfs3(center, v)
	}

	if flag[center] == 1 {
		arr = append(arr, center+1)
	}

	for i := 0; i < k; i++ {
		res = append(res, []int{arr[i], arr[i+k], center + 1})
	}

	d = append(d, center+1)

	return
}
