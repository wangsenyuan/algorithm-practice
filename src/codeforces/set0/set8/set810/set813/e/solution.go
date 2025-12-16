package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
	"slices"
)

func main() {
	debug.SetGCPercent(-1)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	queries := make([][]int, m)
	for i := range m {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		queries[i] = []int{x, y}
	}
	return solve(n, k, a, queries)
}

func solve(n int, k int, a []int, queries [][]int) []int {
	ax := slices.Max(a)
	pos := make([][]int, ax+1)
	trs := make([]*node, n+1)
	trs[0] = build(n + 1)

	for i, v := range a {
		pos[v] = append(pos[v], i)
		m := len(pos[v])
		trs[i+1] = trs[i]
		if m > k {
			j := pos[v][m-k-1]
			trs[i+1] = trs[i+1].update(j, 1)
		}
	}

	find := func(l int, r int) int {
		if l > r {
			l, r = r, l
		}
		w := r - l + 1
		w -= count(trs[r], l-1)
		return w
	}

	ans := make([]int, len(queries))
	var last int

	for i, cur := range queries {
		x, y := cur[0], cur[1]
		l := (x+last)%n + 1
		r := (y+last)%n + 1
		last = find(l, r)
		ans[i] = last
	}

	return ans
}

type node struct {
	lf, rg *node
	l, r   int
	cnt    int
}

func build(n int) *node {
	var f func(l int, r int) *node
	f = func(l int, r int) *node {
		n := new(node)
		n.l = l
		n.r = r
		if l < r {
			mid := (l + r) >> 1
			n.lf = f(l, mid)
			n.rg = f(mid+1, r)
		}
		return n
	}
	return f(0, n-1)
}

func (a *node) pull() {
	a.cnt = a.lf.cnt + a.rg.cnt
}

func (a node) update(p int, v int) *node {
	if a.l == a.r {
		if v > 0 {
			a.cnt = 1
		} else {
			a.cnt = 0
		}
	} else {
		mid := (a.l + a.r) >> 1
		if p <= mid {
			a.lf = a.lf.update(p, v)
		} else {
			a.rg = a.rg.update(p, v)
		}
		a.pull()
	}

	return &a
}

func count(a *node, p int) int {
	if a.l == a.r {
		// cnt 不应该相减
		return a.cnt
	}
	mid := (a.l + a.r) >> 1

	if p > mid {
		return count(a.rg, p)
	}

	// p <= mid
	res := count(a.lf, p)
	res += a.rg.cnt
	return res
}
