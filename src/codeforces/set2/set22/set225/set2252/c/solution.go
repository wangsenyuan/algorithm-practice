package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	v := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &v[i])
	}
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(v, a)
}

func solve(v []int, a [][]int) int {
	var nums []int
	for _, row := range a {
		nums = append(nums, row...)
	}
	slices.Sort(nums)
	nums = slices.Compact(nums)
	tr := build(len(nums))

	for _, row := range a {
		for _, v := range row {
			i := sort.SearchInts(nums, v)
			tr.update(i, v)
		}
	}
	best := len(a[0])

	for i, row := range a {
		// 找出这一行及以下的, 最少的sum >= v
		w := tr.find(v[i])
		if w > 0 {
			best = min(best, w)
		}
		for _, v := range row {
			i := sort.SearchInts(nums, v)
			tr.update(i, -v)
		}
	}

	return best
}

type tree struct {
	sum []int
	cnt []int
	sz  int
}

func build(n int) *tree {
	h := bits.Len(uint(n))
	sum := make([]int, 2<<h)
	cnt := make([]int, 2<<h)
	return &tree{sum, cnt, n}
}

func (tr *tree) update(pos int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			tr.sum[i] += v
			if v < 0 {
				tr.cnt[i]--
			} else {
				tr.cnt[i]++
			}
			return
		}
		mid := (l + r) >> 1
		if pos < mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid, r)
		}
		tr.cnt[i] = tr.cnt[2*i+1] + tr.cnt[i*2+2]
		tr.sum[i] = tr.sum[2*i+1] + tr.sum[i*2+2]
	}
	f(0, 0, tr.sz)
}

// 找到 suf sum >= val 的最大的位置
func (tr *tree) find(val int) int {
	if tr.sum[0] < val {
		return -1
	}
	// tr.sum[0] >= val
	var f func(i int, l int, r int, v int) int
	f = func(i int, l int, r int, v int) int {
		if v == 0 {
			return 0
		}
		// tr.sum[i] >= v must holds
		// v > 0
		if l+1 == r {
			w := tr.sum[i] / tr.cnt[i]
			// 有可能不能整除
			return (v + w - 1) / w
		}
		mid := (l + r) / 2
		if tr.sum[i*2+2] >= v {
			return f(i*2+2, mid, r, v)
		}
		return tr.cnt[i*2+2] + f(i*2+1, l, mid, v-tr.sum[i*2+2])
	}
	return f(0, 0, tr.sz, val)
}
