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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a, b []int) int {
	n := len(a)
	dp := make([]int, n+1)
	pa := make([]int, n+1)
	pb := make([]int, n+1)
	for i := range n + 1 {
		pa[i] = n
		pb[i] = n
		dp[i] = n
	}

	var ans int
	for i := n - 1; i >= 0; i-- {
		a[i]--
		b[i]--
		pa[a[i]] = i
		pb[b[i]] = i
		if a[i] == b[i] {
			x := a[i] + 1
			if pa[x] == pb[x] {
				dp[i] = dp[pa[x]]
			} else {
				dp[i] = min(pa[x], pb[x])
			}
		}
		if pa[0] != pb[0] {
			ans += min(pa[0], pb[0]) - i
		} else {
			ans += dp[pa[0]] - i
		}
	}

	return ans
}
