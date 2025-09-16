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
	_, res := drive(reader)
	fmt.Println(len(res))
	if len(res) > 0 {
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
}

func drive(reader *bufio.Reader) (a []int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a)
	return
}

func solve(a []int) []int {
	n := len(a)

	nums := slices.Clone(a)
	slices.Sort(nums)
	nums = slices.Compact(nums)

	m := len(nums)
	pos := make([][]int, m)

	// 要找出，出现在区间k+1...i-1位置中间的，且在k之前也出现了数的最大位置
	// 从k的角度看，就是找到一个区间l...r (a[l] = a[r]) 且 l < k and k < r
	// 最大的l
	fp := make([]int, n)
	for i := range n {
		fp[i] = -1
	}

	for i := 0; i < n; i++ {
		j := sort.SearchInts(nums, a[i])
		pos[j] = append(pos[j], i)
		if len(pos[j]) >= 4 {
			fp[i] = pos[j][len(pos[j])-4]
		}
	}

	tr := NewSegTree(n)

	for i, v := range a {
		j := sort.SearchInts(nums, v)

		pos[j] = pos[j][1:]
		if len(pos[j]) > 0 {
			r := pos[j][0]
			fp[r] = max(fp[r], fp[i])
			if i+1 <= r-1 {
				fp[r] = max(fp[r], tr.Get(i+1, r))
			}
			tr.Update(r, i)
		}
	}

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i] = dp[i-1]
		}
		j := fp[i]
		if j >= 0 {
			val := 4
			if j > 0 {
				val += dp[j-1]
			}
			if val > dp[i] {
				dp[i] = val
			}
		}
	}

	var res []int

	for i := n - 1; i > 0; {
		if dp[i] == dp[i-1] {
			i--
			continue
		}
		// dp[i].val > dp[i-1].val
		j := fp[i]
		res = append(res, a[i], a[j], a[i], a[j])
		i = j - 1
	}

	slices.Reverse(res)
	return res
}

const inf = 1 << 60

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = -inf
	}
	return SegTree(arr)
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = v

	for p > 1 {
		tr[p>>1] = max(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	var res int = -inf
	for l < r {
		if l&1 == 1 {
			res = max(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
