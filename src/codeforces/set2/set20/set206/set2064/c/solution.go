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
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)

	var best int

	dp := make([]int, n)
	for i := range n {
		if i > 0 {
			dp[i] = dp[i-1]
		}
		if a[i] > 0 {
			dp[i] += a[i]
		}
		best = dp[i]
	}

	var fp int
	for i := n - 1; i >= 0; i-- {
		if a[i] < 0 {
			fp -= a[i]
		}
		if i > 0 {
			best = max(best, fp+dp[i-1])
		} else {
			best = max(best, fp)
		}
	}

	return best
}

func abs(num int) int {
	return max(num, -num)
}
