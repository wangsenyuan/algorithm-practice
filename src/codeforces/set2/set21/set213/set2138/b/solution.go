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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		for _, ok := range drive(reader) {
			if ok {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}
}

func drive(reader *bufio.Reader) []bool {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][2]int, q)
	for i := range q {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(a, queries)
}

func solve(a []int, queries [][2]int) []bool {
	n := len(a)
	todo := make([][]int, n)
	for i, cur := range queries {
		r := cur[1] - 1
		todo[r] = append(todo[r], i)
	}

	s1 := NewSegTree(n)
	s2 := NewSegTree(n)
	for i := range n {
		s1.Update(i, -inf)
		s2.Update(i, -inf)
	}

	ans := make([]bool, len(queries))

	last := -inf

	for i := range n {
		a[i]--
		w := s1.Get(a[i], n)
		s2.Update(a[i], w)

		last = max(last, s2.Get(a[i]+1, n))

		s1.Update(a[i], i)

		for _, id := range todo[i] {
			l := queries[id][0] - 1
			if last >= l {
				ans[id] = false
			} else {
				ans[id] = true
			}
		}
	}

	return ans
}

const inf = 1 << 60

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	return SegTree(arr)
}

func (t SegTree) Update(p int, v int) {
	n := len(t) / 2
	p += n
	t[p] = v
	for p > 1 {
		t[p>>1] = max(t[p], t[p^1])
		p >>= 1
	}
}

func (t SegTree) Get(l int, r int) int {
	n := len(t) / 2
	l += n
	r += n
	res := -inf
	for l < r {
		if l&1 == 1 {
			res = max(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
