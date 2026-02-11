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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	p1 := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &p1[i])
	}
	p2 := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &p2[i])
	}
	return solve(n, p1, p2)
}

func solve(n int, fa1 []int, fa2 []int) int {
	adj2 := make([][]int, n)

	for i, p := range fa2 {
		adj2[p-1] = append(adj2[p-1], i+1)
	}

	sz := make([]int, n)
	st := make([]int, n)
	var timer int

	var dfs func(u int)

	at := make([]int, n)

	dfs = func(u int) {
		st[u] = timer
		at[timer] = u
		timer++
		sz[u] = 1
		for _, v := range adj2[u] {
			dfs(v)
			sz[u] += sz[v]
		}
	}

	dfs(0)

	isAnc := func(u int, v int) bool {
		return st[u] < st[v] && st[v] < st[u]+sz[u]
	}

	rec := make([]int, n)

	sub := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})

	fa := NewSegTree(n, -1, func(a, b int) int {
		return max(a, b)
	})

	var cliqueSize int

	insert := func(u int) {
		// 如果存在一个子节点v
		iv := sub.Get(st[u], st[u]+sz[u])
		if iv != n && isAnc(u, at[iv]) {
			// 貌似只要iv存在，就肯定是u的子节点
			rec[u] = -2
			return
		}
		cliqueSize++
		rec[u] = -1
		// 找到u的父节点w
		iw := fa.Get(0, st[u])
		if iw >= 0 && isAnc(at[iw], u) {
			cliqueSize--
			// 要把w删除掉
			fa.Update(iw, -1)
			rec[u] = at[iw]
		}
		sub.Update(st[u], st[u])
		fa.Update(st[u], st[u])
	}

	remove := func(u int) {
		if rec[u] == -2 {
			return
		}
		cliqueSize--
		if rec[u] != -1 {
			cliqueSize++
			w := rec[u]
			fa.Update(st[w], st[w])
		}
		sub.Update(st[u], n)
		fa.Update(st[u], -1)
	}

	var ans int

	adj1 := make([][]int, n)
	for i, p := range fa1 {
		adj1[p-1] = append(adj1[p-1], i+1)
	}

	var dfs1 func(u int)
	dfs1 = func(u int) {
		insert(u)
		ans = max(ans, cliqueSize)
		for _, v := range adj1[u] {
			dfs1(v)
		}
		remove(u)
	}

	dfs1(0)

	return ans
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
