package main

import (
	"bufio"
	"os"

	. "fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m, q int
	Fscan(reader, &n, &m, &q)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			Fscan(reader, &a[i][j])
		}
	}
	queries := make([][]int, q)
	for i := range q {
		var t, x, y int
		Fscan(reader, &t, &x, &y)
		queries[i] = []int{t, x, y}
	}
	return solve(a, queries)
}

func solve(a [][]int, queries [][]int) []int {
	n := len(a)
	m := len(a[0])
	cols := make([]*Set, m)
	rows := make([]*Set, n)
	for i := range m {
		cols[i] = NewSet(n)
	}
	for i := range n {
		rows[i] = NewSet(m)
	}

	for i := range n {
		for j := range m {
			if a[i][j] == 0 {
				cols[j].Set(i)
				rows[i].Set(j)
			}
		}
	}

	query := func(x int, y int) int {

		var best int
		lo, hi := -1, n

		// 如果x,y是右边（也包括角落)
		for j := y; j >= 0; j-- {
			if a[x][j] == 0 {
				break
			}
			u := cols[j].LowerBound(x)
			v := cols[j].UpperBound(x)
			lo = max(lo, u)
			hi = min(hi, v)
			best = max(best, (y-j+1)*(hi-lo-1))
		}
		lo, hi = -1, n
		for j := y; j < m; j++ {
			if a[x][j] == 0 {
				break
			}
			u := cols[j].LowerBound(x)
			v := cols[j].UpperBound(x)
			lo = max(lo, u)
			hi = min(hi, v)
			best = max(best, (j-y+1)*(hi-lo-1))
		}
		// 还有上下边
		lo, hi = -1, m
		for i := x; i >= 0; i-- {
			if a[i][y] == 0 {
				break
			}
			u := rows[i].LowerBound(y)
			v := rows[i].UpperBound(y)
			lo = max(lo, u)
			hi = min(hi, v)
			best = max(best, (x-i+1)*(hi-lo-1))
		}
		lo, hi = -1, m
		for i := x; i < n; i++ {
			if a[i][y] == 0 {
				break
			}
			u := rows[i].LowerBound(y)
			v := rows[i].UpperBound(y)
			lo = max(lo, u)
			hi = min(hi, v)
			best = max(best, (i-x+1)*(hi-lo-1))
		}

		return best
	}

	var res []int

	for _, cur := range queries {
		x, y := cur[1]-1, cur[2]-1
		if cur[0] == 1 {
			if a[x][y] == 0 {
				cols[y].Clear(x)
				rows[x].Clear(y)
			} else {
				cols[y].Set(x)
				rows[x].Set(y)
			}
			a[x][y] ^= 1
		} else {
			res = append(res, query(x, y))
		}
	}

	return res
}

const inf = 1 << 60

type Set struct {
	n  int
	t1 *SegTree
	t2 *SegTree
}

func NewSet(n int) *Set {
	t1 := NewSegTree(n, inf, func(x int, y int) int {
		return min(x, y)
	})
	t2 := NewSegTree(n, -inf, func(x int, y int) int {
		return max(x, y)
	})
	return &Set{n: n, t1: t1, t2: t2}
}

func (s *Set) Set(p int) {
	s.t1.Update(p, p)
	s.t2.Update(p, p)
}

func (s *Set) Clear(p int) {
	s.t1.Update(p, inf)
	s.t2.Update(p, -inf)
}

func (s *Set) UpperBound(p int) int {
	// res >= p
	return s.t1.Query(p, s.n)
}
func (s *Set) LowerBound(p int) int {
	// res <= p
	return s.t2.Query(0, p+1)
}

type SegTree struct {
	val          []int
	initialValue int
	merge        func(int, int) int
}

func NewSegTree(n int, initialValue int, merge func(int, int) int) *SegTree {
	val := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		val[i] = initialValue
	}
	return &SegTree{val: val, initialValue: initialValue, merge: merge}
}

func (st *SegTree) Update(p int, v int) {
	n := len(st.val) / 2
	p += n
	st.val[p] = v
	for p > 1 {
		st.val[p>>1] = st.merge(st.val[p], st.val[p^1])
		p >>= 1
	}
}

func (st *SegTree) Query(l int, r int) int {
	n := len(st.val) / 2
	l += n
	r += n
	res := st.initialValue
	for l < r {
		if l&1 == 1 {
			res = st.merge(res, st.val[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = st.merge(res, st.val[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
