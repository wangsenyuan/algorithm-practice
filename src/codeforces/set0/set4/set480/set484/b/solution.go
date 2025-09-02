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
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	sort.Ints(a)
	n := len(a)
	m := a[n-1]

	tr := NewSegTree(m+1, 0, func(a, b int) int {
		return max(a, b)
	})

	var res int

	for i := n - 1; i >= 0; i-- {
		if i+1 < n && a[i] == a[i+1] || a[i] == 1 || a[i] <= res {
			continue
		}

		x := a[i]

		for y := x; y <= m; y += x {
			v := tr.Find(y+1, min(m+1, y+x))
			res = max(res, v%x)
		}

		tr.Update(x, x)
	}

	return res
}

func NewSegTree(n int, iv int, fn func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = iv
	}
	return &SegTree{arr, n, iv, fn}
}

type SegTree struct {
	arr       []int
	sz        int
	initValue int
	fn        func(int, int) int
}

func (tree *SegTree) Update(pos int, v int) {
	pos += tree.sz
	tree.arr[pos] = v
	for pos > 0 {
		tree.arr[pos>>1] = tree.fn(tree.arr[pos], tree.arr[pos^1])
		pos >>= 1
	}
}

func (tree *SegTree) Find(l, r int) int {
	l += tree.sz
	r += tree.sz

	ans := tree.initValue

	for l < r {
		if l&1 == 1 {
			ans = tree.fn(ans, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = tree.fn(ans, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}
