package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	fmt.Println(res)
}

func solve(n int) int {
	// 4 * 3 * pow(4, n - 3)
	res := 2 * 4 * 3 * (1 << (2 * (n - 3)))
	// 4 * 3 * 3 * pow(4, n - 4)
	if n >= 4 {
		// n - 3 个起点位置
		res += (n - 3) * 4 * 3 * 3 * (1 << (2 * (n - 4)))
	}
	return res
}

func solve1(n int) int {
	// n <= 30， 2 *n- 2
	// 4种，正好连续n个相同的，剩余 2 * n - 2 - n = n - 2 个位置,
	// dp[i][j] 表示到i为止，切好有j个数和a[i]相同的case
	dp := make([][2]int, n+1)
	ndp := make([][2]int, n+1)
	dp[1][0] = 4
	for i := 2; i <= 2*n-2; i++ {
		// 和前面相同的情况
		for j := 2; j <= n; j++ {
			ndp[j][0] = dp[j-1][0]
			ndp[j][1] = dp[j-1][1]
			if j == n {
				ndp[j][1] += dp[j-1][0]
			}
		}
		// 和前面不同的情况
		var sum0, sum1 int
		for j := 1; j < n; j++ {
			sum0 += dp[j][0]
			sum1 += dp[j][1]
		}
		sum1 += dp[n][1]

		ndp[1][0] = sum0 * 3
		ndp[1][1] = sum1 * 3

		for j := 1; j <= n; j++ {
			dp[j][0] = ndp[j][0]
			dp[j][1] = ndp[j][1]
			ndp[j][0] = 0
			ndp[j][1] = 0
		}
	}

	var res int
	for j := 1; j <= n; j++ {
		res += dp[j][1]
	}

	return res
}
