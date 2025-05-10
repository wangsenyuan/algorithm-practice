package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	u := readString(reader)
	res := solve(s, u)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

const inf = 1 << 30

func solve(s string, u string) int {
	n := len(s)
	// dp[i][j]表示s[i:]和u[j:]开始相同的最长长度
	m := len(u)

	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, m+1)
	}
	dp[n][m] = 0
	ans := m
	for j := m - 1; j >= 0; j-- {
		for i := n - 1; i >= 0; i-- {
			if s[i] == u[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = dp[i+1][j+1]
			}
			ans = min(ans, m-dp[i][j])
		}
	}

	return ans
}
