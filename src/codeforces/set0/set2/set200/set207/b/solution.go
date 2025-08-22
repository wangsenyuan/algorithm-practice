package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func min_pair(a, b pair) pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

const inf = 1 << 60

func solve(a []int) int {
	n := len(a)
	nums := make([]int, 2*n)
	for i := range 2 * n {
		nums[i] = max(0, i-a[i%n])
	}
	seg := NewSegTree(nums)
	fa := make([][]int, 2*n)
	h := bits.Len(uint(2 * n))
	for i := range 2 * n {
		fa[i] = make([]int, h)
		fa[i][0] = seg.Query(nums[i], i+1).second
	}
	for j := 1; j < h; j++ {
		for i := range 2 * n {
			fa[i][j] = fa[fa[i][j-1]][j-1]
		}
	}
	var ans int
	for i := n; i < 2*n; i++ {
		if i-nums[i] >= n-1 {
			ans++
			continue
		}
		u := i
		for j := h - 1; j >= 0; j-- {
			if i-nums[fa[u][j]] < n-1 {
				ans += 1 << j
				u = fa[u][j]
			}
		}
		ans += 2
	}
	return ans
}

type SegTree []pair

func NewSegTree(nums []int) SegTree {
	n := len(nums)
	arr := make([]pair, 2*n)
	for i := n; i < 2*n; i++ {
		arr[i] = pair{nums[i-n], i - n}
	}
	for i := n - 1; i > 0; i-- {
		arr[i] = min_pair(arr[i<<1], arr[i<<1|1])
	}
	return SegTree(arr)
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p].first = v
	for p > 1 {
		tr[p>>1] = min_pair(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Query(l int, r int) pair {
	l += len(tr) / 2
	r += len(tr) / 2
	res := pair{inf, -1}
	for l < r {
		if l&1 == 1 {
			res = min_pair(tr[l], res)
			l++
		}
		if r&1 == 1 {
			r--
			res = min_pair(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
