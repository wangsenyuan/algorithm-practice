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
	var n, x int
	fmt.Fscan(reader, &n, &x)
	segs := make([][]int, n)
	for i := range n {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		segs[i] = []int{l, r}
	}
	return solve(x, segs)
}

const inf = 1 << 60

func solve(x int, segs [][]int) int {
	var arr []int
	arr = append(arr, x)
	for _, seg := range segs {
		arr = append(arr, seg[0])
		arr = append(arr, seg[1])
	}

	slices.Sort(arr)
	arr = slices.Compact(arr)

	m := len(arr)
	dp := make([]int, m)
	for i := range m {
		dp[i] = inf
	}
	dp[sort.SearchInts(arr, x)] = 0

	ndp := make([]int, m)

	for _, cur := range segs {
		l, r := cur[0], cur[1]
		i := sort.SearchInts(arr, l)
		j := sort.SearchInts(arr, r)

		for u := range m {
			ndp[u] = inf
		}
		// 计算ndp[u], 如果u在i的前面, 那么要加上 l - arr[u], 如果在j的后面, 要加上 arr[u] - r
		best := dp[0] - arr[0]
		for u := range m {
			// 可以是原来的位置
			best = min(best, dp[u]-arr[u])
			if u < i {
				ndp[u] = min(ndp[u], arr[u]+best+l-arr[u])
			} else if j < u {
				ndp[u] = min(ndp[u], arr[u]+best+arr[u]-r)
			} else {
				ndp[u] = min(ndp[u], arr[u]+best)
			}
		}
		best = dp[m-1] + arr[m-1]
		for u := m - 1; u >= 0; u-- {
			best = min(best, dp[u]+arr[u])
			if j < u {
				ndp[u] = min(ndp[u], best-arr[u]+arr[u]-r)
			} else if u < i {
				ndp[u] = min(ndp[u], best-arr[u]+l-arr[u])
			} else {
				ndp[u] = min(ndp[u], best-arr[u])
			}
		}
		copy(dp, ndp)
	}

	return slices.Min(dp)
}
