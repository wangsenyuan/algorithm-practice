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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, 2)
		fmt.Fscan(reader, &s[i][0], &s[i][1])
	}
	return solve(s)
}

func solve(s [][]int) int {
	// n := len(s)
	var sum int
	var arr []int
	var sumL int
	for _, cur := range s {
		l, r := cur[0], cur[1]
		sum += r - l
		// 如果将 -l 换成r, 变化 = -(-l) + r = r + l
		arr = append(arr, r+l)
		sumL += l
	}

	slices.Sort(arr)
	arr = slices.Compact(arr)
	tr := NewTree(len(arr))

	for _, cur := range s {
		l, r := cur[0], cur[1]
		i := sort.SearchInts(arr, l+r)
		tr.Update(i, l+r)
	}

	n := len(s)
	if n&1 == 0 {
		return sum - sumL + tr.GetBest(n/2)
	}

	if n == 1 {
		return sum
	}

	var best int

	for _, cur := range s {
		// 如果cur不算入
		l, r := cur[0], cur[1]
		i := sort.SearchInts(arr, l+r)
		tr.Update(i, -(l + r))
		tmp := -(sumL - l) + tr.GetBest(n/2) + sum
		best = max(best, tmp)
		tr.Update(i, l+r)
	}

	return best
}

type Tree struct {
	cnt []int
	sum []int
	sz  int
}

func NewTree(n int) *Tree {
	t := new(Tree)
	t.cnt = make([]int, 4*n)
	t.sum = make([]int, 4*n)
	t.sz = n
	return t
}

func (t *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			t.sum[i] += v

			if v < 0 {
				t.cnt[i]--
			} else {
				t.cnt[i]++
			}
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(2*i+1, l, mid)
		} else {
			f(2*i+2, mid+1, r)
		}
		t.sum[i] = t.sum[2*i+1] + t.sum[2*i+2]
		t.cnt[i] = t.cnt[2*i+1] + t.cnt[2*i+2]
	}
	f(0, 0, t.sz-1)
}

func (t *Tree) GetBest(k int) int {
	if k == 0 {
		return 0
	}
	var f func(i int, l int, r int, k int) int

	f = func(i int, l int, r int, k int) int {
		if l == r {
			return t.sum[i] / t.cnt[i] * k
		}
		mid := (l + r) / 2
		if t.cnt[2*i+2] >= k {
			return f(2*i+2, mid+1, r, k)
		}
		return t.sum[2*i+2] + f(2*i+1, l, mid, k-t.cnt[2*i+2])
	}

	return f(0, 0, t.sz-1, k)
}
