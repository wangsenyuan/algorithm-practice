package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
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
	suf := make([]int, n+1)
	var sufCount int
	for _, v := range a {
		suf[v]++
		if suf[v] == 1 {
			sufCount++
		}
	}

	last := make([]int, n+1)
	for i := range n + 1 {
		last[i] = -1
	}
	var prefCount int

	var res int

	tr := Build(n)

	for i, v := range a {
		suf[v]--
		if suf[v] == 0 {
			sufCount--
		}

		if j := last[v]; j >= 0 {
			tr.update(j, i-1, 1)
		} else {
			prefCount++
		}
		last[v] = i
		tmp := prefCount + sufCount + tr.query(0, i)
		res = max(res, tmp)
	}

	return res
}

type Tree struct {
	val  []int
	lazy []int
}

func Build(n int) *Tree {
	val := make([]int, n*4)
	lazy := make([]int, n*4)
	return &Tree{val, lazy}
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

func (tr *Tree) update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if L == l && r == R {
			tr.apply(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) / 2
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		tr.pull(i)
	}
	n := len(tr.val) / 4
	f(0, 0, n-1, L, R)
}

func (tr *Tree) query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if L == l && r == R {
			return tr.val[i]
		}
		tr.push(i)
		mid := (l + r) / 2
		var res int
		if L <= mid {
			res = max(res, f(i*2+1, l, mid, L, min(mid, R)))
		}
		if mid < R {
			res = max(res, f(i*2+2, mid+1, r, max(mid+1, L), R))
		}
		return res
	}

	return f(0, 0, len(tr.val)/4-1, L, R)
}
