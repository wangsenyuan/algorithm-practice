package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) []int {
	s := readString(reader)
	return solve(s)
}

func solve(s string) []int {
	n := len(s)

	fp := make([][]bool, n)
	for i := range n {
		fp[i] = make([]bool, n)
	}

	for r := 0; r < n; r++ {
		for l := r; l >= 0; l-- {
			if l == r {
				fp[l][r] = true
			} else if l+1 == r && s[l] == s[r] {
				fp[l][r] = true
			} else if s[l] == s[r] {
				fp[l][r] = fp[l+1][r-1]
			}
		}
	}
	// dp[i][j] 表示区间i...j的最大的阶
	ans := make([]int, n+1)
	dp := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = make([]int, n)
		for j := i; j < n; j++ {
			if fp[i][j] {
				// 至少是1
				dp[i][j] = 1
				if i < j {
					mid := (i + j) / 2
					if (j-i+1)%2 == 0 {
						dp[i][j] = min(dp[i][mid], dp[mid+1][j]) + 1
					} else {
						// (0, 2) => 1, (0, 0) and (2, 2)
						dp[i][j] = min(dp[i][mid-1], dp[mid+1][j]) + 1
					}
				}
				ans[dp[i][j]]++
			}
		}
	}

	for i := n - 1; i > 0; i-- {
		ans[i] += ans[i+1]
	}

	return ans[1:]
}
