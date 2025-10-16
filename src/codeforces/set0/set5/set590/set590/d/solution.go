package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k, s int
	fmt.Fscan(reader, &n, &k, &s)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k, s)
}
func solve(a []int, k int, s int) int {
	// 如果能够把最小的k个人都移动到左边，那么就可以了
	n := len(a)
	if s >= n*(n-1)/2 {
		sort.Ints(a)
		var res int
		for i := range k {
			res += a[i]
		}
		return res
	}
	M := k*(k+1)/2 + s
	// dp[i][j][p] = 前i个数中，选择j个位置，不超过M时的最优解
	dp := make([][]int, k+1)
	for i := range k + 1 {
		dp[i] = make([]int, M+1)
		for j := range M + 1 {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0
	for i := range n {
		// 如果选择i
		for j := min(i+1, k); j > 0; j-- {
			for p := M; p >= i+1; p-- {
				dp[j][p] = min(dp[j][p], dp[j-1][p-i-1]+a[i])
			}
		}
		for j := 0; j <= i+1 && j <= k; j++ {
			for p := 1; p <= M; p++ {
				dp[j][p] = min(dp[j][p], dp[j][p-1])
			}
		}
	}

	return dp[k][M]
}

const inf = 1 << 60

func abs(num int) int {
	return max(num, -num)
}
