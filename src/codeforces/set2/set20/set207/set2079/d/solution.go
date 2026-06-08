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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(m, a)
}

func solve(m int, a []int) int {
	arr := slices.Clone(a)
	slices.Sort(arr)
	arr = slices.Compact(arr)
	tr := NewTree(len(arr))

	var best int

	for pos, v := range a {
		// 如果最后一个是v, 那么需要找到m-1的最大的数
		tmp := tr.GetBest(m - 1)
		best = max(best, tmp+v+(pos+1))

		i := sort.SearchInts(arr, v)
		tr.Set(i, v)
	}

	return best
}

type Tree struct {
	val []int
	cnt []int
}

func NewTree(n int) *Tree {
	return &Tree{
		val: make([]int, 4*n),
		cnt: make([]int, 4*n),
	}
}

func (tr *Tree) Set(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.val[i] += v
			tr.cnt[i]++
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.val[i] = tr.val[i*2+1] + tr.val[i*2+2]
		tr.cnt[i] = tr.cnt[i*2+1] + tr.cnt[i*2+2]
	}
	n := len(tr.val) / 4
	f(0, 0, n-1)
}

const inf = 1 << 60

func (tr *Tree) GetBest(k int) int {
	if k == 0 || k > tr.cnt[0] {
		return 0
	}
	var f func(i int, l int, r int, k int) int
	f = func(i int, l int, r int, k int) int {
		if tr.cnt[i] == k {
			return tr.val[i]
		}
		// k > 0
		if l == r {
			return tr.val[i] / tr.cnt[i] * k
		}
		mid := (l + r) >> 1
		if tr.cnt[i*2+2] >= k {
			return f(i*2+2, mid+1, r, k)
		}
		return tr.val[i*2+2] + f(i*2+1, l, mid, k-tr.cnt[i*2+2])
	}
	n := len(tr.val) / 4
	return f(0, 0, n-1, k)
}
