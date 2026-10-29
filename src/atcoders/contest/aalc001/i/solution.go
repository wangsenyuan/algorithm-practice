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

	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int64 {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var pos, x int
			fmt.Fscan(reader, &pos, &x)
			queries[i] = []int{t, pos, x}
		} else {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			queries[i] = []int{t, l, r}
		}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int64 {
	t := build(a)
	var res []int64

	for _, cur := range queries {
		if cur[0] == 1 {
			i, x := cur[1], cur[2]
			t.update(i-1, x)
		} else {
			l, r := cur[1], cur[2]
			res = append(res, int64(t.query(l-1, r)))
		}
	}
	return res
}

type node struct {
	best     int
	best1    int
	best2    int
	sum      int
	pref     int
	suf      int
	bestPref int
	bestSuf  int // 包含右端点的最优解
}

func merge(left, right node) node {
	var res node
	res.best = max(left.best, right.best, left.best1+right.best1, left.bestSuf+right.pref, left.suf+right.bestPref)
	res.best1 = max(left.best1, right.best1, left.suf+right.pref)
	res.sum = left.sum + right.sum
	res.pref = max(left.pref, left.sum+right.pref)
	res.suf = max(right.suf, right.sum+left.suf)

	// bestPref：从左端点开始的一段，再加它右边至多一段。best2 把右端点的一段接到右孩子的前缀上。
	res.bestPref = max(left.bestPref, left.sum+right.bestPref, left.pref+right.best1, left.best2+max(0, right.pref))
	res.bestSuf = max(right.bestSuf, right.sum+left.bestSuf, right.suf+left.best1, right.best2+max(0, left.suf))

	// best2：左端点的一段加右端点的一段。右端点那段要覆盖整个右孩子，只接前缀就到不了右端点。
	res.best2 = max(left.pref+right.suf, left.best2+right.sum, left.sum+right.best2, res.pref, res.suf)

	return res
}

type tree []node

func build(a []int) tree {
	n := len(a)
	t := make(tree, 4*n)

	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			t[i] = node{max(0, a[l]), max(0, a[l]), max(0, a[l]), a[l], max(0, a[l]), max(0, a[l]), max(0, a[l]), max(0, a[l])}
			return
		}
		mid := (l + r) / 2
		f(i*2+1, l, mid)
		f(i*2+2, mid, r)
		t[i] = merge(t[i*2+1], t[i*2+2])
	}

	f(0, 0, n)

	return t
}

func (t tree) update(pos int, x int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			t[i] = node{max(0, x), max(0, x), max(0, x), x, max(0, x), max(0, x), max(0, x), max(0, x)}
			return
		}
		mid := (l + r) / 2
		if pos < mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid, r)
		}
		t[i] = merge(t[i*2+1], t[i*2+2])
	}
	f(0, 0, len(t)/4)
}

func (t tree) query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) node
	f = func(i int, l int, r int, L int, R int) node {
		if l == L && r == R {
			return t[i]
		}
		mid := (l + r) / 2
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid <= L {
			return f(i*2+2, mid, r, L, R)
		}
		return merge(f(i*2+1, l, mid, L, mid), f(i*2+2, mid, r, mid, R))
	}
	res := f(0, 0, len(t)/4, L, R)
	return res.best
}
