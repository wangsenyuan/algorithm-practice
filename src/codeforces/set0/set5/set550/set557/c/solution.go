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
	l := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &l[i])
	}
	d := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &d[i])
	}
	return solve(l, d)
}

type data struct {
	height int
	price  int
	id     int
}

func solve(l []int, d []int) int {
	n := len(l)
	arr := make([]data, n)
	for i := range n {
		arr[i] = data{l[i], d[i], i}
	}

	slices.SortFunc(arr, func(a, b data) int {
		return cmp.Or(a.height-b.height, a.id-b.id)
	})

	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = dp[i+1] + arr[i].price
	}

	nums := slices.Clone(arr)

	slices.SortFunc(nums, func(a, b data) int {
		return cmp.Or(a.price-b.price, a.id-b.id)
	})

	search := func(p data) int {
		return sort.Search(n, func(i int) bool {
			return nums[i].price > p.price || nums[i].price == p.price && nums[i].id >= p.id
		})
	}

	fp := NewTree(n)

	res := dp[1]
	for i := 0; i < n; {
		j := i
		for i < n && arr[i].height == arr[j].height {
			i++
		}
		// (i - j) * 2 > i - l
		// l > 2 * j - i
		l := min(j, max(0, 2*j-i+1))
		tmp := dp[i]
		if l > 0 {
			tmp += fp.FindKthPrefixSum(l)
		}
		res = min(res, tmp)
		for j < i {
			p := search(arr[j])
			fp.Update(p, arr[j].price)
			j++
		}
	}

	return res
}

type Tree struct {
	sum []int
	cnt []int
}

func NewTree(n int) *Tree {
	return &Tree{
		sum: make([]int, 4*n),
		cnt: make([]int, 4*n),
	}
}

func (tr *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.sum[i] += v
			tr.cnt[i]++
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.sum[i] = tr.sum[i*2+1] + tr.sum[i*2+2]
		tr.cnt[i] = tr.cnt[i*2+1] + tr.cnt[i*2+2]
	}
	n := len(tr.sum) / 4
	f(0, 0, n-1)
}

func (tr *Tree) FindKthPrefixSum(k int) int {
	var f func(i int, l int, r int, k int) int
	f = func(i int, l int, r int, k int) int {
		if l == r || k == tr.cnt[i] {
			// tr.cnt[i] > 0
			return tr.sum[i]
		}
		mid := (l + r) >> 1
		if k <= tr.cnt[i*2+1] {
			return f(i*2+1, l, mid, k)
		}
		return tr.sum[i*2+1] + f(i*2+2, mid+1, r, k-tr.cnt[i*2+1])
	}
	n := len(tr.sum) / 4
	return f(0, 0, n-1, k)
}
