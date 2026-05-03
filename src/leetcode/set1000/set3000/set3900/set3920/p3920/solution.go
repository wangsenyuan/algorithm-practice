package p3920

import (
	"cmp"
	"slices"
)

func maxFixedPoints(nums []int) int {
	n := len(nums)
	// 对于当前(i, v)来说，要找到左边(j, w), w < v, v - w <= i - j
	// i - v >= j - w (这个范围内的最大值)
	// 但是这里没法保证 v < w
	// 按照v升序处理
	type pair struct {
		first  int
		second int
	}

	arr := make([]pair, n)
	for i, v := range nums {
		arr[i] = pair{v, i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	dp := make(SegTree, 2*n)

	for i := 0; i < n; {
		j := i
		var todo []pair
		for i < n && arr[i].first == arr[j].first {
			if arr[i].first <= arr[i].second {
				d := arr[i].second - arr[i].first
				v := dp.Get(0, d+1) + 1
				todo = append(todo, pair{i, v})
			}
			i++
		}
		for _, cur := range todo {
			i, v := cur.first, cur.second
			d := arr[i].second - arr[i].first
			dp.Set(d, v)
		}
	}

	return dp.Get(0, n)
}

type SegTree []int

func (t SegTree) Set(p int, v int) {
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
