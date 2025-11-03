package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
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

	var res int
	for i := range n {
		if a[i] == 0 {
			res++
		}
	}

	sort.Ints(a)

	tr := NewTree(a)

	buf := make([]int, n)
	for i := range n {
		// 如果以i为开始位置
		tr.Update(i, -inf)
		for j := range n {
			if i == j || a[i] == 0 && a[j] == 0 {
				continue
			}
			tr.Update(j, -inf)
			u, v := i, j
			var it int
			for {
				next := a[u] + a[v]
				if tr[0] < next {
					break
				}
				w := tr.UpperBound(next)
				if a[w] != next {
					break
				}
				tr.Update(w, -inf)
				buf[it] = w
				it++
				u, v = v, w
			}

			res = max(res, it+2)
			for k := 0; k < it; k++ {
				tr.Update(buf[k], a[buf[k]])
			}
			tr.Update(j, a[j])
		}
		tr.Update(i, a[i])
	}

	return res
}

const inf = 1 << 60

type Tree []int

func NewTree(a []int) Tree {
	n := len(a)
	tr := make(Tree, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			tr[i] = a[l]
			return
		}
		mid := (l + r) >> 1
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		tr[i] = max(tr[i*2+1], tr[i*2+2])
	}

	build(0, 0, n-1)
	return tr
}

func (tr Tree) Update(p int, v int) {

	var f func(i int, l int, r int)

	f = func(i int, l int, r int) {
		if l == r {
			tr[i] = v
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr[i] = max(tr[i*2+1], tr[i*2+2])
	}

	n := len(tr) / 4

	f(0, 0, n-1)
}

func (tr Tree) UpperBound(v int) int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		mid := (l + r) >> 1
		if tr[2*i+1] >= v {
			return f(i*2+1, l, mid)
		}
		return f(i*2+2, mid+1, r)
	}
	n := len(tr) / 4
	return f(0, 0, n-1)
}
