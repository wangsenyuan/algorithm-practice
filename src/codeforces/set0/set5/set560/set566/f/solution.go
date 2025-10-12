package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	dp := make([]int, n)

	fp := make([]int, a[n-1]+1)
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		var cnt int
		for w := v; w <= a[n-1]; w += v {
			if fp[w] > i {
				cnt = max(cnt, dp[fp[w]])
			}
		}
		dp[i] = cnt + 1
		fp[v] = i
	}

	return slices.Max(dp)
}
