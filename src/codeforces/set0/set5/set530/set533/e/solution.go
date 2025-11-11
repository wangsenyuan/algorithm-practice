package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	readString(reader)
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

func solve(s string, t string) int {
	n := len(s)
	if n == 1 {
		// a vs b => ab and ba
		return 2
	}
	dp := make([][3]bool, n+1)
	dp[n][0] = true
	dp[n][1] = true
	// -1
	dp[n][2] = true

	for i := n - 1; i >= 0; i-- {
		if s[i] == t[i] {
			dp[i][0] = dp[i+1][0]
		}
		// 替换t[i] 为 s[i]
		dp[i][1] = dp[i+1][0]
		if (i == n-1 || s[i] == t[i+1]) && dp[i+1][1] {
			dp[i][1] = true
		}
		dp[i][2] = dp[i+1][0]
		if (i == n-1 || t[i] == s[i+1]) && dp[i+1][2] {
			dp[i][2] = true
		}
	}

	// dp[0][0] 肯定是false， 因为 S != T
	var ans int

	var i int
	for s[i] == t[i] {
		i++
	}

	// 到目前为止 s[:i] = t[:i]
	if dp[i][1] {
		ans++
	}
	if dp[i][2] {
		ans++
	}

	return ans
}
