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
	var s string
	fmt.Fscan(reader, &s)
	queries := make([][]int, q)
	for i := range q {
		var t, x, y int
		fmt.Fscan(reader, &t, &x, &y)
		queries[i] = []int{t, x, y}
	}
	return solve(a, s, queries)
}

func solve(a []int, s string, queries [][]int) []int {
	n := len(a)
	t := build(n)
	for i := range n {
		t.update(i, func(nd *node) {
			switch s[i] {
			case 'R':
				nd.r = a[i]
			case 'G':
				nd.g = a[i]
			default:
				nd.b = a[i]
			}
		})
	}

	var res []int

	for _, cur := range queries {
		if cur[0] == 1 {
			i, x := cur[1]-1, cur[2]
			t.update(i, func(nd *node) {
				switch s[i] {
				case 'R':
					nd.r = x
				case 'G':
					nd.g = x
				default:
					nd.b = x
				}
			})
		} else {
			l, r := cur[1]-1, cur[2]-1
			res = append(res, t.query(l, r+1))
		}
	}

	return res
}

const mod = 998244353

func add(nums ...int) int {
	var res int
	for _, num := range nums {
		res += num
		if res >= mod {
			res -= mod
		}
	}
	return res
}

func mul(nums ...int) int {
	res := 1
	for _, num := range nums {
		res = res * num % mod
	}
	return res
}

type node struct {
	val int
	rg  int
	b   int
	r   int
	gb  int
	g   int
}

func merge(first node, second node) node {
	var res node
	res.val = add(first.val, second.val, mul(first.rg, second.b), mul(first.r, second.gb))

	res.rg = add(first.rg, second.rg, mul(first.r, second.g))
	res.gb = add(first.gb, second.gb, mul(first.g, second.b))
	res.b = add(first.b, second.b)
	res.r = add(first.r, second.r)
	res.g = add(first.g, second.g)

	return res
}

type tree []node

func build(n int) tree {
	return make(tree, n*4)
}

func (t tree) update(pos int, callback func(*node)) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			callback(&t[i])
			return
		}
		mid := (l + r) / 2
		if pos < mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid, r)
		}
		t[i] = merge(t[i*2+1], t[i*2+2])
	}
	n := len(t) / 4
	f(0, 0, n)
}

func (t tree) query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) node

	f = func(i int, l int, r int, L int, R int) node {
		if l == L && r == R {
			return t[i]
		}
		mid := (l + r) / 2
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid <= L {
			return f(i*2+2, mid, r, L, R)
		}
		x := f(i*2+1, l, mid, L, mid)
		y := f(i*2+2, mid, r, mid, R)
		return merge(x, y)
	}
	n := len(t) / 4
	return f(0, 0, n, L, R).val
}
