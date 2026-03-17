package main

import "slices"

func main() {

}

func solve(n int, t int, relations [][]int) int {
	slices.SortFunc(relations, func(a, b []int) int {
		return a[2] - b[2]
	})

	adj := make([][]int, n)
	for i, relation := range relations {
		h := relation[0] - 1
		adj[h] = append(adj[h], i)
	}

	// 找出最大的n个边
	var tot int
	k := len(relations)
	for i := max(0, k-n); i < k; i++ {
		tot += relations[i][2]
	}

	dp := make([]int, tot+1)
	dp[0] = 1

	var path []int

	vis := make([]bool, n)
	pair := make([]int, n)
	for i := range n {
		pair[i] = -1
	}

	var dfs func(u int) bool
	dfs = func(u int) bool {
		if vis[u] {
			return false
		}
		vis[u] = true
		for _, i := range adj[u] {
			v := (relations[i][0] - 1) ^ (relations[i][1] - 1) ^ u
			if pair[v] == -1 || dfs(pair[v]) {
				path = append(path, relations[i][2])
				pair[v] = u
				return true
			}
		}
		return false
	}

	for _, cur := range relations {
		h, w, r := cur[0]-1, cur[1]-1, cur[2]

		path = path[:0]
		clear(vis)
		// 这里要怎么强制使用这条边呢？

	}
}
