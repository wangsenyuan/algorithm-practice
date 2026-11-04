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

	s := fmt.Sprintf("%v", drive(reader))
	fmt.Fprintln(writer, s[1:len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, n-1)
	for i := range edges {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	ops := make([][]int, m)
	for i := range ops {
		ops[i] = make([]int, 2)
		fmt.Fscan(reader, &ops[i][0], &ops[i][1])
	}
	return solve(n, edges, ops)
}

func solve(n int, edges [][]int, ops [][]int) []int {
	g := make([][]int, n)
	for _, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	in := make([]int, n)
	sz := make([]int, n)
	var timer int
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		in[u] = timer
		timer++
		sz[u]++
		for _, v := range g[u] {
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0)

	todo := make([][]int, n)
	for i, cur := range ops {
		a, b := cur[0]-1, cur[1]-1
		todo[a] = append(todo[a], i)
		todo[b] = append(todo[b], i)
	}

	c := make([]int, n)

	tr := build(n)

	var dfs2 func(p int, u int, flag int)
	dfs2 = func(p int, u int, flag int) {
		for _, i := range todo[u] {
			flag = 1
			a, b := ops[i][0]-1, ops[i][1]-1
			tr.update(in[a], in[a]+sz[a], 1)
			tr.update(in[b], in[b]+sz[b], 1)
		}
		if flag > 0 {
			covered := n
			if tr.val[1] == 0 {
				covered = n - tr.ans[1]
			}
			c[u] = covered - 1
		}
		for _, v := range g[u] {
			if p != v {
				dfs2(u, v, flag)
			}
		}
		for _, i := range todo[u] {
			a, b := ops[i][0]-1, ops[i][1]-1
			tr.update(in[a], in[a]+sz[a], -1)
			tr.update(in[b], in[b]+sz[b], -1)
		}
	}

	dfs2(-1, 0, 0)

	return c
}

type tree struct {
	val []int
	ans []int
	tag []int
	sz  int
}

func build(n int) *tree {
	tr := &tree{
		val: make([]int, 4*n),
		ans: make([]int, 4*n),
		tag: make([]int, 4*n),
		sz:  n,
	}

	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		tr.ans[i] = r - l
		if l+1 < r {
			mid := (l + r) / 2
			f(i*2, l, mid)
			f(i*2+1, mid, r)
		}
	}
	f(1, 0, n)
	return tr
}

func (tr *tree) apply(i int, v int) {
	// val[i] 是这一段的最小值
	tr.val[i] += v
	tr.tag[i] += v
}

func (tr *tree) merge(i int) {
	l := 2 * i
	r := l + 1
	tr.val[i] = min(tr.val[l], tr.val[r])
	if tr.val[l] < tr.val[r] {
		tr.ans[i] = tr.ans[l]
	} else if tr.val[l] > tr.val[r] {
		tr.ans[i] = tr.ans[r]
	} else {
		tr.ans[i] = tr.ans[l] + tr.ans[r]
	}
}

func (tr *tree) push(i int) {
	if tr.tag[i] != 0 {
		tr.apply(i*2, tr.tag[i])
		tr.apply(i*2+1, tr.tag[i])
		tr.tag[i] = 0
	}
}

func (tr *tree) update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.apply(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if L < mid {
			f(i*2, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+1, mid, r, max(L, mid), R)
		}
		tr.merge(i)
	}
	f(1, 0, tr.sz, L, R)
}
