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

	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		var t, l, r int
		fmt.Fscan(reader, &t, &l, &r)

		queries[i] = []int{t, l, r}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	trs := make([]tree, 30)
	for d := range 30 {
		trs[d] = build(d, a)
	}

	update := func(pos int, x int) {
		for d := range 30 {
			trs[d].update(pos, (x>>d)&1)
		}
		a[pos] = x
	}
	query := func(l int, r int) int {
		if l == r {
			return a[l]
		}
		var ans int
		for d := range 30 {
			n := trs[d].query(l+1, r+1)
			ans |= n.apply((a[l]>>d)&1) << d
		}
		return ans
	}
	var res []int
	for _, cur := range queries {
		if cur[0] == 1 {
			update(cur[1]-1, cur[2])
		} else {
			res = append(res, query(cur[1]-1, cur[2]-1))
		}
	}
	return res
}

type node struct {
	f0 int
	f1 int
}

func (g node) apply(s int) int {
	if s == 0 {
		return g.f0
	}
	return g.f1
}

func compose(a node, b node) node {
	return node{b.apply(a.f0), b.apply(a.f1)}
}

type tree []node

func build(d int, a []int) tree {
	tr := make([]node, len(a)*4)

	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			tr[i] = node{1, 1 - (a[l]>>d)&1}
			return
		}
		mid := (l + r) >> 1
		f(i*2+1, l, mid)
		f(i*2+2, mid, r)
		tr[i] = compose(tr[i*2+1], tr[i*2+2])
	}
	f(0, 0, len(a))
	return tr
}

// v (0, 1)
func (tr tree) update(pos int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			tr[i] = node{1, 1 - v}
			return
		}
		mid := (l + r) >> 1
		if pos < mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid, r)
		}
		tr[i] = compose(tr[i*2+1], tr[i*2+2])
	}
	f(0, 0, len(tr)/4)
}

func (tr tree) query(L int, R int) node {
	var f func(i int, l int, r int, L int, R int) node
	f = func(i int, l int, r int, L int, R int) node {
		if l == L && r == R {
			return tr[i]
		}
		mid := (l + r) >> 1
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid <= L {
			return f(i*2+2, mid, r, L, R)
		}
		return compose(f(i*2+1, l, mid, L, mid), f(i*2+2, mid, r, mid, R))
	}
	return f(0, 0, len(tr)/4, L, R)
}
