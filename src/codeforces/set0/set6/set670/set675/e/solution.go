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
	a := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, a)
}

func solve(n int, a []int) int {
	dp := NewSegTree(n + 1)
	dp.Set(n, n)
	var res int

	for i := n - 1; i > 0; i-- {
		r := a[i-1]
		w := dp.Get(i+1, r+1) - r
		// 假设最优的位置在j, 那么w - (r - j) 才可以
		w += n - i
		res += w
		dp.Set(i, w+i)
	}

	return res
}

const inf = 1 << 60

type SegTree []int

func NewSegTree(n int) SegTree {
	res := make(SegTree, 2*n)
	for i := range res {
		res[i] = inf
	}
	return res
}

func (this SegTree) Set(p int, v int) {
	n := len(this) >> 1
	p += n
	this[p] = v
	for p > 1 {
		this[p>>1] = min(this[p], this[p^1])
		p >>= 1
	}
}

func (this SegTree) Get(l int, r int) int {
	n := len(this) >> 1
	l += n
	r += n
	res := inf
	for l < r {
		if l&1 == 1 {
			res = min(res, this[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, this[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
