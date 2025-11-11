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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var h, n int
	fmt.Fscan(reader, &h, &n)
	queries := make([][]int, n)
	for i := range n {
		queries[i] = make([]int, 4)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2], &queries[i][3])
	}
	return solve(h, queries)
}

type data struct {
	l    int
	r    int
	flag int
}

func solve(h int, queries [][]int) string {
	var arr []data

	for _, cur := range queries {
		i, l, r := cur[0], cur[1], cur[2]
		for i < h {
			l *= 2
			r = 2*r + 1
			i++
		}

		arr = append(arr, data{l, r, cur[3]})
	}

	var pos []int
	pos = append(pos, (1 << (h - 1)))
	pos = append(pos, (1<<h)-1)

	for _, cur := range arr {
		pos = append(pos, cur.l, cur.r)
		if cur.l > 1<<(h-1) {
			pos = append(pos, cur.l-1)
		}
		if cur.r < (1<<h)-1 {
			pos = append(pos, cur.r+1)
		}
	}

	slices.Sort(pos)
	pos = slices.Compact(pos)

	n := len(pos)
	tr := NewTree(n)
	tr.Update(0, n-1, 1)

	for _, cur := range arr {
		l := sort.SearchInts(pos, cur.l)
		r := sort.SearchInts(pos, cur.r)
		if cur.flag == 0 {
			tr.Update(l, r, 0)
		} else {
			if l > 0 {
				tr.Update(0, l-1, 0)
			}
			if r < n-1 {
				tr.Update(r+1, n-1, 0)
			}
		}
	}
	cnt := tr.Query(0, n-1)

	if cnt == 0 {
		return "Game cheated!"
	}
	if cnt > 1 {
		return "Data not sufficient!"
	}
	i := tr.FindThePosition()
	// 必须只有一个点
	if i > 0 && pos[i] != pos[i-1]+1 || i < n-1 && pos[i]+1 != pos[i+1] {
		return "Data not sufficient!"
	}
	return fmt.Sprintf("%d", pos[i])
}

type Tree struct {
	arr  []int
	lazy []int
	sz   int
}

func NewTree(sz int) *Tree {
	arr := make([]int, 4*sz)
	lazy := make([]int, 4*sz)
	for i := range lazy {
		lazy[i] = -1
	}
	return &Tree{arr, lazy, sz}
}

func (tr *Tree) apply(i int, l int, r int, v int) {
	tr.arr[i] = v * (r - l + 1)
	tr.lazy[i] = v
}

func (tr *Tree) push(i int, l int, r int) {
	mid := (l + r) >> 1
	if l < r && tr.lazy[i] != -1 {
		tr.apply(i*2+1, l, mid, tr.lazy[i])
		tr.apply(i*2+2, mid+1, r, tr.lazy[i])
		tr.lazy[i] = -1
	}
}

func (tr *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.apply(i, l, r, v)
			return
		}
		tr.push(i, l, r)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		tr.arr[i] = tr.arr[i*2+1] + tr.arr[i*2+2]
	}
	f(0, 0, tr.sz-1, L, R)
}

func (tr *Tree) Query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return tr.arr[i]
		}
		tr.push(i, l, r)
		mid := (l + r) >> 1
		var res int
		if L <= mid {
			res += f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			res += f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		return res
	}
	return f(0, 0, tr.sz-1, L, R)
}

func (tr *Tree) FindThePosition() int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		tr.push(i, l, r)
		mid := (l + r) >> 1
		if tr.arr[i*2+1] > 0 {
			return f(i*2+1, l, mid)
		}
		return f(i*2+2, mid+1, r)
	}

	return f(0, 0, tr.sz-1)
}
