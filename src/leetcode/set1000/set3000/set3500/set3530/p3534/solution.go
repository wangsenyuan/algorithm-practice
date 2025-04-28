package p3534

import (
	"cmp"
	"math/bits"
	"slices"
)

type pair struct {
	first  int
	second int
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{nums[i], i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})
	var m int
	pos := make([]int, n)
	for i := 0; i < n; {
		j := i
		for i < n && arr[i].first == arr[j].first {
			pos[arr[i].second] = m
			i++
		}
		arr[m] = arr[j]
		m++
	}

	arr = arr[:m]
	h := bits.Len(uint(m)) + 1

	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, h)
	}

	for i, r := m-1, m-1; i >= 0; i-- {
		for r > i && arr[r].first-arr[i].first > maxDiff {
			r--
		}
		dp[i][0] = r
		for j := 1; j < h; j++ {
			dp[i][j] = dp[dp[i][j-1]][j-1]
		}
	}

	find := func(u int, v int) int {
		if u == v {
			return 0
		}
		if pos[u] > pos[v] {
			u, v = v, u
		}
		// pos[u] <= pos[v]
		// dp[pos[u]][i] >= pos[v]
		l, r := pos[u], pos[v]
		if dp[l][h-1] < r {
			return -1
		}
		var res int
		for d := h - 1; d >= 0; d-- {
			if dp[l][d] < r {
				res += 1 << d
				l = dp[l][d]
			}
		}
		return res + 1
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur[0], cur[1])
	}

	return ans
}
