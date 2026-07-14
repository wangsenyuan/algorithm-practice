package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for _, v := range drive(reader) {
		fmt.Println(v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	return solve(n, m)
}

const mod = 998244353

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

func mul(a, b int) int {
	return a * b % mod
}

func solve(n, m int) []int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := n; j >= max(1, (i-1)/m); j-- {
			dp[j] = add(dp[j-1], mul(dp[j], j-(i-1)/m))
		}
		dp[0] = 0
	}

	return dp[1:]
}
