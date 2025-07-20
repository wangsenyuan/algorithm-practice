package p3624

import "math/bits"

var f [61]int

func init() {
	f[1] = 0

	for i := 2; i <= 60; i++ {
		j := bits.OnesCount(uint(i))
		f[i] = f[j] + 1
	}
}

func popcountDepth(nums []int64, queries [][]int64) []int {
	tr := NewTree(nums)

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			l, r, k := int(cur[1]), int(cur[2]), int(cur[3])
			tmp := tr.Query(l, r)
			ans = append(ans, tmp[k])
		} else {
			i, v := int(cur[1]), int(cur[2])
			tr.Update(i, v)
		}
	}
	return ans
}

func g(x int) int {
	if x < 61 {
		return f[x]
	}
	return g(bits.OnesCount(uint(x))) + 1
}

type Tree struct {
	cnt [][]int
	val []int
	sz  int
}

func merge(c []int, a []int, b []int) []int {
	for i := range 6 {
		c[i] = a[i] + b[i]
	}
	return c
}

func NewTree(arr []int64) *Tree {
	n := len(arr)
	cnt := make([][]int, 4*n)
	val := make([]int, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		cnt[i] = make([]int, 6)
		if l == r {
			val[i] = int(arr[l])
			cnt[i][g(val[i])]++
			return
		}
		mid := (l + r) >> 1
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)

		merge(cnt[i], cnt[i*2+1], cnt[i*2+2])
	}
	build(0, 0, n-1)
	return &Tree{cnt, val, n}
}

func (tr *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.cnt[i][g(tr.val[i])]--
			tr.val[i] = v
			tr.cnt[i][g(tr.val[i])]++
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			loop(i*2+1, l, mid)
		} else {
			loop(i*2+2, mid+1, r)
		}
		merge(tr.cnt[i], tr.cnt[i*2+1], tr.cnt[i*2+2])
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) Query(L int, R int) []int {
	var loop func(i int, l int, r int, L int, R int) []int
	loop = func(i int, l int, r int, L int, R int) []int {
		if L == l && R == r {
			return tr.cnt[i]
		}
		mid := (l + r) >> 1
		if R <= mid {
			return loop(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return loop(i*2+2, mid+1, r, L, R)
		}
		a := loop(i*2+1, l, mid, L, mid)
		b := loop(i*2+2, mid+1, r, mid+1, R)
		c := make([]int, 6)
		merge(c, a, b)
		return c
	}
	return loop(0, 0, tr.sz-1, L, R)
}
