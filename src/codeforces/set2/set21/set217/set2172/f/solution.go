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
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(p)
}

func solve(p []int) int {
	n := len(p)
	dp := make([]int, n)
	dp[0] = p[0]
	for i := 1; i < n; i++ {
		dp[i] = gcd(dp[i-1], p[i])
	}
	res := dp[n-1]

	fp := p[n-1]
	for i := n - 2; i > 0; i-- {
		fp = gcd(fp, p[i])
		res += min(dp[i], fp)
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
