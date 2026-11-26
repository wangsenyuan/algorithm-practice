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

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	ops := make([][]int, m)
	for i := range m {
		var tpy int
		fmt.Fscan(reader, &tpy)
		switch tpy {
		case 1:
			var l, r int
			fmt.Fscan(reader, &l, &r)
			ops[i] = []int{tpy, l, r}
		case 2:
			var l, r, x int
			fmt.Fscan(reader, &l, &r, &x)
			ops[i] = []int{tpy, l, r, x}
		default:
			var k, x int
			fmt.Fscan(reader, &k, &x)
			ops[i] = []int{tpy, k, x}
		}
	}
	return solve(a, ops)
}

func solve(a []int, ops [][]int) (ans []int64) {
	tr := build(a)

	for _, cur := range ops {
		switch cur[0] {
		case 1:
			l, r := cur[1], cur[2]
			ans = append(ans, int64(tr.query(l-1, r)))
		case 2:
			l, r, x := cur[1], cur[2], cur[3]
			tr.update(l-1, r, x)
		default:
			// cur[0] == 3
			k, x := cur[1], cur[2]
			tr.set(k-1, x)
		}
	}

	return
}

type node struct {
	sum int
	mx  int
}

func merge(a node, b node) node {
	return node{
		sum: a.sum + b.sum,
		mx:  max(a.mx, b.mx),
	}
}

type tree []node

func build(a []int) tree {
	n := len(a)
	tr := make([]node, 4*n)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			tr[i] = node{a[l], a[l]}
			return
		}
		mid := (l + r) >> 1
		f(i*2+1, l, mid)
		f(i*2+2, mid, r)
		tr[i] = merge(tr[i*2+1], tr[i*2+2])
	}
	f(0, 0, n)
	return tr
}

func (tr tree) update(L int, R int, x int) {
	var g func(i int, l int, r int)
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			g(i, l, r)
			return
		}
		if tr[i].mx < x {
			return
		}
		mid := (l + r) >> 1
		if L < mid {
			f(i*2+1, l, mid, L, min(R, mid))
		}
		if mid < R {
			f(i*2+2, mid, r, max(L, mid), R)
		}
		tr[i] = merge(tr[i*2+1], tr[i*2+2])
	}

	g = func(i int, l int, r int) {
		if tr[i].mx < x {
			return
		}
		if l+1 == r {
			tr[i].sum %= x
			tr[i].mx = tr[i].sum
			return
		}
		mid := (l + r) >> 1
		g(i*2+1, l, mid)
		g(i*2+2, mid, r)
		tr[i] = merge(tr[i*2+1], tr[i*2+2])
	}

	f(0, 0, len(tr)/4, L, R)
}

func (tr tree) set(pos int, x int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			tr[i].sum = x
			tr[i].mx = x
			return
		}
		mid := (l + r) >> 1
		if pos < mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid, r)
		}
		tr[i] = merge(tr[i*2+1], tr[i*2+2])
	}
	f(0, 0, len(tr)/4)
}

func (tr tree) query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return tr[i].sum
		}
		mid := (l + r) >> 1
		var res int
		if L < mid {
			res += f(i*2+1, l, mid, L, min(R, mid))
		}
		if mid < R {
			res += f(i*2+2, mid, r, max(L, mid), R)
		}
		return res
	}
	return f(0, 0, len(tr)/4, L, R)
}
