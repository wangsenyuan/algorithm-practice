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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}
	foe := make([][]int, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		foe[i] = []int{a, b}
	}
	return solve(p, foe)
}

func solve(p []int, foe [][]int) int {
	n := len(p)

	pos := make([]int, n)
	for i := range n {
		pos[p[i]-1] = i
	}

	pair := make([]int, n)
	for i := range n {
		pair[i] = -1
	}
	for _, cur := range foe {
		a, b := pos[cur[0]-1], pos[cur[1]-1]
		if a > b {
			a, b = b, a
		}
		pair[b] = max(pair[b], a)
	}

	tr := NewSegTree(n)

	var res int
	for l, r := 0, 0; r < n; r++ {
		if pair[r] >= 0 {
			tr.Update(r, pair[r])
		}
		for l < r && tr.Get(l, r+1) >= l {
			l++
		}
		res += r - l + 1
	}

	return res
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = -1
	}
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = max(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	res := -1
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res = max(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
