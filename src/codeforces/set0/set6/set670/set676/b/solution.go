package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, t int
	fmt.Fscan(reader, &n, &t)
	return solve(n, t)
}

func solve(n int, t int) int {
	dp := make([][]float64, n)
	for i := range n {
		dp[i] = make([]float64, i+1)
	}
	dp[0][0] = float64(t)

	for i := range n {
		for j := 0; j <= i; j++ {
			if dp[i][j] <= 1 {
				continue
			}
			extra := dp[i][j] - 1
			dp[i][j] = 1
			if i+1 < n {
				dp[i+1][j] += extra / 2
				dp[i+1][j+1] += extra / 2
			}
		}
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if dp[i][j] >= 1 {
				res++
			}
		}
	}
	return res
}
