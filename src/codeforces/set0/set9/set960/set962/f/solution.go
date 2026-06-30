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
	fmt.Fprintln(writer, len(res))
	if len(res) > 0 {
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
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
	return solve(n, edges)
}

func solve(n int, edges [][]int) []int {
	adj := make([][]int, n)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}

	getEdgeTo := func(eid int, u int) int {
		return (edges[eid][0] - 1) ^ (edges[eid][1] - 1) ^ u
	}

	low := make([]int, n)
	dfn := make([]int, n)
	m := len(edges)
	stack := make([]int, 0, m)
	deg := make([]int, n)
	vis := make([]int, n)
	var timer, label int
	var ans []int

	check := func(comp []int) {
		label++
		var vs []int
		for _, eid := range comp {
			u, v := edges[eid][0]-1, edges[eid][1]-1
			if vis[u] != label {
				vis[u] = label
				vs = append(vs, u)
			}
			if vis[v] != label {
				vis[v] = label
				vs = append(vs, v)
			}
			deg[u]++
			deg[v]++
		}
		ok := len(comp) == len(vs)
		if !ok {
			return
		}
		for _, u := range vs {
			ok = ok && deg[u] == 2
			deg[u] = 0
		}
		if ok {
			for _, eid := range comp {
				ans = append(ans, eid+1)
			}
		}
	}

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		timer++
		dfn[u] = timer
		low[u] = timer

		for _, eid := range adj[u] {
			if eid == p {
				continue
			}
			v := getEdgeTo(eid, u)
			if dfn[v] == 0 {
				stack = append(stack, eid)
				dfs(eid, v)
				low[u] = min(low[u], low[v])
				if low[v] >= dfn[u] {
					var comp []int
					for {
						last := stack[len(stack)-1]
						stack = stack[:len(stack)-1]
						comp = append(comp, last)
						if last == eid {
							break
						}
					}
					check(comp)
				}
			} else if dfn[v] < dfn[u] {
				stack = append(stack, eid)
				low[u] = min(low[u], dfn[v])
			}
		}
	}

	for u := range n {
		if dfn[u] == 0 {
			dfs(-1, u)
		}
	}

	slices.Sort(ans)

	return ans
}
