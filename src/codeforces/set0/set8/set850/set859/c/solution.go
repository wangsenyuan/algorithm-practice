package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res[0], res[1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	n := len(a)
	suf := make([]int, n+1)
	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + a[i]
		// dp[i] = 0
		for j := i; j < n; j++ {
			dp[i] = max(dp[i], suf[j]-dp[j+1])
		}
	}

	return []int{suf[0] - dp[0], dp[0]}
}
