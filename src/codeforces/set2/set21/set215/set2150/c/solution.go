package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int

	fmt.Fscan(reader, &n)
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(c, a, b)
}

func solve(c []int, a []int, b []int) int {
	n := len(c)
	pb := make([]int, n+1)
	for i, v := range b {
		pb[v] = i + 1
	}

	p := make([]int, n+1)
	for i := range n {
		p[i+1] = pb[a[i]]
	}

	v := make([]int, n+1)
	for i := range n {
		v[i+1] = c[a[i]-1]
	}

	tr := NewTree(n + 1)

	tr.chmax(0, 0)

	for i := 1; i <= n; i++ {
		val := tr.query(0, p[i]-1)
		tr.chmax(p[i], val)
		// j < p[i], update with v[i]
		tr.update(0, p[i]-1, v[i])
	}

	return tr.query(0, n)
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

const inf = 1 << 60

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	for i := range 4 * n {
		val[i] = -inf
	}
	lazy := make([]int, 4*n)
	return &Tree{val, lazy, n}
}

func (t *Tree) apply(i int, v int) {
	t.val[i] += v
	t.lazy[i] += v
}

func (t *Tree) push(i int) {
	if t.lazy[i] != 0 {
		t.apply(2*i+1, t.lazy[i])
		t.apply(2*i+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) pull(i int) {
	t.val[i] = max(t.val[2*i+1], t.val[2*i+2])
}

func (t *Tree) update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.apply(i, v)
			return
		}
		t.push(i)
		mid := (l + r) / 2
		if L <= mid {
			f(2*i+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(2*i+2, mid+1, r, max(mid+1, L), R)
		}
		t.pull(i)
	}

	f(0, 0, t.sz-1, L, R)
}

func (t *Tree) chmax(pos int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			t.val[i] = max(t.val[i], v)
			return
		}
		t.push(i)
		mid := (l + r) / 2
		if pos <= mid {
			f(2*i+1, l, mid)
		} else {
			f(2*i+2, mid+1, r)
		}
		t.pull(i)
	}

	f(0, 0, t.sz-1)
}

func (t *Tree) query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return t.val[i]
		}
		t.push(i)
		mid := (l + r) / 2
		res := -inf
		if L <= mid {
			res = max(res, f(2*i+1, l, mid, L, min(mid, R)))
		}
		if mid < R {
			res = max(res, f(2*i+2, mid+1, r, max(mid+1, L), R))
		}
		return res
	}

	return f(0, 0, t.sz-1, L, R)
}
