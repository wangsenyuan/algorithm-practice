package p3639

func minTime(s string, order []int, k int) int {
	n := len(order)

	t1 := NewSegTree(n, -1, func(a, b int) int {
		return max(a, b)
	})

	t2 := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})

	var cnt int

	for j, i := range order {
		l := t1.Query(0, i)
		r := t2.Query(i, n)
		cnt += (r - i) * (i - l)
		if cnt >= k {
			return j
		}
		t1.Update(i, i)
		t2.Update(i, i)
	}

	return -1
}

type SegTree struct {
	f   func(int, int) int
	arr []int
	iv  int
	n   int
}

func NewSegTree(n int, iv int, f func(int, int) int) *SegTree {
	arr := make([]int, n*2)
	for i := 0; i < len(arr); i++ {
		arr[i] = iv
	}
	return &SegTree{f, arr, iv, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.n
	t.arr[p] = v

	for p > 1 {
		t.arr[p>>1] = t.f(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Query(l int, r int) int {
	res := t.iv
	l += t.n
	r += t.n
	for l < r {
		if l&1 == 1 {
			res = t.f(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = t.f(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
