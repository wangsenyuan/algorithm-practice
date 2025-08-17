package p3651

import (
	"slices"
	"sort"
)

const inf = 1 << 60

func minCost(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	// 就运行k层就好了
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
		for j := range m {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0

	type data struct {
		r int
		c int
		u int
		v int
	}

	arr := make([]data, n*m)
	for i := range n {
		for j := range m {
			arr[i*m+j] = data{i, j, grid[i][j], inf}
		}
	}

	arr[0].v = 0

	slices.SortFunc(arr, func(a, b data) int {
		return a.u - b.u
	})

	for i := len(arr) - 2; i >= 0; i-- {
		arr[i].v = min(arr[i].v, arr[i+1].v)
	}

	for x := range k + 1 {
		for i := range n {
			for j := range m {
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j]+grid[i][j])
				}
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1]+grid[i][j])
				}
				if x > 0 {
					w := sort.Search(len(arr), func(w int) bool {
						return arr[w].u >= grid[i][j]
					})
					if w < len(arr) {
						dp[i][j] = min(dp[i][j], arr[w].v)
					}
				}
			}
		}
		for i := len(arr) - 1; i >= 0; i-- {
			r, c := arr[i].r, arr[i].c
			arr[i].v = dp[r][c]
			if i+1 < len(arr) {
				arr[i].v = min(arr[i].v, arr[i+1].v)
			}
		}
	}

	return dp[n-1][m-1]
}
