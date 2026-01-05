package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(a, queries)
}

const mod = 95542721

func mul(nums ...int) int {
	res := 1
	for _, num := range nums {
		res = res * num % mod
	}
	return res
}

func add(nums ...int) int {
	res := 0
	for _, num := range nums {
		res = (res + num) % mod
	}
	return res
}

func solve(a []int, queries [][]int) []int {
	tr := NewTree(a)
	var ans []int
	for _, cur := range queries {
		l, r := cur[1]-1, cur[2]-1
		if cur[0] == 1 {
			ans = append(ans, tr.Query(l, r))
		} else {
			tr.Update(l, r, 1)
		}
	}
	return ans
}

type Tree struct {
	sum  [][48]int
	lazy []int
	sz   int
}

func NewTree(a []int) *Tree {
	n := len(a)
	sum := make([][48]int, 4*n)
	lazy := make([]int, 4*n)
	sz := n
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			sum[i][0] = a[l]
			for j := 1; j < 48; j++ {
				sum[i][j] = mul(sum[i][j-1], sum[i][j-1], sum[i][j-1])
			}
			return
		}

		mid := (l + r) >> 1
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)

		for j := range 48 {
			sum[i][j] = add(sum[i*2+1][j], sum[i*2+2][j])
		}
	}
	build(0, 0, n-1)

	return &Tree{sum, lazy, sz}
}

func rotate(a []int, v int) {
	v %= len(a)
	slices.Reverse(a[:v])
	slices.Reverse(a[v:])
	slices.Reverse(a)
}

func (tr *Tree) apply(i int, v int) {
	rotate(tr.sum[i][:], v)
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
	for j := range 48 {
		tr.sum[i][j] = add(tr.sum[i*2+1][j], tr.sum[i*2+2][j])
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
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		tr.pull(i)
	}
	f(0, 0, tr.sz-1, L, R)
}

func (tr *Tree) Query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return tr.sum[i][0]
		}
		tr.push(i)
		mid := (l + r) >> 1
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return f(i*2+2, mid+1, r, L, R)
		}
		ans := f(i*2+1, l, mid, L, mid)
		ans = add(ans, f(i*2+2, mid+1, r, mid+1, R))
		return ans
	}
	return f(0, 0, tr.sz-1, L, R)
}
