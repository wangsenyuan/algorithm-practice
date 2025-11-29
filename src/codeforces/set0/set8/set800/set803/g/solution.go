package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	queries := make([][]int, m)
	for i := range m {
		var t, l, r, x int
		fmt.Fscan(reader, &t, &l, &r)
		if t == 1 {
			fmt.Fscan(reader, &x)
		}
		queries[i] = []int{t, l, r, x}
	}
	return solve(k, b, queries)
}

func solve(k int, b []int, queries [][]int) []int {
	n := len(b)

	var arr []int
	arr = append(arr, 0)

	for _, cur := range queries {
		l, r := cur[1], cur[2]
		l--
		r--
		arr = append(arr, l, r, r+1)
		if l > 0 {
			arr = append(arr, l-1)
		}
	}

	arr = append(arr, n*k)
	slices.Sort(arr)
	arr = slices.Compact(arr)
	// n2 := len(arr)
	tr := NewTree(b, arr)

	var ans []int

	for _, cur := range queries {
		l, r := cur[1]-1, cur[2]-1
		i := sort.SearchInts(arr, l)
		j := sort.SearchInts(arr, r)
		if cur[0] == 1 {
			x := cur[3]
			tr.Update(i, j, x)
		} else {
			// query
			x := tr.Get(i, j)
			ans = append(ans, x)
		}
	}

	return ans
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

const inf = 1 << 60

func NewTree(b []int, pos []int) *Tree {
	st := NewSegTree(len(b))
	for i, v := range b {
		st.Update(i, v)
	}

	n := len(pos) - 1
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	var f func(i int, l int, r int)

	f = func(i int, l int, r int) {
		if l == r {
			// 这个不大对
			u := pos[l]
			v := pos[l+1]
			if v-u >= len(b) {
				val[i] = st.Get(0, len(b))
			} else {
				// 只是部分区间
				if u%len(b) < v%len(b) {
					val[i] = st.Get(u%len(b), v%len(b))
				} else {
					val[i] = min(st.Get(u%len(b), len(b)), st.Get(0, v%len(b)))
				}
			}

			return
		}
		mid := (l + r) >> 1
		f(i*2+1, l, mid)
		f(i*2+2, mid+1, r)
		val[i] = min(val[i*2+1], val[i*2+2])
	}

	f(0, 0, n-1)
	return &Tree{val, lazy, n}
}

func (tr *Tree) apply(i int, v int) {
	// 把整个区间变成v
	tr.val[i] = v
	tr.lazy[i] = v
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.apply(i*2+1, tr.lazy[i])
		tr.apply(i*2+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.apply(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		tr.val[i] = min(tr.val[i*2+1], tr.val[i*2+2])
	}

	f(0, 0, tr.sz-1, L, R)
}

func (tr *Tree) Get(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return tr.val[i]
		}
		tr.push(i)
		mid := (l + r) >> 1
		res := inf
		if L <= mid {
			res = min(res, f(i*2+1, l, mid, L, min(mid, R)))
		}
		if mid < R {
			res = min(res, f(i*2+2, mid+1, r, max(mid+1, L), R))
		}
		return res
	}
	return f(0, 0, tr.sz-1, L, R)
}

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = inf
	}
	return SegTree(arr)
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = v

	for p > 1 {
		tr[p>>1] = min(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	var res int = inf
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
