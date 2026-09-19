package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	qs := make([][]int, m)
	for i := range m {
		var p, v int
		fmt.Fscan(reader, &p, &v)
		qs[i] = []int{p, v}
	}
	return solve(a, qs)
}

type fenwick []int

func (t fenwick) reset(i int) {
	for ; i < len(t); i += i & -i {
		t[i] = 0
	}
}

func (t fenwick) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] = max(t[i], val)
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, t[i])
	}
	return
}

func solve(a []int, queries [][]int) int {
	type data struct{ i, v, mn, mx, f int }

	n := len(a)

	d := make([]data, n)
	for i, v := range a {
		d[i] = data{i, v, v, v, 1}
	}
	for _, cur := range queries {
		p, v := cur[0], cur[1]
		p--
		d[p].mn = min(d[p].mn, v)
		d[p].mx = max(d[p].mx, v)
	}

	t := make(fenwick, 1e5+1)
	ans := 1
	var play func(int, int)
	play = func(l, r int) {
		if l+1 == r {
			return
		}
		mid := (l + r) >> 1
		play(l, mid)

		slices.SortFunc(d[l:mid], func(a, b data) int { return a.mx - b.mx })
		slices.SortFunc(d[mid:r], func(a, b data) int { return a.v - b.v })
		j := l
		for i := mid; i < r; i++ {
			for ; j < mid && d[j].mx <= d[i].v; j++ {
				t.update(d[j].v, d[j].f)
			}
			d[i].f = max(d[i].f, t.pre(d[i].mn)+1)
			ans = max(ans, d[i].f)
		}
		for j--; j >= l; j-- {
			t.reset(d[j].v)
		}

		slices.SortFunc(d[l:r], func(a, b data) int { return a.i - b.i })

		play(mid, r)
	}

	play(0, n)

	return ans
}
