package p3933

import (
	"cmp"
	"math/bits"
	"slices"
)

type data struct {
	r int
	c int
	v int
}

func countLocalMaximums1(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])
	arr := make([]data, n*m)
	for i := range n {
		for j := range m {
			arr[i*m+j] = data{i, j, matrix[i][j]}
		}
	}
	slices.SortFunc(arr, func(a data, b data) int {
		return cmp.Or(b.v-a.v, a.r-b.r, a.c-b.c)
	})

	rows := make([]BIT, n)
	for i := range n {
		rows[i] = make(BIT, m+2)
	}

	check := func(u int, v int) bool {
		x := matrix[u][v]
		if x == 0 {
			return false
		}

		lo := u - x
		hi := u + x

		for i := max(0, lo); i <= min(n-1, hi); i++ {
			l, r := v-x, v+x
			if i == lo || i == hi {
				l++
				r--
			}
			if l > r {
				continue
			}
			l = max(0, l)
			r = min(m-1, r)
			cnt := rows[i].getRange(l, r)
			if cnt > 0 {
				return false
			}
		}
		return true
	}

	var res int
	for i := 0; i < n*m; {
		j := i
		for i < n*m && arr[i].v == arr[j].v {
			r, c := arr[i].r, arr[i].c
			if check(r, c) {
				res++
			}
			i++
		}

		for j < i {
			r, c := arr[j].r, arr[j].c
			rows[r].update(c, 1)
			j++
		}
	}

	return res
}

type BIT []int

func (bit BIT) update(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) get(p int) int {
	p++
	var res int
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}

func (bit BIT) getRange(l int, r int) int {
	return bit.get(r) - bit.get(l-1)
}

type sparseTable struct {
	st [][]int
	op func(int, int) int
}

func newSparseTable(a []int, op func(int, int) int) *sparseTable {
	n := len(a)
	lg := bits.Len(uint(n))
	st := make([][]int, lg)
	for i := range lg {
		st[i] = make([]int, n)
	}
	copy(st[0], a)
	for i := 1; i < lg; i++ {
		for j := 0; j+(1<<i) <= n; j++ {
			st[i][j] = op(st[i-1][j], st[i-1][j+(1<<(i-1))])
		}
	}
	return &sparseTable{st, op}
}

func (st *sparseTable) query(l int, r int) int {
	k := bits.Len(uint(r-l+1)) - 1
	return st.op(st.st[k][l], st.st[k][r-(1<<k)+1])
}

type segTree []*sparseTable

func newSegTree(mat [][]int) segTree {
	n := len(mat)
	m := len(mat[0])

	op := func(a int, b int) int {
		return max(a, b)
	}

	nodes := make(segTree, 4*n)

	var f func(int, int, int)

	f = func(i int, l int, r int) {
		if l == r {
			nodes[i] = newSparseTable(mat[l], op)
		} else {
			mid := (l + r) / 2
			f(i*2+1, l, mid)
			f(i*2+2, mid+1, r)
			arr := make([]int, m)
			for j := range m {
				arr[j] = op(nodes[i*2+1].st[0][j], nodes[i*2+2].st[0][j])
			}
			nodes[i] = newSparseTable(arr, op)
		}
	}

	f(0, 0, n-1)

	return nodes
}

func (t segTree) Query(r1 int, c1 int, r2 int, c2 int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return t[i].query(c1, c2)
		}
		mid := (l + r) / 2
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return f(i*2+2, mid+1, r, L, R)
		}
		a := f(i*2+1, l, mid, L, mid)
		b := f(i*2+2, mid+1, r, mid+1, R)
		return max(a, b)
	}
	return f(0, 0, len(t)/4-1, r1, r2)
}

func countLocalMaximums(matrix [][]int) int {
	t1 := newSegTree(matrix)
	t2 := newSegTree(matrix)

	var res int

	n := len(matrix)
	m := len(matrix[0])

	for r, cur := range matrix {
		for c, x := range cur {
			if x > 0 {
				u1 := t1.Query(max(0, r-x), max(0, c-x+1), min(n-1, r+x), min(m-1, c+x-1))
				u2 := t2.Query(max(0, r-x+1), max(0, c-x), min(n-1, r+x-1), min(m-1, c+x))
				if max(u1, u2) <= x {
					res++
				}
			}
		}
	}
	return res
}
