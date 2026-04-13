package p3901

func countGoodSubseq(nums []int, p int, queries [][]int) int {
	n := len(nums)
	t := NewSegTree(n)

	update := func(i int, v int) {
		if v%p != 0 {
			v = 0
		}

		t.Update(i, v/p)
	}

	for i, v := range nums {
		update(i, v)
	}

	get := func(l int, r int) int {
		if l > r {
			return 0
		}
		return t.Query(l, r)
	}

	check := func() bool {
		for i := range n {
			x := get(0, i)
			y := get(i+1, n)
			if gcd(x, y) == 1 {
				return true
			}
		}
		return false
	}

	var res int

	for _, cur := range queries {
		i, v := cur[0], cur[1]
		update(i, v)
		if t.Query(0, n) == 1 && (n > 6 || check()) {
			// 怎么保证不是有全部组成的呢？
			res++
		}
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type SegTree []int

func NewSegTree(n int) SegTree {
	return make(SegTree, 2*n)
}

func (t SegTree) Update(p int, v int) {
	p += len(t) / 2
	t[p] = v
	for p > 1 {
		t[p>>1] = gcd(t[p], t[p^1])
		p >>= 1
	}
}

func (t SegTree) Query(l int, r int) int {
	var res int
	l += len(t) / 2
	r += len(t) / 2
	for l < r {
		if l&1 == 1 {
			res = gcd(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = gcd(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
