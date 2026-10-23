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

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		var t, x, y int
		fmt.Fscan(reader, &t, &x, &y)
		queries[i] = []int{t, x, y}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	tr := build(a)
	var ans []int
	for _, cur := range queries {
		if cur[0] == 1 {
			i, x := cur[1]-1, cur[2]
			tr.update(i, x)
		} else {
			l, r := cur[1]-1, cur[2]-1
			ans = append(ans, tr.query(l, r))
		}
	}
	return ans
}

type node struct {
	pref       int
	suf        int
	firstValue int
	lastValue  int
	best       int
	len        int
}

func merge(a node, b node) node {
	var c node
	c.best = max(a.best, b.best)
	if a.lastValue == b.firstValue {
		c.best = max(c.best, a.suf+b.pref)
	}
	c.len = a.len + b.len
	c.firstValue = a.firstValue
	c.lastValue = b.lastValue
	c.pref = a.pref
	c.suf = b.suf
	if a.pref == a.len && a.lastValue == b.firstValue {
		c.pref += b.pref
	}
	if b.suf == b.len && a.lastValue == b.firstValue {
		c.suf += a.suf
	}

	return c
}

type tree []node

func build(a []int) tree {
	n := len(a)
	tr := make(tree, 4*n)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			tr[i] = node{
				pref:       1,
				suf:        1,
				firstValue: a[l],
				lastValue:  a[l],
				best:       1,
				len:        1,
			}
			return
		}
		m := (l + r) / 2
		f(2*i+1, l, m)
		f(2*i+2, m, r)
		tr[i] = merge(tr[2*i+1], tr[2*i+2])
	}
	f(0, 0, n)
	return tr
}

func (tr tree) update(pos int, val int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			tr[i].firstValue = val
			tr[i].lastValue = val
			return
		}
		m := (l + r) / 2
		if pos < m {
			f(2*i+1, l, m)
		} else {
			f(2*i+2, m, r)
		}
		tr[i] = merge(tr[2*i+1], tr[2*i+2])
	}
	f(0, 0, len(tr)/4)
}

func (tr tree) query(l int, r int) int {
	var f func(i int, l int, r int, L int, R int) node
	f = func(i int, l int, r int, L int, R int) node {
		if l == L && r == R {
			return tr[i]
		}
		mid := (l + r) / 2
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid <= L {
			return f(i*2+2, mid, r, L, R)
		}
		a := f(i*2+1, l, mid, L, mid)
		b := f(i*2+2, mid, r, mid, R)
		return merge(a, b)
	}
	return f(0, 0, len(tr)/4, l, r+1).best
}
