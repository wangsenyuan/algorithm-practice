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
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(p)
}

func solve(p []int) int {
	n := len(p)
	start := 0
	for i := range n {
		if p[i] == 1 {
			start = i
			break
		}
	}
	q := make([]int, n)
	copy(q, p[start:])
	copy(q[n-start:], p[:start])

	pos := make([]int, n+1)
	for i, v := range q {
		pos[v] = i
	}

	st := newLazySegTree(n)
	seen := make([]bool, n+2)
	for i, v := range q {
		seen[v] = true
		if !seen[v-1] {
			st.add(i, n, 1)
		}
		if seen[v+1] {
			st.add(i, n, -1)
		}
	}

	ans := 0
	if st.max() <= 2 {
		ans++
	}

	for i := 1; i < n; i++ {
		v := q[i-1]
		if v == 1 {
			st.add(1, n, -1)
		} else {
			x := pos[v-1]
			addCircular(st, i-1, x, n, -1)
		}
		if v < n {
			x := pos[v+1]
			if i-1 < x {
				st.add(x, n, 1)
				st.add(0, i-1, 1)
			} else {
				st.add(x, i-1, 1)
			}
		}
		st.add(i, i+1, 1)
		if st.max() <= 2 {
			ans++
		}
	}

	return ans
}

func addCircular(st *lazySegTree, l int, r int, n int, v int) {
	if l < r {
		st.add(l, r, v)
	} else {
		st.add(l, n, v)
		st.add(0, r, v)
	}
}

type lazySegTree struct {
	val []int
	mx  []int
	sz  int
}

func newLazySegTree(n int) *lazySegTree {
	sz := 1
	for sz < n {
		sz <<= 1
	}
	return &lazySegTree{make([]int, 2*sz), make([]int, 2*sz), sz}
}

func (st *lazySegTree) add(l int, r int, v int) {
	st.addAt(l, r, v, 0, st.sz, 1)
}

func (st *lazySegTree) addAt(l int, r int, v int, tl int, tr int, k int) {
	if r <= l || r <= tl || tr <= l {
		return
	}
	if l <= tl && tr <= r {
		st.val[k] += v
		st.mx[k] += v
		return
	}
	tm := (tl + tr) / 2
	st.addAt(l, r, v, tl, tm, k*2)
	st.addAt(l, r, v, tm, tr, k*2+1)
	st.mx[k] = st.val[k] + max(st.mx[k*2], st.mx[k*2+1])
}

func (st *lazySegTree) max() int {
	return st.mx[1]
}
