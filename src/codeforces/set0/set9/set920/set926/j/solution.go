package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
	"slices"
	"sort"
)

func main() {
	debug.SetGCPercent(-1)
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	buf := make([]byte, 4096)
	var _i, _n int
	rc := func() byte {
		if _i == _n {
			_n, _ = reader.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	rd := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n := rd()
	segments := make([][]int, n)
	for i := 0; i < n; i++ {
		segments[i] = []int{rd(), rd()}
	}

	return solve(segments)
}

func solve(segments [][]int) []int {
	var arr []int
	for _, cur := range segments {
		arr = append(arr, cur[0], cur[1])
	}
	sort.Ints(arr)
	arr = slices.Compact(arr)
	m := len(arr)
	var cnt int

	set := NewSet(m)

	ptr := make([]int, m)
	for i := 0; i < m; i++ {
		ptr[i] = i
	}

	work := func(l int, r int) {
		a := set.LowerBound(l)
		if a >= 0 && ptr[a] >= l {
			if r <= ptr[a] {
				return
			}
			cnt--
			set.Unset(a)
			l = a
		}

		for {
			b := set.UpperBound(l)
			if r < b {
				break
			}
			// b <= r
			set.Unset(b)
			cnt--
			r = max(r, ptr[b])
		}
		cnt++
		ptr[l] = r
		set.Set(l)
	}

	ans := make([]int, len(segments))

	for i, cur := range segments {
		l := sort.SearchInts(arr, cur[0])
		r := sort.SearchInts(arr, cur[1])
		work(l, r)
		ans[i] = cnt
	}

	return ans
}

type Set struct {
	min_pos *SegTree
	max_pos *SegTree
	n       int
}

func NewSet(n int) Set {
	min_pos := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})
	max_pos := NewSegTree(n, -1, func(a, b int) int {
		return max(a, b)
	})
	return Set{min_pos, max_pos, n}
}

func (set *Set) Set(p int) {
	set.min_pos.Update(p, p)
	set.max_pos.Update(p, p)
}

func (set *Set) Unset(p int) {
	set.min_pos.Update(p, set.n)
	set.max_pos.Update(p, -1)
}

func (set Set) LowerBound(p int) int {
	return set.max_pos.Get(0, p)
}

func (set Set) UpperBound(p int) int {
	return set.min_pos.Get(p, set.n)
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
