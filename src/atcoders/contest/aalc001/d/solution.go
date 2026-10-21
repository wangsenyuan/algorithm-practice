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

	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		var t, x, y int
		fmt.Fscan(reader, &t, &x, &y)
		queries[i] = []int{t, x, y}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	tr := buildTree(n)
	for i, v := range a {
		tr.update(i, v)
	}
	var ans []int
	for _, cur := range queries {
		if cur[0] == 1 {
			i, x := cur[1]-1, cur[2]
			tr.update(i, x)
		} else {
			l, r := cur[1]-1, cur[2]
			ans = append(ans, tr.query(l, r))
		}
	}
	return ans
}

type node struct {
	val int
	cnt [3]int
}

type tree []node

func buildTree(n int) tree {
	return make(tree, n*4)
}

func merge(a node, b node) node {
	var c node
	for i := range 3 {
		c.cnt[i] = a.cnt[i] + b.cnt[i]
	}
	c.val = a.val + b.val
	c.val += a.cnt[2] * (b.cnt[0] + b.cnt[1])
	c.val += a.cnt[1] * b.cnt[0]
	return c
}

func (tr tree) update(pos int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			clear(tr[i].cnt[:])
			tr[i].cnt[v]++
		} else {
			mid := (l + r) >> 1
			if pos < mid {
				f(i*2+1, l, mid)
			} else {
				f(i*2+2, mid, r)
			}
			tr[i] = merge(tr[i*2+1], tr[i*2+2])
		}
	}
	n := len(tr) / 4
	f(0, 0, n)
}

func (tr tree) query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) node
	f = func(i int, l int, r int, L int, R int) node {
		if l == L && r == R {
			return tr[i]
		}
		mid := (l + r) / 2
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid <= L {
			return f(i*2+2, mid, r, L, R)
		}
		a := f(i*2+1, l, mid, L, mid)
		b := f(i*2+2, mid, r, mid, R)
		return merge(a, b)
	}
	n := len(tr) / 4
	res := f(0, 0, n, L, R)
	return res.val
}
