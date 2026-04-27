package p3915

import (
	"slices"
	"sort"
)

func maxAlternatingSum(nums []int, k int) int64 {
	// dp[i][0]表示到i为止, 以i为峰时的最大值
	// dp[i][1] 表示到i为止，以i为谷时的最优解
	arr := slices.Clone(nums)
	slices.Sort(arr)
	arr = slices.Compact(arr)
	m := len(arr)
	t1 := make(SegTree, 2*m)
	t2 := make(SegTree, 2*m)

	n := len(nums)
	dp := make([][2]int, n)

	for i, v := range nums {
		if i-k >= 0 {
			w := nums[i-k]
			j := sort.SearchInts(arr, w)
			y := t1.Get(j, j+1)
			if y < dp[i-k][0] {
				t1.Update(j, dp[i-k][0])
			}
			y = t2.Get(j, j+1)
			if y < dp[i-k][1] {
				t2.Update(j, dp[i-k][1])
			}
		}
		j := sort.SearchInts(arr, v)
		dp[i][0] = t2.Get(0, j) + v
		dp[i][1] = t1.Get(j+1, m) + v
	}

	var res int
	for i := range n {
		res = max(res, dp[i][0], dp[i][1])
	}

	return int64(res)
}

type SegTree []int

func (t SegTree) Update(p int, v int) {
	n := len(t) / 2
	p += n
	t[p] = v
	for p > 1 {
		t[p>>1] = max(t[p], t[p^1])
		p >>= 1
	}
}

func (t SegTree) Get(l int, r int) int {
	n := len(t) / 2
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
