package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func solve(s string) int {
	n := len(s)
	f := make([][]int, n+1)
	for i := range n + 1 {
		f[i] = make([]int, n+1)
	}
	f[n][n] = 0
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			if s[i] != s[j] {
				f[i][j] = 0
			} else {
				f[i][j] = 1 + f[i+1][j+1]
			}
		}
	}
	//dp[i][j] 表示以i结尾，且长度为j的子串的方案
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n+1)
	}

	for r := 0; r < n; r++ {
		for w := 1; w <= r; w++ {
			i := r - w
			if s[i+1] != '0' {
				dp[r][w] = add(dp[r][w], dp[i][w-1])
			}
			l := r - 2*w + 1
			if l >= 0 && s[l] != '0' && s[i+1] != '0' {
				k := f[l][i+1]
				if k < w && s[l+k] < s[i+1+k] {
					dp[r][w] = add(dp[r][w], sub(dp[i][w], dp[i][w-1]))
				}
			}
		}
		dp[r][r+1] = 1
		for j := 1; j <= n; j++ {
			dp[r][j] = add(dp[r][j], dp[r][j-1])
		}
	}

	return dp[n-1][n]
}
