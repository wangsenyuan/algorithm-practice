package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(k, a, b)
}

const inf = 1 << 30

func solve(k int, a []int, b []int) int {
	n := len(a)
	for i := range n {
		b[i] *= k
	}
	max_diff := findMaxDiff(a, b)

	dp := make([]int, 2*max_diff+1)

	for i := range dp {
		dp[i] = -inf
	}
	dp[max_diff] = 0

	ndp := make([]int, 2*max_diff+1)
	copy(ndp, dp)

	for i := range n {
		x, y := a[i], b[i]

		for diff, u := range dp {
			if u < 0 {
				continue
			}
			diff -= max_diff
			newDiff := diff + x - y
			newDiff += max_diff
			ndp[newDiff] = max(ndp[newDiff], u+x)
		}
		copy(dp, ndp)
	}

	res := dp[max_diff]
	if res <= 0 {
		return -1
	}
	return res
}

func findMaxDiff(a []int, b []int) int {
	n := len(a)

	// 先要算出一组最大的diff
	var sum int
	for i := range n {
		sum += a[i]
	}

	dp := make([]int, sum+1)
	fp := make([]int, sum+1)
	for i := range sum + 1 {
		dp[i] = -inf
		fp[i] = inf
	}
	dp[0] = 0
	fp[0] = 0

	for i := range n {
		x, y := a[i], b[i]
		for s := sum - x; s >= 0; s-- {
			dp[s+x] = max(dp[s+x], dp[s]+y)
			fp[s+x] = min(fp[s+x], fp[s]+y)
		}
	}

	var max_diff int
	for s := range sum + 1 {
		if dp[s] > 0 {
			max_diff = max(max_diff, dp[s]-s)
		}
		if fp[s] < inf {
			max_diff = max(max_diff, s-fp[s])
		}
	}
	return max_diff
}
