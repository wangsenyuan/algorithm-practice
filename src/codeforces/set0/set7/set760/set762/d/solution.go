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
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, 3)
	for i := range 3 {
		a[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

const inf = 1 << 60

func solve(a [][]int) int {
	n := len(a[0])

	dp := make([][]int, 3)
	for i := range 3 {
		dp[i] = make([]int, n+1)
		for j := range n + 1 {
			dp[i][j] = -inf
		}
	}

	dp[0][0] = 0

	getCol := func(j int) int {
		var res int
		for i := range 3 {
			res += a[i][j]
		}
		return res
	}

	for j := range n {
		// 只能往下一列运动
		for i := range 3 {
			switch i {
			case 0:
				dp[0][j+1] = max(dp[0][j+1], dp[i][j]+a[i][j])
				dp[1][j+1] = max(dp[1][j+1], dp[i][j]+a[i][j]+a[i+1][j])
				dp[2][j+1] = max(dp[2][j+1], dp[i][j]+a[i][j]+a[i+1][j]+a[i+2][j])
				if j+2 <= n {
					dp[2][j+2] = max(dp[2][j+2], dp[i][j]+getCol(j)+getCol(j+1))
				}
			case 1:
				dp[0][j+1] = max(dp[0][j+1], dp[i][j]+a[i][j]+a[i-1][j])
				dp[1][j+1] = max(dp[1][j+1], dp[i][j]+a[i][j])
				dp[2][j+1] = max(dp[2][j+1], dp[i][j]+a[i][j]+a[i+1][j])
			default:
				// i == 2
				dp[2][j+1] = max(dp[2][j+1], dp[i][j]+a[i][j])
				dp[1][j+1] = max(dp[1][j+1], dp[i][j]+a[i][j]+a[i-1][j])
				dp[0][j+1] = max(dp[0][j+1], dp[i][j]+a[i][j]+a[i-1][j]+a[i-2][j])
				if j+2 <= n {
					dp[0][j+2] = max(dp[0][j+2], dp[i][j]+getCol(j)+getCol(j+1))
				}
			}
		}
	}

	return dp[2][n]
}
