package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, m)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, m int) int {
	n := len(a)

	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}

	slices.SortFunc(arr, func(x, y pair) int {
		return x.first - y.first
	})

	// far[i] 表示区间[far[i]...i] 是一个连续(下标)的区间
	// 那么在这个区间里面， 只能取ceil(l / 2)个数
	// 但是问题出在，移动左边的时候，它有可能变成一个非连续的区间
	// 比如 [1, 3, 5, 2, 4]
	// 整个区间可以得到ceil(5/2) = 3个数，但是将5移出后，只能取2个数，4移动后，也能取2个数
	//

	best := arr[n-1].first - arr[0].first
	tr := NewTree(n)

	for l, r := 0, 0; r < n; r++ {
		tr.Update(arr[r].second, 1)
		for tr.val[0] >= m {
			tr.Update(arr[l].second, 0)
			if tr.val[0] < m {
				tr.Update(arr[l].second, 1)
				break
			}

			l++
		}
		if tr.val[0] >= m {
			best = min(best, arr[r].first-arr[l].first)
		}
	}

	return best

}

type Tree struct {
	val  []int // 这个区间内可以贡献多少个数
	cnt  []int // 这个区间内，现在有多少个数
	pref []int
	suf  []int
	sz   int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	cnt := make([]int, 4*n)
	pref := make([]int, 4*n)
	suf := make([]int, 4*n)
	return &Tree{val, cnt, pref, suf, n}
}

func (t *Tree) maintain(i int, l int, mid int, r int) {
	t.cnt[i] = t.cnt[i*2+1] + t.cnt[i*2+2]
	t.pref[i] = t.pref[i*2+1]
	if t.cnt[i*2+1] == mid-l+1 {
		t.pref[i] += t.pref[i*2+2]
	}
	t.suf[i] = t.suf[i*2+2]
	if t.cnt[i*2+2] == r-mid {
		t.suf[i] += t.suf[i*2+1]
	}

	// 但是可能会出现重叠的地方
	t.val[i] = t.val[i*2+1] + t.val[i*2+2]
	if t.suf[i*2+1] != 0 {
		t.val[i] -= (t.suf[i*2+1] + 1) / 2
	}
	if t.pref[i*2+2] != 0 {
		t.val[i] -= (t.pref[i*2+2] + 1) / 2
	}

	w := t.suf[i*2+1] + t.pref[i*2+2]

	t.val[i] += (w + 1) / 2
}

func (t *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			t.val[i] = v
			t.cnt[i] = v
			t.pref[i] = v
			t.suf[i] = v
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		t.maintain(i, l, mid, r)
	}
	f(0, 0, t.sz-1)
}
