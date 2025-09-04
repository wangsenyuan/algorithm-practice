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
	_, _, res := drive(reader)
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (d int, h []int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &d)
	h = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &h[i])
	}
	res = solve(d, h)
	return
}

func solve(d int, h []int) []int {
	nums := slices.Clone(h)
	slices.Sort(nums)
	nums = slices.Compact(nums)

	m := len(nums)
	n := len(h)

	pos := make([]int, m+1)
	for i := range m + 1 {
		pos[i] = n
	}

	dp := make([]int, m)
	next := make([]int, n)
	for i := range n {
		next[i] = n
	}

	tr := NewSegTree(m)

	for i := n - 1; i >= 0; i-- {
		p := sort.SearchInts(nums, h[i])
		r := sort.SearchInts(nums, h[i]+d)
		if r < m {
			tmp := tr.Get(r, m)
			if dp[p] < tmp.first+1 {
				dp[p] = tmp.first + 1
				next[i] = pos[tmp.second]
			}
		}

		l := sort.SearchInts(nums, h[i]-d)
		// h[l] >= h[i] - d
		if l < m && nums[l] == h[i]-d {
			l++
		}
		if 0 < l {
			tmp := tr.Get(0, l)
			if dp[p] < tmp.first+1 {
				dp[p] = tmp.first + 1
				next[i] = pos[tmp.second]
			}
		}
		dp[p] = max(dp[p], 1)
		tmp := tr.Get(p, p+1)
		if tmp.first < dp[p] {
			tr.Update(p, dp[p])
			pos[p] = i
		}
	}
	best := tr.Get(0, m).second

	var res []int
	for i := pos[best]; i < n; i = next[i] {
		res = append(res, i+1)
	}

	return res
}

type pair struct {
	first  int
	second int
}

func max_pair(a, b pair) pair {
	if a.first > b.first || a.first == b.first && a.second > b.second {
		return a
	}
	return b
}

type SegTree []pair

func NewSegTree(n int) SegTree {
	arr := make([]pair, 2*n)
	for i := n; i < 2*n; i++ {
		arr[i] = pair{0, i - n}
	}
	for i := n - 1; i > 0; i-- {
		arr[i] = max_pair(arr[i<<1], arr[i<<1|1])
	}
	return SegTree(arr)
}

func (t SegTree) Update(p int, v int) {
	p += len(t) / 2
	t[p].first = v
	for p > 1 {
		t[p>>1] = max_pair(t[p], t[p^1])
		p >>= 1
	}
}

func (t SegTree) Get(l int, r int) pair {
	res := pair{0, len(t) / 2}
	l += len(t) / 2
	r += len(t) / 2

	for l < r {
		if l&1 == 1 {
			res = max_pair(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max_pair(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
