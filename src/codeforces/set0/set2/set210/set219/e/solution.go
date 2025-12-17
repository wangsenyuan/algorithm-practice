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
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	records := make([][]int, m)
	for i := range m {
		records[i] = make([]int, 2)
		fmt.Fscan(reader, &records[i][0], &records[i][1])
	}
	return solve(n, records)
}

func solve(n int, records [][]int) []int {
	var pid int
	for _, record := range records {
		pid = max(pid, record[1])
	}

	pos := make([]int, pid+1)

	var ans []int

	t1 := NewSegTree(n, -1, func(a, b int) int {
		return max(a, b)
	})

	t2 := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})

	getSafeDist := func(l int, r int) int {
		if l == 0 || r == n {
			return r - l
		}
		return (r-1+l)/2 - (l - 1)
	}

	tr := NewTree(n)
	tr.Update(0, getSafeDist(0, n))

	for _, cur := range records {
		if cur[0] == 1 {
			// park
			id := cur[1]

			tmp := tr.GetParkingLot()

			l := tmp.second
			r := t2.Get(l+1, n)
			if l == 0 {
				pos[id] = 0
				tr.Update(0, 0)
				tr.Update(1, getSafeDist(1, r))
			} else if r == n {
				pos[id] = n - 1
				tr.Update(l, getSafeDist(l, n-1))
			} else {
				pos[id] = (l + r - 1) >> 1
				tr.Update(l, getSafeDist(l, pos[id]))
				tr.Update(pos[id]+1, getSafeDist(pos[id]+1, r))
			}

			t1.Update(pos[id], pos[id])
			t2.Update(pos[id], pos[id])

			ans = append(ans, pos[id]+1)
		} else {
			where := pos[cur[1]]
			l := t1.Get(0, where)
			r := t2.Get(where+1, n)
			tr.Update(l+1, getSafeDist(l+1, r))
			if where+1 < n {
				tr.Update(where+1, 0)
			}

			t1.Update(where, -1)
			t2.Update(where, n)
		}
	}

	return ans
}

type Tree struct {
	val []int
	sz  int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)

	for i := range 4 * n {
		val[i] = 0
	}
	return &Tree{val, n}
}

func (tr *Tree) pull(i int) {
	tr.val[i] = max(tr.val[i*2+1], tr.val[i*2+2])
}

func (tr *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = v
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.pull(i)
	}
	f(0, 0, tr.sz-1)
}

type pair struct {
	first  int
	second int
}

func (tr *Tree) GetParkingLot() pair {
	var f func(i int, l int, r int) pair

	f = func(i int, l int, r int) pair {
		if l == r {
			return pair{tr.val[i], l}
		}
		mid := (l + r) >> 1
		if tr.val[i*2+1] >= tr.val[i*2+2] {
			return f(i*2+1, l, mid)
		}
		return f(i*2+2, mid+1, r)
	}
	return f(0, 0, tr.sz-1)
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
