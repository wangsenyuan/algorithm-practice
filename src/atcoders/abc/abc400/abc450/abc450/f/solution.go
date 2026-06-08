package main

import (
	"bufio"
	"cmp"
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
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	return solve(n, edges)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int, edges [][]int) int {
	slices.SortFunc(edges, func(a []int, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	dp := NewTree(n + 1)

	dp.Add(1, 1)

	for _, cur := range edges {
		l, r := cur[0], cur[1]
		dp.Mul(r, n, 2)
		sum := dp.Get(l, r-1)
		dp.Add(r, sum)
	}

	return dp.Get(n, n)
}

type Tree struct {
	val  []int
	lazy []int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	for i := range 4 * n {
		lazy[i] = 1
	}
	return &Tree{val, lazy}
}

func (tr *Tree) apply(i int, v int) {
	tr.val[i] = mul(tr.val[i], v)
	tr.lazy[i] = mul(tr.lazy[i], v)
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 1 {
		tr.apply(i*2+1, tr.lazy[i])
		tr.apply(i*2+2, tr.lazy[i])
		tr.lazy[i] = 1
	}
}

func (tr *Tree) Mul(L int, R int, v int) {
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
		tr.val[i] = add(tr.val[i*2+1], tr.val[i*2+2])
	}

	n := len(tr.val) / 4
	f(0, 0, n-1, L, R)
}

func (tr *Tree) Add(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = add(tr.val[i], v)
			return
		}

		tr.push(i)
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.val[i] = add(tr.val[i*2+1], tr.val[i*2+2])
	}

	n := len(tr.val) / 4
	f(0, 0, n-1)
}

func (tr *Tree) Get(L int, R int) int {
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
		return add(f(i*2+1, l, mid, L, mid), f(i*2+2, mid+1, r, mid+1, R))
	}

	n := len(tr.val) / 4
	return f(0, 0, n-1, L, R)
}
