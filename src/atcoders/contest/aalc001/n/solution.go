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
	var n, q int
	fmt.Fscan(reader, &n, &q)
	rects := make([][]int, q)
	for i := range q {
		var a, b, c, d int
		fmt.Fscan(reader, &a, &b, &c, &d)
		rects[i] = []int{a, b, c, d}
	}
	return solve(n, rects)
}

const mod = 998244353

func add(a, b int) int {
	return (a + b) % mod
}

func sub(a, b int) int {
	res := (a - b + mod) % mod
	if res < 0 {
		res += mod
	}
	return res
}

func mul(a, b int) int {
	return (a * b) % mod
}

func solve(n int, rects [][]int) int {
	todo := make([][]int, n+1)
	for i, rect := range rects {
		r1, r2 := rect[0]-1, rect[2]
		todo[r1] = append(todo[r1], i+1)
		todo[r2] = append(todo[r2], -(i + 1))
	}
	tr := build(n)
	var ans int
	for r := range n {
		for _, i := range todo[r] {
			if i > 0 {
				i--
				c1, c2 := rects[i][1]-1, rects[i][3]
				tr.update(c1, c2, 1)
			} else {
				i = -i - 1
				c1, c2 := rects[i][1]-1, rects[i][3]
				tr.update(c1, c2, -1)
			}
		}
		if tr[1].value == 0 {
			ans = add(ans, mul(tr[1].sum%mod, (r+1)))
		}
	}
	return ans
}

type node struct {
	value int
	sum   int
	lazy  int
}

func merge(left, right node) node {
	if left.value < right.value {
		return node{value: left.value, sum: left.sum}
	}
	if left.value > right.value {
		return node{value: right.value, sum: right.sum}
	}
	return node{
		value: left.value,
		sum:   left.sum + right.sum,
	}
}

func (n *node) apply(x int) {
	n.value += x
	n.lazy += x
}

type segtree []node

func build(n int) segtree {
	tr := make(segtree, 4*n)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			tr[i] = node{
				value: 0,
				sum:   r,
			}
			return
		}
		mid := (l + r) / 2
		f(i*2, l, mid)
		f(i*2+1, mid, r)
		tr[i] = merge(tr[i*2], tr[i*2+1])
	}
	f(1, 0, n)
	return tr
}

func (tr segtree) push(i int) {
	if tr[i].lazy != 0 {
		tr[i*2].apply(tr[i].lazy)
		tr[i*2+1].apply(tr[i].lazy)
		tr[i].lazy = 0
	}
}

func (tr segtree) update(L int, R int, x int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr[i].apply(x)
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if L < mid {
			f(i*2, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+1, mid, r, max(mid, L), R)
		}
		tr[i] = merge(tr[i*2], tr[i*2+1])
	}

	f(1, 0, len(tr)/4, L, R)
}
