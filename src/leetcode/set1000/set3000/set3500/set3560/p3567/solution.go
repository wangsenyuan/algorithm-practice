package p3567

import "sort"

func minAbsDiff(grid [][]int, k int) [][]int {
	n := len(grid)
	m := len(grid[0])

	ans := make([][]int, n-k+1)
	for i := range ans {
		ans[i] = make([]int, m-k+1)
	}

	for i := 0; i < n-k+1; i++ {
		for j := 0; j < m-k+1; j++ {
			var arr []int
			for u := 0; u < k; u++ {
				for v := 0; v < k; v++ {
					arr = append(arr, grid[i+u][j+v])
				}
			}
			sort.Ints(arr)
			ans[i][j] = inf
			for k := 0; k+1 < len(arr); k++ {
				if arr[k] != arr[k+1] {
					ans[i][j] = min(ans[i][j], arr[k+1]-arr[k])
				}
			}
			if ans[i][j] == inf {
				ans[i][j] = 0
			}
		}
	}

	return ans
}

const inf = 1 << 60
