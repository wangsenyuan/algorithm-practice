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
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	orders := make([][]int, m)
	for i := 0; i < m; i++ {
		var l, r, k int
		fmt.Fscan(reader, &l, &r, &k)
		orders[i] = []int{l, r, k}
	}
	return solve(n, orders)
}

func solve(n int, orders [][]int) []int {
	var dates []int
	for _, cur := range orders {
		dates = append(dates, cur[0]-1, cur[1])
	}
	sort.Ints(dates)
	dates = slices.Compact(dates)

	m := len(dates)
	set := NewSet(m)

	active := make([]int, m)
	for i := range m {
		active[i] = inf
	}

	var sum int

	add := func(r int, l int) {
		sum += dates[r] - l + 1
		set.Update(r)
		active[r] = l
	}

	remove := func(r int) {
		sum -= dates[r] - active[r] + 1
		set.Clear(r)
		active[r] = inf
	}

	deactive := func(i int, l int) {
		// i左边的部分（包括i）已经处理好了
		ni := set.UpperBound(i)
		if ni < m && active[ni] <= dates[i] {
			if active[ni] <= l {
				// 完全被覆盖了
				return
			}
			// 部分重叠
			remove(ni)
			add(ni, dates[i]+1)
		}
		add(i, l)
	}

	activate := func(i int, l int) {
		ni := set.UpperBound(i)
		if ni < m && active[ni] <= dates[i] {
			nl := active[ni]
			remove(ni)
			add(ni, dates[i]+1)
			if nl < l {
				// 中间挖个洞出来
				add(sort.SearchInts(dates, l-1), nl)
			}
		}
	}

	ans := make([]int, len(orders))

	for i, cur := range orders {
		l, r, k := cur[0], cur[1], cur[2]
		u := sort.SearchInts(dates, r)
		j := u
		for {
			j = set.LowerBound(j)
			if j < 0 || dates[j] < l {
				break
			}
			// l <= dates[j]
			ul := active[j]
			remove(j)
			if ul < l {
				// 有重叠区间
				add(sort.SearchInts(dates, l-1), ul)
				break
			}
		}
		if k == 1 {
			deactive(u, l)
		} else {
			activate(u, l)
		}

		ans[i] = n - sum
	}

	return ans
}

type Set struct {
	t1 *SegTree
	t2 *SegTree
	sz int
}

const inf = 1 << 30

func NewSet(n int) *Set {
	t1 := NewSegTree(n, inf, func(a, b int) int {
		return min(a, b)
	})
	t2 := NewSegTree(n, -inf, func(a, b int) int {
		return max(a, b)
	})
	return &Set{t1, t2, n}
}

func (set *Set) Update(p int) {
	set.t1.Update(p, p)
	set.t2.Update(p, p)
}

func (set *Set) Clear(p int) {
	set.t1.Update(p, inf)
	set.t2.Update(p, -inf)
}

func (set *Set) LowerBound(p int) int {
	// 在区间[0...p]之间 最大的数
	return set.t2.Get(0, p+1)
}

func (set *Set) UpperBound(p int) int {
	return set.t1.Get(p, set.sz)
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
