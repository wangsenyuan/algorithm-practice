package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	segs := make([][]int, n)
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		segs[i] = []int{l, r}
	}
	return solve(n, segs)
}

type event struct {
	pos int
	id  int
	tt  int
}

func solve(n int, segs [][]int) []int {
	arr := make([]event, 2*n)
	nums := make([]int, 2*n)

	for i, cur := range segs {
		l, r := cur[0], cur[1]
		arr[i*2] = event{l, i, 1}
		arr[i*2+1] = event{r, i, -1}
		nums[i*2] = l
		nums[i*2+1] = r
	}

	slices.Sort(nums)
	nums = slices.Compact(nums)
	m := len(nums)
	cnt := NewSegTree(m)

	slices.SortFunc(arr, func(a, b event) int {
		return a.pos - b.pos
	})

	ans := make([]int, n)

	for _, cur := range arr {
		if cur.tt == -1 {
			l := segs[cur.id][0]
			r := cur.pos
			i := sort.SearchInts(nums, l)
			j := sort.SearchInts(nums, r)
			ans[cur.id] = cnt.Query(i, j)
			cnt.Add(i, 1)
		}
	}

	return ans
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (seg *SegTree) Add(p int, v int) {
	p += seg.sz
	seg.arr[p] += v
	for p > 1 {
		seg.arr[p>>1] = seg.arr[p] + seg.arr[p^1]
		p >>= 1
	}
}

func (seg *SegTree) Query(l int, r int) int {
	l += seg.sz
	r += seg.sz
	var res int
	for l < r {
		if l&1 == 1 {
			res += seg.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += seg.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
