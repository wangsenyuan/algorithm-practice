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
	a := readN(reader, n)
	b := readN(reader, n)
	c := readN(reader, n)
	queries := make([][]int, q)
	for i := range q {
		var t int
		fmt.Fscan(reader, &t)
		if t == 4 {
			var pos int
			fmt.Fscan(reader, &pos)
			queries[i] = []int{t, pos}
			continue
		}
		var l, r int
		fmt.Fscan(reader, &l, &r)
		queries[i] = []int{t, l, r}
	}
	return solve(a, b, c, queries)
}

func readN(reader *bufio.Reader, n int) []int {
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return a
}

func solve(a []int, b []int, c []int, queries [][]int) []int {
	// new[i] = old[perm[i]]. The three swaps are transpositions, not one shared rotation.
	perms := [4][3]int{
		{},
		{1, 0, 2},
		{0, 2, 1},
		{2, 1, 0},
	}
	tr := build(a, b, c)
	var res []int
	for _, query := range queries {
		if query[0] == 4 {
			j := tr.firstPostAfter(query[1] - 1)
			if j < 0 {
				j = len(a) + 1
			} else {
				j++
			}
			res = append(res, j)
			continue
		}
		tr.roll(query[1]-1, query[2], perms[query[0]])
	}
	return res
}

type node struct {
	sum    [3]int64
	best   [3]int64 // maximum prefix sum from the left end of this node
	lazy   [3]int
	lazyOn bool
}

func (n *node) apply(p [3]int) {
	var sum, best [3]int64
	for i := range 3 {
		sum[i] = n.sum[p[i]]
		best[i] = n.best[p[i]]
	}
	n.sum = sum
	n.best = best
	if !n.lazyOn {
		n.lazy = p
		n.lazyOn = true
		return
	}
	var composed [3]int
	for i := range 3 {
		composed[i] = n.lazy[p[i]]
	}
	n.lazy = composed
}

func merge(n *node, l *node, r *node) {
	for i := range 3 {
		n.sum[i] = l.sum[i] + r.sum[i]
		n.best[i] = max(l.best[i], l.sum[i]+r.best[i])
	}
}

type tree []node

func build(a []int, b []int, c []int) tree {
	n := len(a)
	t := make(tree, n*4)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			t[i] = node{
				sum:  [3]int64{int64(a[l]), int64(b[l]), int64(c[l])},
				best: [3]int64{int64(a[l]), int64(b[l]), int64(c[l])},
			}
			return
		}
		mid := (l + r) / 2
		f(i*2+1, l, mid)
		f(i*2+2, mid, r)
		merge(&t[i], &t[i*2+1], &t[i*2+2])
	}
	f(0, 0, n)
	return t
}

func (t tree) push(i int) {
	if !t[i].lazyOn {
		return
	}
	left, right := i*2+1, i*2+2
	if right < len(t) {
		t[left].apply(t[i].lazy)
		t[right].apply(t[i].lazy)
	}
	t[i].lazyOn = false
}

func (t tree) roll(L int, R int, perm [3]int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t[i].apply(perm)
			return
		}
		t.push(i)
		mid := (l + r) / 2
		if L < mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid, r, max(L, mid), R)
		}
		merge(&t[i], &t[i*2+1], &t[i*2+2])
	}
	f(0, 0, len(t)/4, L, R)
}

// firstPostAfter returns the smallest 0-based j >= pos with sum(A[pos..j]) > 0, or -1.
func (t tree) firstPostAfter(pos int) int {
	var walk func(i int, l int, r int, need int64) int
	walk = func(i int, l int, r int, need int64) int {
		if l+1 == r {
			return l
		}
		t.push(i)
		mid := (l + r) / 2
		if t[i*2+1].best[0] > need {
			return walk(i*2+1, l, mid, need)
		}
		return walk(i*2+2, mid, r, need-t[i*2+1].sum[0])
	}

	var find func(i int, l int, r int) (int64, int)
	find = func(i int, l int, r int) (int64, int) {
		if r <= pos {
			return 0, -1
		}
		if l+1 == r {
			if t[i].sum[0] > 0 {
				return 0, l
			}
			return t[i].sum[0], -1
		}
		t.push(i)
		mid := (l + r) / 2
		if pos >= mid {
			return find(i*2+2, mid, r)
		}
		leftSum, ans := find(i*2+1, l, mid)
		if ans >= 0 {
			return 0, ans
		}
		if t[i*2+2].best[0]+leftSum > 0 {
			return 0, walk(i*2+2, mid, r, -leftSum)
		}
		return leftSum + t[i*2+2].sum[0], -1
	}
	_, ans := find(0, 0, len(t)/4)
	return ans
}
