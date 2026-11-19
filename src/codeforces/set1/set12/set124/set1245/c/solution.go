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
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

const mod = 1_000_000_007

func add(a int, b int) int {
	return (a + b) % mod
}

func solve(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := range s {
		if s[i] == 'm' || s[i] == 'w' {
			return 0
		}

		dp[i+1] = dp[i]

		if i > 0 && s[i] == s[i-1] && (s[i] == 'u' || s[i] == 'n') {
			dp[i+1] = add(dp[i+1], dp[i-1])
		}
	}

	return dp[len(s)]
}
