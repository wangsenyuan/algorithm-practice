package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	particles := make([][]int, n)
	for i := range n {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		particles[i] = []int{x, y}
	}
	shop := make([][]int, m)
	for i := range m {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		shop[i] = []int{x, y}
	}
	return solve(particles, shop)
}

func solve(particles [][]int, shop [][]int) []int {

	n := len(particles)
	var xs []int

	ys := make([][]int, n+1)

	for _, cur := range particles {
		xs = append(xs, cur[0])
		ys[cur[1]] = append(ys[cur[1]], cur[0])
	}
	slices.Sort(xs)
	xs = slices.Compact(xs)

	tr := NewTree(len(xs))

	dp := make(SegTree, 2*(n+1))

	var best int
	for y := n; y >= 0; y-- {
		for _, x := range ys[y] {
			i := sort.SearchInts(xs, x)
			tr.Update(i, x)
		}
		best = max(best, tr.GetBest(y+1))
		dp.Update(y, tr.GetBest(y))
	}

	ans := make([]int, len(shop))
	for i, cur := range shop {
		x, y := cur[0], cur[1]
		ans[i] = max(best, x+dp.Get(0, y+1))
	}

	return ans
}

type SegTree []int

func (t SegTree) Update(p int, v int) {
	p += len(t) / 2
	t[p] = max(t[p], v)
	for p > 1 {
		t[p>>1] = max(t[p], t[p^1])
		p >>= 1
	}
}

func (t SegTree) Get(l int, r int) int {
	n := len(t) / 2
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

type Tree struct {
	cnt []int
	val []int
	sz  int
}

func NewTree(n int) *Tree {
	cnt := make([]int, 4*n)
	val := make([]int, 4*n)
	return &Tree{cnt, val, n}
}

func (t *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			t.cnt[i]++
			t.val[i] += v
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(2*i+1, l, mid)
		} else {
			f(2*i+2, mid+1, r)
		}
		t.cnt[i] = t.cnt[2*i+1] + t.cnt[2*i+2]
		t.val[i] = t.val[2*i+1] + t.val[2*i+2]
	}
	f(0, 0, t.sz-1)
}

func (t *Tree) GetBest(k int) int {
	if k == 0 {
		return 0
	}
	if k >= t.cnt[0] {
		return t.val[0]
	}

	var f func(i int, l int, r int, k int) int
	f = func(i int, l int, r int, k int) int {
		if l == r {
			return t.val[i] / t.cnt[i] * k
		}
		mid := (l + r) >> 1
		if t.cnt[2*i+2] >= k {
			return f(2*i+2, mid+1, r, k)
		}
		return t.val[2*i+2] + f(2*i+1, l, mid, k-t.cnt[2*i+2])
	}
	return f(0, 0, t.sz-1, k)
}
