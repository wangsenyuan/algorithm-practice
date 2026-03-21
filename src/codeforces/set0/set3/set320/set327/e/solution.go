package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	var k int
	fmt.Fscan(reader, &k)
	unlucky := make([]int, k)
	for i := range unlucky {
		fmt.Fscan(reader, &unlucky[i])
	}
	fmt.Println(solve(a, unlucky))
}

const mod = 1000000007

func solve(a, unlucky []int) int {
	n := len(a)
	full := 1 << n
	dp := make([]int, full)

	// Subset sums via lowest-bit decomposition
	for i := 1; i < full; i++ {
		j := bits.TrailingZeros(uint(i))
		dp[i] = dp[i^(1<<j)] + a[j]
	}

	// Mark subsets whose sum hits a bad point
	for i := range dp {
		if slices.Contains(unlucky, dp[i]) {
			dp[i] = -1
		}
	}

	// dp[mask] = valid orderings of elements in mask with no bad prefix sum
	dp[0] = 1
	for i := 1; i < full; i++ {
		if dp[i] < 0 {
			dp[i] = 0
			continue
		}
		dp[i] = 0
		for j := range n {
			if i&(1<<j) != 0 {
				dp[i] += dp[i^(1<<j)]
			}
		}
		dp[i] %= mod
	}
	return dp[full-1]
}
