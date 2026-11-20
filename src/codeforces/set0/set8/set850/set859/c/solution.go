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
	dp := make([]int, n+1)
	var sum int
	for i := n - 1; i >= 0; i-- {
		sum += a[i]
		dp[i] = max(dp[i+1], sum-dp[i+1])
	}

	return []int{sum - dp[0], dp[0]}
}
