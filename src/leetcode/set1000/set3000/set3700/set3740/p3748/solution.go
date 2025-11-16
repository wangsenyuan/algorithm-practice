package p3748

import (
	"cmp"
	"math"
	"slices"
)

func countStableSubarrays(nums []int, queries [][]int) []int64 {
	n := len(nums)
	sum := make([]int64, n+1)
	var cnt int
	for i := range n {
		if i == 0 || nums[i] < nums[i-1] {
			cnt = 0
		}
		cnt++
		sum[i+1] = sum[i] + int64(cnt)
	}

	next := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || nums[i] > nums[i+1] {
			next[i] = i + 1
		} else {
			next[i] = next[i+1]
		}
	}

	ans := make([]int64, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		l2 := next[l]
		if l2 > r {
			m := int64(r - l + 1)
			ans[i] = m * (m + 1) / 2
		} else {
			ans[i] = sum[r+1] - sum[l2]
			m := int64(l2 - l)
			ans[i] += m * (m + 1) / 2
		}
	}
	return ans
}

type pair struct {
	first  int
	second int
}

type query struct {
	l  int
	r  int
	id int
}

func countStableSubarrays1(nums []int, queries [][]int) []int64 {
	n := len(nums)
	L := make([]int, n)

	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{nums[i], i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(b.first-a.first, b.second-a.second)
	})

	tr := NewSegTree(n, -1, func(a, b int) int {
		return max(a, b)
	})

	for _, cur := range arr {
		i := cur.second
		// 左边第一个大于nums[i]的位置
		L[i] = tr.Get(0, i)
		tr.Update(i, i)
	}

	for i := 1; i < n; i++ {
		L[i] = max(L[i], L[i-1])
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	R := make([]int, n)

	tr = NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})

	for _, cur := range arr {
		i := cur.second
		// 右边第一个小于nums[i]的位置
		R[i] = tr.Get(i, n)
		tr.Update(i, i)
	}

	for i := n - 2; i >= 0; i-- {
		R[i] = min(R[i+1], R[i])
	}

	qs := make([]query, len(queries))
	for i, cur := range queries {
		qs[i] = query{cur[0], cur[1], i}
	}

	block_size := int(math.Sqrt(float64(n)))

	slices.SortFunc(qs, func(a, b query) int {
		if a.r/block_size != b.r/block_size {
			return a.r - b.r
		}
		if (a.r/block_size)%2 == 0 {
			return a.l - b.l
		}
		return b.l - a.l
	})

	ans := make([]int64, len(queries))

	var sum int
	var l, r int

	add := func(i int, left bool) {
		if left {
			sum += min(r, R[i]) - i
		} else {
			sum += i - max(l-1, L[i])
		}
	}

	rem := func(i int, left bool) {
		if left {
			sum -= min(r, R[i]) - i
		} else {
			sum -= i - max(l-1, L[i])
		}
	}

	for _, cur := range qs {
		for r <= cur.r {
			add(r, false)
			r++
		}
		for r-1 > cur.r {
			r--
			rem(r, false)
		}

		for l > cur.l {
			l--
			add(l, true)
		}
		for l < cur.l {
			rem(l, true)
			l++
		}
		ans[cur.id] = int64(sum)
	}

	return ans
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
