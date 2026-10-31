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
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var pos, x int
			fmt.Fscan(reader, &pos, &x)
			queries[i] = []int{t, pos, x}
		} else {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			queries[i] = []int{t, l, r}
		}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	s := make(segtree, 2*n)
	for i := range n {
		s.update(i, a[i])
	}
	var ans []int
	for _, cur := range queries {
		switch cur[0] {
		case 1:
			s.update(cur[1]-1, cur[2])
		case 2:
			ans = append(ans, s.query(cur[1]-1, cur[2]))
		}
	}
	return ans
}

const inf = 1 << 60

type segtree []int

func (s segtree) update(i int, x int) {
	n := len(s) / 2
	i += n
	s[i] = x
	for i > 1 {
		s[i>>1] = min(s[i], s[i^1])
		i >>= 1
	}
}

func (s segtree) query(l int, r int) int {
	n := len(s) / 2
	l += n
	r += n
	res := inf
	for l < r {
		if l&1 == 1 {
			res = min(res, s[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, s[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
