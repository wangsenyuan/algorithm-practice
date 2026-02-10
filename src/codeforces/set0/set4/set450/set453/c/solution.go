package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, ok, res := drive(reader)
	if !ok {
		fmt.Println("-1")
		return
	}
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (n int, edges [][]int, a []int, ok bool, res []int) {
	var m int

	fmt.Fscan(reader, &n, &m)

	edges = make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	ok, res = solve(n, edges, a)
	return
}

func solve(n int, edges [][]int, a []int) (bool, []int) {
	if slices.Max(a) == 0 {
		return true, nil
	}

	// m := len(edges)

	adj := make([][]int, n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	occ := make([]int, n)
	vis := make([]bool, n)
	var path []int
	add := func(u int) {
		occ[u] ^= 1
		path = append(path, u+1)
	}

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		vis[u] = true
		add(u)
		for _, v := range adj[u] {
			if !vis[v] {
				dfs(u, v)
				add(u)
			}
		}
		if a[u] != occ[u] && p >= 0 {
			add(p)
			add(u)
		}
	}

	for i := range n {
		if a[i] == 1 {
			dfs(-1, i)
			if a[i] != occ[i] {
				occ[i] ^= 1
				path = path[:len(path)-1]
			}
			if !slices.Equal(occ, a) {
				return false, nil
			}
			return true, path
		}
	}
	return false, nil
}
