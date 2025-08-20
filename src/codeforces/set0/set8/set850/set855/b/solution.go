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
	var n, p, q, r int
	fmt.Fscan(reader, &n, &p, &q, &r)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}
	return solve(p, q, r, nums)
}

const inf = 1 << 62

func solve(p int, q int, r int, nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = p * nums[i]
		if i > 0 {
			dp[i] = max(dp[i], dp[i-1])
		}
	}
	best := -inf
	suf := -inf
	for i := n - 1; i >= 0; i-- {
		suf = max(suf, r*nums[i])
		best = max(best, dp[i]+q*nums[i]+suf)
	}

	return best
}
