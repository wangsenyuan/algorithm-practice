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
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	companies := make([][]int, m)
	for i := range m {
		companies[i] = make([]int, 3)
		fmt.Fscan(reader, &companies[i][0], &companies[i][1], &companies[i][2])
	}
	return solve(n, k, companies)
}

func solve(n int, k int, companies [][]int) int {
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, n+1)
		for j := i; j <= n; j++ {
			dp[i][j] = inf
		}
	}

	for _, cur := range companies {
		l, r, c := cur[0]-1, cur[1]-1, cur[2]
		dp[l][r] = min(dp[l][r], c)
	}

	for r := 0; r < n; r++ {
		for l := r - 1; l >= 0; l-- {
			// dp[l][r] = dp[l][i] + dp[j][r] (j <= i + 1)
			tmp := inf
			for i := l; i < r; i++ {
				tmp = min(tmp, dp[i+1][r])
				dp[l][r] = min(dp[l][r], dp[l][i]+tmp)
			}
		}
		for l := 1; l <= r; l++ {
			dp[l][r] = min(dp[l][r], dp[l-1][r])
		}
	}
	// fp[r][j] = fp[r][i] + fp[r-i][j-i]
	fp := make([][]int, n+1)
	for i := range n + 1 {
		fp[i] = make([]int, n+1)
		for j := range n + 1 {
			fp[i][j] = inf
		}
	}
	fp[0][0] = 0
	for r := 1; r <= n; r++ {
		for j := 0; j <= r; j++ {
			// l = r - 1 - j + 1, 正好j个人
			fp[r][j] = min(fp[r-1][j], dp[r-j][r-1])
			for i := j - 1; i > 0; i-- {
				fp[r][j] = min(fp[r][j], dp[r-i][r-1]+fp[r-i][j-i])
			}
		}
	}
	ans := inf
	for i := k; i <= n; i++ {
		ans = min(ans, fp[n][i])
	}
	if ans == inf {
		return -1
	}
	return ans
}

const inf = 1 << 60
