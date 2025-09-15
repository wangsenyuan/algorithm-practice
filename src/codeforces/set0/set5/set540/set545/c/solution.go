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
	trees := make([][]int, n)
	for i := range n {
		trees[i] = make([]int, 2)
		for j := range 2 {
			fmt.Fscan(reader, &trees[i][j])
		}
	}
	return solve(trees)
}

func solve(trees [][]int) int {
	// trees already sorted by x
	dp := make([]int, 3)
	// dp[0] fell to left, dp[1] fell to right, dp[2] stand
	dp[0] = 1
	// 第一个始终fell to left
	n := len(trees)
	for i := 1; i < n; i++ {
		ndp := make([]int, 3)
		ndp[2] = max(dp[0], dp[1], dp[2])

		x0 := trees[i][0] - trees[i][1]
		if trees[i-1][0] < x0 {
			ndp[0] = max(dp[0], dp[2]) + 1
		}
		if trees[i-1][0]+trees[i-1][1] < x0 {
			ndp[0] = max(ndp[0], dp[1]+1)
		}

		x1 := trees[i][0] + trees[i][1]

		if i == n-1 || x1 < trees[i+1][0] {
			// 往后倒的时候，只考虑不撞到后面的树
			ndp[1] = max(dp[0], dp[2]) + 1
			if trees[i-1][0]+trees[i-1][1] < trees[i][0] {
				ndp[1] = max(ndp[1], dp[1]+1)
			}
		}

		copy(dp, ndp)
	}

	return max(dp[0], dp[1], dp[2])
}
