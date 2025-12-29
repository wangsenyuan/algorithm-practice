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
	_, best, ans := drive(reader)
	fmt.Println(best)
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (coupons [][]int, best int, ans []int) {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	coupons = make([][]int, n)
	for i := range n {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		coupons[i] = []int{l, r}
	}
	best, ans = solve(k, coupons)
	return
}

func solve(k int, coupons [][]int) (int, []int) {
	// n := len(coupons)
	var arr []int
	for _, cur := range coupons {
		arr = append(arr, cur[0], cur[1])
	}

	slices.Sort(arr)
	arr = slices.Compact(arr)

	m := len(arr)

	open := make([][]int, m)
	close := make([][]int, m)

	for _, cur := range coupons {
		l := sort.SearchInts(arr, cur[0])
		r := sort.SearchInts(arr, cur[1])
		open[l] = append(open[l], r)
		close[r] = append(close[r], l)
	}

	var best []int

	tr := NewTree(m)

	var j int
	for i := range m {
		for _, r := range open[i] {
			tr.Update(i, r, 1)
		}
		if tr.Get(i, i) >= k {
			for j < i && tr.Get(j, i) < k {
				j++
			}
			if len(best) == 0 || arr[i]-arr[j]+1 > best[1]-best[0] {
				best = []int{arr[j], arr[i]}
			}
		}
		for _, l := range close[i] {
			tr.Update(l, i, -1)
		}
	}

	var ans []int

	if len(best) == 0 {
		for i := range k {
			ans = append(ans, i+1)
		}
		return 0, ans
	}

	l, r := best[0], best[1]
	for i, cur := range coupons {
		if cur[0] <= l && cur[1] >= r {
			ans = append(ans, i+1)
		}
		if len(ans) == k {
			break
		}
	}

	return best[1] - best[0] + 1, ans
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	sz := n
	return &Tree{val, lazy, sz}
}

func (tr *Tree) apply(i int, v int) {
	tr.val[i] += v
	tr.lazy[i] += v
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.apply(2*i+1, tr.lazy[i])
		tr.apply(2*i+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *Tree) pull(i int) {
	tr.val[i] = min(tr.val[2*i+1], tr.val[2*i+2])
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
			f(2*i+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(2*i+2, mid+1, r, max(mid+1, L), R)
		}
		tr.pull(i)
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
		if R <= mid {
			return f(2*i+1, l, mid, L, R)
		}
		if mid < L {
			return f(2*i+2, mid+1, r, L, R)
		}
		x := f(2*i+1, l, mid, L, mid)
		y := f(2*i+2, mid+1, r, mid+1, R)
		return min(x, y)
	}
	return f(0, 0, tr.sz-1, L, R)
}
