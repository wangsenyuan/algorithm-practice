package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, v := range res {
		buf.WriteString(fmt.Sprintf("%d\n", v))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n, k, a, b, q int
	fmt.Fscan(reader, &n, &k, &a, &b, &q)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var d, x int
			fmt.Fscan(reader, &d, &x)
			queries[i] = []int{t, d, x}
		} else {
			var p int
			fmt.Fscan(reader, &p)
			queries[i] = []int{t, p}
		}
	}
	return solve(k, n, a, b, queries)
}

func solve(k int, n int, a int, b int, queries [][]int) []int {
	pref := NewTree(n, b)
	suf := NewTree(n, a)

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			d, x := cur[1], cur[2]
			d--
			pref.Update(d, x)
			suf.Update(d, x)
		} else {
			p := cur[1] - 1
			tmp := pref.Query(0, p-1) + suf.Query(p+k, n-1)
			ans = append(ans, tmp)
		}
	}
	return ans
}

type Tree struct {
	val []int
	sz  int
	x   int
}

func NewTree(n int, x int) *Tree {
	return &Tree{
		val: make([]int, 4*n),
		sz:  n,
		x:   x,
	}
}

func (tr *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.val[i] += v
			if tr.val[i] > tr.x {
				tr.val[i] = tr.x
			}
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.val[i] = tr.val[i*2+1] + tr.val[i*2+2]
	}
	f(0, 0, tr.sz-1)
}

func (tr *Tree) Query(L int, R int) int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if r < L || R < l {
			return 0
		}
		if L <= l && r <= R {
			return tr.val[i]
		}
		mid := (l + r) >> 1
		return f(i*2+1, l, mid) + f(i*2+2, mid+1, r)
	}
	return f(0, 0, tr.sz-1)
}
