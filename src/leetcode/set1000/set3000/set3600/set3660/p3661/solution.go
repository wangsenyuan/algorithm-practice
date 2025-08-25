package p3661

import (
	"slices"
	"sort"
)

type pair struct {
	first  int
	second int
}

func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{robots[i], distance[i]}
	}
	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})
	slices.Sort(walls)

	destroyLeft := func(x int, d int) []int {
		p1 := sort.SearchInts(walls, x-d)
		p2 := sort.SearchInts(walls, x)
		if p2 == len(walls) || walls[p2] > x {
			p2--
		}
		return []int{p1, p2}
	}
	destroyRight := func(x int, d int) []int {
		p1 := sort.SearchInts(walls, x)
		p2 := sort.SearchInts(walls, x+d)
		if p2 == len(walls) || walls[p2] > x+d {
			p2--
		}
		return []int{p1, p2}
	}

	dp := make([]int, 2)
	ndp := make([]int, 2)

	L := make([][]int, n)
	R := make([][]int, n)
	pos := make([]int, n)

	for i := range n {
		L[i] = destroyLeft(arr[i].first, arr[i].second)
		R[i] = destroyRight(arr[i].first, arr[i].second)
		pos[i] = sort.SearchInts(walls, arr[i].first)
	}

	for i := range n {
		if i == 0 {
			dp[0] = L[i][1] - L[i][0] + 1
			dp[1] = R[i][1] - R[i][0] + 1
			if i+1 < n {
				dp[1] = min(pos[i+1]-1, R[i][1]) - R[i][0] + 1
			}
			continue
		}
		// L[i-1][1]是上一个往左的时候的开始墙的位置
		p1 := max(L[i-1][1]+1, L[i][0])
		ndp[0] = dp[0] + max(0, L[i][1]-p1+1)
		// 或者前面一个往右
		p1 = min(R[i-1][1], pos[i]-1)
		p1 = max(p1+1, L[i][0])
		ndp[0] = max(ndp[0], dp[1]+max(0, L[i][1]-p1+1))
		p2 := R[i][1]
		if i+1 < n {
			p2 = min(p2, pos[i+1]-1)
		}
		ndp[1] = p2 - R[i][0] + 1 + dp[0]
		// 或者前面往右
		ndp[1] = max(ndp[1], p2-R[i][0]+1+dp[1])
		copy(dp, ndp)
	}

	return max(dp[0], dp[1])
}
