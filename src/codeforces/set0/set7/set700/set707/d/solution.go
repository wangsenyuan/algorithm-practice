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
	res := drive(reader)
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)
	queries := make([][]int, q)
	for i := range q {
		var t, r, c int
		fmt.Fscan(reader, &t, &r)
		if t <= 2 {
			fmt.Fscan(reader, &c)
		}
		queries[i] = []int{t, r, c}
	}
	return solve(n, m, queries)
}

func solve(n int, m int, queries [][]int) []int {
	w := len(queries)

	trs := make([]*node, w+1)
	trs[0] = NewTree(n * m)

	res := make([]int, w)

	for i, cur := range queries {
		switch cur[0] {
		case 1:
			r, c := cur[1]-1, cur[2]-1
			trs[i+1] = trs[i].Set(r*m+c, 1)
		case 2:
			r, c := cur[1]-1, cur[2]-1
			trs[i+1] = trs[i].Set(r*m+c, 0)
		case 3:
			r := cur[1] - 1
			trs[i+1] = trs[i].Invert(r*m, (r+1)*m-1)
		case 4:
			k := cur[1]
			trs[i+1] = trs[k]
		}

		res[i] = trs[i+1].val
	}

	return res
}

// lazy + 持久树, 还有点麻烦呐
type node struct {
	lf, rg   *node
	l, r     int
	val      int
	inverted int
}

func NewTree(n int) *node {
	var build func(l int, r int) *node
	build = func(l int, r int) *node {
		if l == r {
			return &node{nil, nil, l, r, 0, 0}
		}
		mid := (l + r) >> 1
		lf := build(l, mid)
		rg := build(mid+1, r)
		return &node{lf, rg, l, r, 0, 0}
	}
	return build(0, n-1)
}

func (n *node) copy() *node {
	return &node{n.lf, n.rg, n.l, n.r, n.val, n.inverted}
}

func (n *node) invert() {
	n.val = n.r - n.l + 1 - n.val
	n.inverted ^= 1
}

func (n *node) push() (*node, *node) {
	if n.inverted == 0 {
		return n.lf, n.rg
	}
	n.inverted = 0

	lf := n.lf.copy()
	lf.invert()
	rg := n.rg.copy()
	rg.invert()
	return lf, rg
}

func (n node) Set(p int, v int) *node {
	if n.l == n.r {
		n.val = v
	} else {
		lf, rg := n.push()
		mid := (n.l + n.r) >> 1

		if p <= mid {
			lf = lf.Set(p, v)
		} else {
			rg = rg.Set(p, v)
		}
		n.val = lf.val + rg.val
		n.lf = lf
		n.rg = rg
	}
	return &n
}

func (n node) Invert(l int, r int) *node {
	if n.l == l && n.r == r {
		n.invert()
	} else {
		mid := (n.l + n.r) >> 1
		lf, rg := n.push()
		if l <= mid {
			lf = lf.Invert(l, min(mid, r))
		}
		if mid < r {
			rg = rg.Invert(max(mid+1, l), r)
		}
		n.val = lf.val + rg.val
		n.lf = lf
		n.rg = rg
	}

	return &n
}
