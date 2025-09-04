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
	p := readString(reader)
	res := solve(s, p)
	x := fmt.Sprintf("%v", res)
	fmt.Println(x[1 : len(x)-1])
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

const inf = 1 << 60

func solve(s string, p string) []int {
	m := len(p)
	n := len(s)
	fp := make([]int, n)
	for i := range n {
		r := i
		var j int
		for r < n && j < m {
			if s[r] == p[j] {
				j++
			}
			r++
		}
		if j < m {
			fp[i] = n + 1
		} else {
			// j == m
			fp[i] = r
		}
	}
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, n+1)
		for j := range n + 1 {
			dp[i][j] = -inf
		}
	}

	dp[n][0] = 0

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= n; j++ {
			// 如果不删除s[i]
			dp[i][j] = dp[i+1][j]
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1])
			}
		}
		if fp[i] <= n {
			w := fp[i] - i - m
			for j := w; j <= n; j++ {
				dp[i][j] = max(dp[i][j], dp[fp[i]][j-w]+1)
			}
		}
	}
	return dp[0]
}
