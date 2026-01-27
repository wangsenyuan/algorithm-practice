package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, _, _, res := drive(reader)
		fmt.Fprintln(writer, res[0], res[1])
	}
}

func drive(reader *bufio.Reader) (m int, t int, c []int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &m, &t)
	c = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	res = solve(m, t, c)
	return
}

type pair struct {
	first  int
	second int
}

func solve(m int, t int, c []int) []int {
	c1 := slices.Clone(c)
	slices.Sort(c1)
	c1 = slices.Compact(c1)

	n := len(c)

	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{c[i], i}
	}

	slices.SortFunc(arr, func(x pair, y pair) int {
		return cmp.Or(x.first-y.first, x.second-y.second)
	})

	sum := NewTree(n)

	check := func(w int) bool {
		if w == 0 {
			return true
		}
		// 可否完成前w个任务
		w1 := w / m * m
		if w1 < w {
			s1 := sum.Get(w1)
			// s1 < t
			s2 := sum.Get(w) - s1
			return s1*2+s2 <= t
		}
		// w1 == w
		s1 := sum.Get(w - m)
		s2 := sum.Get(w) - s1
		return s1*2+s2 <= t
	}

	find := func(r int) int {
		// when d = d0
		// 找到最大的w, 花费的时间 <= t
		var l int
		r++
		for l < r {
			mid := (l + r) >> 1
			if check(mid) {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l - 1
	}

	best := []int{0, 1}

	for i, j := 0, 0; i < len(c1); i++ {
		for j < n && arr[j].first == c1[i] {
			sum.Update(arr[j].second, c1[i])
			j++
		}
		tmp := find(j)
		if tmp > best[0] {
			best[0] = tmp
			best[1] = c1[i]
		}
	}

	return best
}

type Tree struct {
	sum []int
	cnt []int
	sz  int
}

func NewTree(n int) *Tree {
	return &Tree{
		sum: make([]int, 4*n),
		cnt: make([]int, 4*n),
		sz:  n,
	}
}

func (tr *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.sum[i] += v
			tr.cnt[i]++
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.sum[i] = tr.sum[2*i+1] + tr.sum[2*i+2]
		tr.cnt[i] = tr.cnt[2*i+1] + tr.cnt[2*i+2]
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) Get(k int) int {
	if k == 0 {
		return 0
	}
	if k >= tr.cnt[0] {
		return tr.sum[0]
	}
	var loop func(i int, l int, r int, k int) int
	loop = func(i int, l int, r int, k int) int {
		if l == r {
			return tr.sum[i]
		}
		mid := (l + r) >> 1
		if k <= tr.cnt[2*i+1] {
			return loop(2*i+1, l, mid, k)
		}
		return tr.sum[2*i+1] + loop(2*i+2, mid+1, r, k-tr.cnt[2*i+1])
	}
	return loop(0, 0, tr.sz-1, k)
}
