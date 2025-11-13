package main

import (
	"bufio"
	"cmp"
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

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	rects := make([][]int, n)
	for i := range n {
		rects[i] = make([]int, 2)
		fmt.Fscan(reader, &rects[i][0], &rects[i][1])
	}
	return solve(n, rects)
}

type pair struct {
	first  int
	second int
}

func solve(n int, rects [][]int) int {
	arr := make([]pair, n)
	var nums []int
	for i := range n {
		// 躺下后， 高度变成w, 宽度变成h
		arr[i] = pair{-rects[i][0] + rects[i][1], i}
		nums = append(nums, rects[i]...)
	}

	slices.SortFunc(arr, func(a pair, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	slices.Sort(nums)
	nums = slices.Compact(nums)

	m := len(nums)

	at := make([][]int, m)
	var W int
	for i := range n {
		w := rects[i][0]
		W += w
		j := sort.SearchInts(nums, w)
		at[j] = append(at[j], i)
	}
	suf := make([]int, m+1)

	for i := m - 1; i >= 0; i-- {
		suf[i] = suf[i+1]
		for _, j := range at[i] {
			h := rects[j][1]
			suf[i] = max(suf[i], h)
		}
	}

	best := W * suf[0]

	cnt := make(BIT, m+2)
	sum := make(BIT, m+2)

	tr := NewTree(n)

	todo := make([][]int, m)

	update := func(w int, h int, j int) {
		if w <= h {
			return
		}
		delta := -w + h
		k := sort.Search(n, func(k int) bool {
			return arr[k].first > delta || arr[k].first == delta && arr[k].second >= j
		})
		tr.Update(k, delta)
	}

	for i, H := range nums {
		for _, j := range at[i] {
			w, h := rects[j][0], rects[j][1]
			k := sort.SearchInts(nums, h)
			cnt.Update(k, 1)
			sum.Update(k, -w+h)
			if k < i {
				update(w, h, j)
			} else {
				todo[k] = append(todo[k], j)
			}
		}

		for _, j := range todo[i] {
			w, h := rects[j][0], rects[j][1]
			if w > h {
				update(w, h, j)
			}
		}

		// 需要知道那些必须要躺下的人数
		cnt1 := cnt.QueryRange(i+1, m)
		// 后面的都是不能躺下的
		if cnt1 <= n/2 && suf[i+1] <= H {
			cnt2 := n/2 - cnt1
			tmp := sum.QueryRange(i+1, m)
			if cnt2 > 0 {
				tmp += tr.GetBestK(cnt2)
			}
			best = min(best, (W+tmp)*H)
		}
	}

	return best
}

type BIT []int

func (bit BIT) Update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) Get(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) QueryRange(l int, r int) int {
	return bit.Get(r) - bit.Get(l-1)
}

const inf = 1 << 60

type Tree struct {
	val []int
	cnt []int
	sz  int
}

func NewTree(n int) *Tree {
	return &Tree{
		val: make([]int, 4*n),
		cnt: make([]int, 4*n),
		sz:  n,
	}
}

func (t *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			t.cnt[i] = 1
			t.val[i] = v
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}

		t.cnt[i] = t.cnt[i*2+1] + t.cnt[i*2+2]
		t.val[i] = t.val[i*2+1] + t.val[i*2+2]
	}
	f(0, 0, t.sz-1)
}

func (t *Tree) GetBestK(k int) int {
	if k >= t.cnt[0] {
		return t.val[0]
	}
	var f func(i int, l int, r int, k int) int
	f = func(i int, l int, r int, k int) int {
		if l == r {
			return t.val[i]
		}
		mid := (l + r) >> 1
		if t.cnt[i*2+1] >= k {
			return f(i*2+1, l, mid, k)
		}
		return t.val[i*2+1] + f(i*2+2, mid+1, r, k-t.cnt[i*2+1])
	}
	return f(0, 0, t.sz-1, k)
}
