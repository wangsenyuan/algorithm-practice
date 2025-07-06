package p3605

import "sort"

func minStable(nums []int, maxC int) int {

	n := len(nums)

	tr := NewSegTree(nums)

	check := func(k int) bool {
		var cnt int
		var l int
		var g int
		for i := 0; i < n; i++ {
			g = gcd(g, nums[i])
			if g == 1 {
				for g == 1 {
					l++
					g = tr.Query(l, i+1)
				}
				continue
			}

			// g > 1
			if i-l == k {
				cnt++
				l = i + 1
				g = 0
			}
		}
		return cnt <= maxC
	}
	res := sort.Search(n, check)
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(nums []int) *SegTree {
	n := len(nums)
	arr := make([]int, 2*n)
	copy(arr[n:], nums)
	for i := n - 1; i > 0; i-- {
		arr[i] = gcd(arr[i<<1], arr[i<<1|1])
	}

	return &SegTree{arr, n}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.sz
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = gcd(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Query(l int, r int) int {
	l += seg.sz
	r += seg.sz
	var res int
	for l < r {
		if l&1 == 1 {
			res = gcd(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = gcd(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
