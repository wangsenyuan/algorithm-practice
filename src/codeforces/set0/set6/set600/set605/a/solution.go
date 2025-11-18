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
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	pos := make([]int, n)
	for i, v := range a {
		pos[v-1] = i
	}

	// 最长的连续递增序列

	dp := make([]int, n)
	for i, v := range a {
		v--
		dp[v] = 1
		if v > 0 && pos[v-1] < i {
			dp[v] = dp[v-1] + 1
		}
	}

	best := slices.Max(dp)
	
	return n - best
}
