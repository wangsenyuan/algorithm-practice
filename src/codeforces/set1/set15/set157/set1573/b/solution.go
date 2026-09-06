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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a, b []int) int {
	n := len(a)
	// 2 * n
	tr := build(2 * n)
	for i, v := range a {
		tr.update(v, i)
	}
	best := inf

	for i, v := range b {
		// 找到 1..v 中间离0最近的位置j
		j := tr.get(1, v)
		best = min(best, i+j)
	}

	return best
}

const inf = 1 << 60

type segtree []int

func build(n int) segtree {
	tr := make(segtree, 2*n)
	for i := range 2 * n {
		tr[i] = inf
	}
	return tr
}

func (tr segtree) update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = v
	for p > 1 {
		tr[p>>1] = min(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr segtree) get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	res := inf
	for l < r {
		if l&1 == 1 {
			res = min(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
