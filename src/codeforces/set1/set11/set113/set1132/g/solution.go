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
	for i, v := range res {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) []int {
	n := len(a)
	firstChild := make([]int, n+1)
	nextSibling := make([]int, n+1)
	for i := 0; i <= n; i++ {
		firstChild[i] = -1
		nextSibling[i] = -1
	}
	fa := make([]int, n+1)
	stack := make([]int, n)
	var top int
	for i := range n {
		for top > 0 && a[stack[top-1]] < a[i] {
			v := stack[top-1]
			fa[v] = i
			nextSibling[v] = firstChild[i]
			firstChild[i] = v
			top--
		}
		stack[top] = i
		top++
	}
	for top > 0 {
		v := stack[top-1]
		fa[v] = n
		nextSibling[v] = firstChild[n]
		firstChild[n] = v
		top--
	}
	dfn := make([]int, n+1)
	sz := make([]int, n+1)
	ord := make([]int, 0, n+1)
	stack = append(stack[:0], n)
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		dfn[u] = len(ord)
		ord = append(ord, u)
		for v := firstChild[u]; v >= 0; v = nextSibling[v] {
			stack = append(stack, v)
		}
	}
	fa[n] = n
	for i := len(ord) - 1; i >= 0; i-- {
		u := ord[i]
		sz[u]++
		if u != n {
			sz[fa[u]] += sz[u]
		}
	}

	ans := make([]int, n-k+1)

	dp := BuildTree(n + 1)

	for i := n - 1; i >= 0; i-- {
		if i+k < n {
			j := dfn[i+k]
			dp.update(j, j+sz[i+k]-1, -1)
		}
		var w int32
		if fa[i] < min(n, i+k) {
			w = dp.get(dfn[fa[i]], dfn[fa[i]])
		}
		w++
		dp.set(dfn[i], w)
		if i < len(ans) {
			ans[i] = int(dp.get(0, n))
		}
	}

	return ans
}

type Tree struct {
	val  []int32
	lazy []int32
}

func BuildTree(n int) *Tree {
	val := make([]int32, 4*n)
	lazy := make([]int32, 4*n)
	return &Tree{val, lazy}
}

func (tr *Tree) apply(i int, v int32) {
	tr.val[i] += v
	tr.lazy[i] += v
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.apply(i*2+1, tr.lazy[i])
		tr.apply(i*2+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *Tree) update(L int, R int, v int32) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.apply(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(L, mid+1), R)
		}
		tr.val[i] = max(tr.val[i*2+1], tr.val[i*2+2])
	}
	n := len(tr.val) / 4
	f(0, 0, n-1, L, R)
}

func (tr *Tree) set(p int, v int32) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = v
			tr.lazy[i] = 0
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.val[i] = max(tr.val[i*2+1], tr.val[i*2+2])
	}
	n := len(tr.val) / 4
	f(0, 0, n-1)
}

func (tr *Tree) get(L int, R int) int32 {
	var f func(i int, l int, r int, L int, R int) int32
	f = func(i int, l int, r int, L int, R int) int32 {
		if l == L && r == R {
			return tr.val[i]
		}
		tr.push(i)
		var res int32
		mid := (l + r) >> 1
		if L <= mid {
			res = max(res, f(i*2+1, l, mid, L, min(mid, R)))
		}
		if mid < R {
			res = max(res, f(i*2+2, mid+1, r, max(mid+1, L), R))
		}
		return res
	}
	n := len(tr.val) / 4
	return f(0, 0, n-1, L, R)
}
