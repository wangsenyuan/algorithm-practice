package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	last1 := make([]int, n+1)
	last2 := make([]int, n+2)

	tr := NewTree(n + 1)

	var res int
	for i, v := range a {
		i++
		if last1[v] > 0 {
			tr.Update(last2[v]+1, last1[v], -1)
		}
		tr.Update(last1[v]+1, i, 1)

		res += i + 1 - tr.Query(0, i)

		last1[v], last2[v] = i, last1[v]
	}

	return res
}

type Tree struct {
	val  []int
	sz   []int
	lazy []int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	sz := make([]int, 4*n)
	lazy := make([]int, 4*n)

	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		sz[i] = r - l + 1
		if l < r {
			mid := (l + r) >> 1
			f(i*2+1, l, mid)
			f(i*2+2, mid+1, r)
		}
	}

	f(0, 0, n-1)

	return &Tree{val, sz, lazy}
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
	tr.val[i] = min(tr.val[i*2+1], tr.val[i*2+2])
	tr.sz[i] = 0
	if tr.val[i] == tr.val[i*2+1] {
		tr.sz[i] += tr.sz[i*2+1]
	}
	if tr.val[i] == tr.val[i*2+2] {
		tr.sz[i] += tr.sz[i*2+2]
	}
}

func (tr *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.apply(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(R, mid))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(L, mid+1), R)
		}
		tr.pull(i)
	}
	n := len(tr.val) / 4
	f(0, 0, n-1, L, R)
}

func (tr *Tree) Query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			if tr.val[i] == 0 {
				return tr.sz[i]
			}
			return 0
		}
		tr.push(i)
		mid := (l + r) >> 1
		var res int
		if L <= mid {
			res += f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			res += f(i*2+2, mid+1, r, max(L, mid+1), R)
		}
		return res
	}
	n := len(tr.val) / 4
	return f(0, 0, n-1, L, R)
}
