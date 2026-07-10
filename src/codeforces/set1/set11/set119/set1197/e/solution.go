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
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	dolls := make([][]int, n)
	for i := range n {
		dolls[i] = make([]int, 2)
		fmt.Fscan(reader, &dolls[i][0], &dolls[i][1])
	}
	return solve(dolls)
}

const inf = 1 << 60

const mod = 1e9 + 7

func solve(dolls [][]int) int {
	var nums []int
	for _, cur := range dolls {
		nums = append(nums, cur[0], cur[1])
	}
	slices.Sort(nums)
	nums = slices.Compact(nums)

	k := len(nums)
	todo1 := make([][]int, k)
	todo2 := make([][]int, k)
	suf := make([]int, k+1)
	for i, cur := range dolls {
		l := sort.SearchInts(nums, cur[1])
		todo1[l] = append(todo1[l], i)
		r := sort.SearchInts(nums, cur[0])
		todo2[r] = append(todo2[r], i)

		suf[l]++
	}

	for i := k - 1; i >= 0; i-- {
		suf[i] += suf[i+1]
	}

	n := len(dolls)
	tr := make(SegTree, 2*k)
	for i := range 2 * k {
		tr[i].val = inf
	}

	dp := make([]data, n)
	for i := range n {
		dp[i] = data{dolls[i][1], 1}
	}

	for i := range k {
		for _, j := range todo2[i] {
			out := dolls[j][0]
			j2 := sort.SearchInts(nums, out)
			tmp := data{dp[j].val - out, dp[j].ways}
			tr.update(j2, tmp)
		}

		for _, j := range todo1[i] {
			in := dolls[j][1]
			// 要找到所有 out <=in中, 最小的 dp[?] - out[?]
			j1 := sort.SearchInts(nums, in)
			tmp := tr.get(0, j1+1)
			tmp.val += in
			dp[j] = mergeData(dp[j], tmp)
		}
	}

	check := func(out int) bool {
		i := sort.SearchInts(nums, out)
		return suf[i] == 0
	}

	var minExtraSpace = inf
	for i := range n {
		if check(dolls[i][0]) {
			minExtraSpace = min(minExtraSpace, dp[i].val)
		}
	}

	var ans int
	for i := range n {
		if dp[i].val == minExtraSpace && check(dolls[i][0]) {
			ans += dp[i].ways
			ans %= mod
		}
	}

	return ans
}

type data struct {
	val  int
	ways int
}

func mergeData(a data, b data) data {
	if a.val < b.val {
		return a
	}
	if a.val > b.val {
		return b
	}
	return data{
		val:  a.val,
		ways: (a.ways + b.ways) % mod,
	}
}

type SegTree []data

func (s SegTree) update(p int, v data) {
	n := len(s) / 2
	p += n
	if s[p].val < v.val {
		return
	}
	s[p] = mergeData(s[p], v)
	for p > 1 {
		s[p>>1] = mergeData(s[p], s[p^1])
		p >>= 1
	}
}

func (s SegTree) get(l int, r int) data {
	n := len(s) / 2
	l += n
	r += n
	res := data{
		val:  inf,
		ways: 0,
	}
	for l < r {
		if l&1 == 1 {
			res = mergeData(res, s[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = mergeData(res, s[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
