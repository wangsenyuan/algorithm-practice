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

func drive(reader *bufio.Reader) int {
	var n, w, k int
	fmt.Fscan(reader, &n, &w, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	t := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &t[i])
	}
	return solve(w, k, t, a)
}

func solve(w int, k int, t []int, a []int) int {
	n := len(t)
	t1 := make([]int, n)
	for i := range n {
		t1[i] = t[i] - (t[i]+1)/2
	}
	slices.Sort(t1)
	t1 = slices.Compact(t1)

	m := len(t1)
	tr := NewTree(m)

	sum := make([]int, n+1)
	suf := make([]int, n+1)

	var best int
	for l, r := n-1, n-1; l >= 0; l-- {
		v := t[l]
		sum[l] = sum[l+1] + v
		suf[l] = suf[l+1] + a[l]
		j := sort.SearchInts(t1, v-(v+1)/2)
		tr.Update(j, v-(v+1)/2)
		for r >= l {
			tmp := sum[l] - sum[r+1]
			// 找到节省最多的w个歌曲
			tmp -= tr.GetLastKSum(min(w, r-l+1))
			if tmp <= k {
				best = max(best, suf[l]-suf[r+1])
				break
			}
			// tmp > k
			u := t[r] - (t[r]+1)/2
			j1 := sort.SearchInts(t1, u)
			tr.Update(j1, -u)
			r--
		}
	}

	return best
}

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
			t.val[i] += v
			if v > 0 {
				t.cnt[i]++
			} else {
				t.cnt[i]--
			}
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		t.val[i] = t.val[i*2+1] + t.val[i*2+2]
		t.cnt[i] = t.cnt[i*2+1] + t.cnt[i*2+2]
	}
	f(0, 0, t.sz-1)
}

func (t *Tree) GetLastKSum(k int) int {
	var f func(i int, l int, r int, k int) int
	f = func(i int, l int, r int, k int) int {
		if l == r {
			return t.val[i] / t.cnt[i] * k
		}

		mid := (l + r) >> 1

		if t.cnt[i*2+2] >= k {
			return f(i*2+2, mid+1, r, k)
		}

		return t.val[i*2+2] + f(i*2+1, l, mid, k-t.cnt[i*2+2])
	}
	return f(0, 0, t.sz-1, k)
}
