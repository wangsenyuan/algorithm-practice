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

func drive(reader *bufio.Reader) int64 {
	var s, t string
	fmt.Fscan(reader, &s, &t)
	return solve(s, t)
}

func solve(s, t string) int64 {
	// 如果 s[l...r]包含t,的最短序列, 那么 ans += r - l (没有+1)
	// dp[i][j] 表示如果s[:i]匹配t[:j], 的起点最大的位置
	// dp[i][j] = dp[i-1][j] or dp[i-1][j-1] if s[i] = t[j]
	m := len(t)
	dp := make([]int, m)
	for i := range m {
		dp[i] = -1
	}

	var ans int

	for i := range len(s) {
		for j := m - 1; j >= 0; j-- {
			if s[i] == t[j] {
				if j > 0 {
					dp[j] = max(dp[j], dp[j-1])
				} else {
					dp[j] = i
				}
			}
		}
		ans += i - dp[m-1]
	}

	return int64(ans)
}
