package p3812

import "slices"

func minimumFlips(n int, edges [][]int, start string, target string) []int {
	color := make([]int, n)
	var tot int
	for i := range n {
		if start[i] != target[i] {
			color[i] = 1
			tot++
		}
	}

	if tot%2 == 1 {
		return []int{-1}
	}

	if tot == 0 {
		return nil
	}

	adj := make([][]int, n)
	for i, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}

	var res []int

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		for _, id := range adj[u] {
			v := edges[id][0] ^ edges[id][1] ^ u
			if p != v {
				dfs(u, v)
				if color[v] == 1 {
					res = append(res, id)
					color[u] ^= 1
					color[v] ^= 1
				}
			}
		}
	}

	for u := range n {
		if color[u] == 1 {
			dfs(-1, u)
			slices.Sort(res)
			return res
		}
	}

	return nil
}
