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
	var n, D, R int
	fmt.Fscan(reader, &n, &D, &R)
	H := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &H[i])
	}
	return solve(H, D, R)
}

func solve(H []int, D int, R int) int {
	// H is a permutation of [1...n]
	n := len(H)
	pos := make([]int, n+1)
	for i, v := range H {
		pos[v] = i
	}

	dp := make([]int, n+1)
	tr := NewTree(n)
	for v := n; v > 0; v-- {
		if v+D <= n {
			tr.Update(pos[v+D], dp[v+D])
		}
		l := max(0, pos[v]-R)
		r := min(n-1, pos[v]+R)
		dp[v] = tr.Query(l, r) + 1
	}

	return slices.Max(dp) - 1
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	return &Tree{
		val:  make([]int, 4*n),
		lazy: make([]int, 4*n),
		sz:   n,
	}
}

func (tr *Tree) apply(i int, v int) {
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

func (tr *Tree) pull(i int) {
	tr.val[i] = max(tr.val[i*2+1], tr.val[i*2+2])
}

func (tr *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = v
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.pull(i)
	}
	f(0, 0, tr.sz-1)
}

func (tr *Tree) Query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return tr.val[i]
		}
		tr.push(i)
		mid := (l + r) >> 1
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return f(i*2+2, mid+1, r, L, R)
		}
		return max(f(i*2+1, l, mid, L, mid), f(i*2+2, mid+1, r, mid+1, R))
	}
	return f(0, 0, tr.sz-1, L, R)
}
