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

	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		var t, l, r int
		fmt.Fscan(reader, &t, &l, &r)
		queries[i] = []int{t, l, r}
	}
	return solve(queries)
}

func solve(queries [][]int) []int {
	var pos []int
	for _, cur := range queries {
		l, r := cur[1], cur[2]
		pos = append(pos, l, r+1)
	}
	slices.Sort(pos)
	pos = slices.Compact(pos)

	t := build(pos)

	var ans []int

	for _, cur := range queries {
		l, _ := slices.BinarySearch(pos, cur[1])
		r, _ := slices.BinarySearch(pos, cur[2]+1)
		if cur[0] == 1 {
			t.flip(l, r)
		} else {
			ans = append(ans, t.query(l, r))
		}
	}

	return ans
}

type node struct {
	val int
	tag int
}

type tree struct {
	arr   []int
	nodes []node
	sz    int
}

func build(a []int) *tree {
	n := len(a)
	return &tree{arr: a, nodes: make([]node, 4*n), sz: n}
}

func (t *tree) length(l int, r int) int {
	return t.arr[r] - t.arr[l]
}

func (t *tree) apply(i int, l int, r int) {
	t.nodes[i].val = t.length(l, r) - t.nodes[i].val
	t.nodes[i].tag ^= 1
}

func (t *tree) push(i int, l int, r int) {
	if t.nodes[i].tag == 0 {
		return
	}
	mid := (l + r) / 2
	t.apply(i*2+1, l, mid)
	t.apply(i*2+2, mid, r)
	t.nodes[i].tag = 0
}

func (t *tree) pull(i int) {
	t.nodes[i].val = t.nodes[i*2+1].val + t.nodes[i*2+2].val
}

func (t *tree) flip(L int, R int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.apply(i, l, r)
			return
		}
		t.push(i, l, r)
		mid := (l + r) / 2
		if L < mid {
			f(i*2+1, l, mid, L, min(R, mid))
		}
		if mid < R {
			f(i*2+2, mid, r, max(L, mid), R)
		}
		t.pull(i)
	}
	f(0, 0, t.sz, L, R)
}

func (t *tree) query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return t.nodes[i].val
		}
		t.push(i, l, r)
		mid := (l + r) / 2
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid <= L {
			return f(i*2+2, mid, r, L, R)
		}
		return f(i*2+1, l, mid, L, mid) + f(i*2+2, mid, r, mid, R)
	}
	return f(0, 0, t.sz, L, R)
}
