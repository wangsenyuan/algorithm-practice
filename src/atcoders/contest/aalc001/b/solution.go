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
		var t, l, r int
		fmt.Fscan(reader, &t, &l, &r)
		if t == 1 {
			var x int
			fmt.Fscan(reader, &x)
			queries[i] = []int{t, l, r, x}
			continue
		}
		queries[i] = []int{t, l, r}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	t := build(a)
	var ans []int
	for _, cur := range queries {
		switch cur[0] {
		case 1:
			t.add(cur[1]-1, cur[2], cur[3])
		case 2:
			ans = append(ans, t.query(cur[1]-1, cur[2]))
		}
	}
	return ans
}

const inf = 1 << 60

type tree struct {
	val []int
	tag []int
}

func build(a []int) *tree {
	n := len(a)
	t := &tree{
		val: make([]int, 4*n),
		tag: make([]int, 4*n),
	}

	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			t.val[i] = a[l]
			return
		}
		mid := (l + r) >> 1
		f(i*2+1, l, mid)
		f(i*2+2, mid, r)
		t.val[i] = min(t.val[i*2+1], t.val[i*2+2])
	}
	f(0, 0, n)
	return t
}

func (t *tree) apply(i int, x int) {
	t.val[i] += x
	t.tag[i] += x
}

func (t *tree) push(i int) {
	if t.tag[i] != 0 {
		t.apply(i*2+1, t.tag[i])
		t.apply(i*2+2, t.tag[i])
		t.tag[i] = 0
	}
}

func (t *tree) add(L int, R int, x int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.apply(i, x)
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		if L < mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid, r, max(mid, L), R)
		}
		t.val[i] = min(t.val[i*2+1], t.val[i*2+2])
	}
	f(0, 0, len(t.val)/4, L, R)
}

func (t *tree) query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return t.val[i]
		}
		t.push(i)
		mid := (l + r) >> 1
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid <= L {
			return f(i*2+2, mid, r, L, R)
		}
		return min(f(i*2+1, l, mid, L, mid), f(i*2+2, mid, r, mid, R))
	}
	return f(0, 0, len(t.val)/4, L, R)
}
