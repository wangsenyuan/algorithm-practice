package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)

	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)

	if n == 1 {
		return a[0]
	}

	dp := make([]int, n+1)
	dp[1] = a[0]
	for i := 1; i < n; i++ {
		dp[i+1] = min(dp[i]+a[i]-1, dp[i-1]+a[i-1]+max(0, a[i]-i))
	}

	return dp[n]
}
