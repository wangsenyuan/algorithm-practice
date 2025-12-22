package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	score, res := drive(reader)
	fmt.Println(score)
	fmt.Println(res[0], res[1], res[2], res[3])
}

func drive(reader *bufio.Reader) (int, []int) {
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

func solve(a [][]int) (score int, res []int) {

	n := len(a)

	dp := make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, n)
		for j := range n {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := range n {
		for j := range n {
			dp[i][j][0] = a[i][j]
			if i > 0 && j > 0 {
				dp[i][j][0] += dp[i-1][j-1][0]
			}
			dp[i][j][1] = a[i][j]
			if i > 0 && j+1 < n {
				dp[i][j][1] += dp[i-1][j+1][1]
			}
		}
	}

	fp := make([][]int, n)
	for i := range n {
		fp[i] = make([]int, 2)
	}

	best := make([]int, 2)
	pos := make([][]int, 2)
	for i := n - 1; i >= 0; i-- {
		for j := range n {
			tmp := dp[i][j][0] + dp[i][j][1] - a[i][j]
			if j > 0 {
				tmp += fp[j-1][0]
			}
			if j+1 < n {
				tmp += fp[j+1][1]
			}
			p := (i + j) % 2
			if tmp >= best[p] {
				best[p] = tmp
				pos[p] = []int{i + 1, j + 1}
			}
		}
		for j := n - 1; j > 0; j-- {
			fp[j][0] = a[i][j] + fp[j-1][0]
		}
		fp[0][0] = a[i][0]
		for j := 0; j+1 < n; j++ {
			fp[j][1] = a[i][j] + fp[j+1][1]
		}
		fp[n-1][1] = a[i][n-1]
	}

	return best[0] + best[1], append(pos[0], pos[1]...)
}
