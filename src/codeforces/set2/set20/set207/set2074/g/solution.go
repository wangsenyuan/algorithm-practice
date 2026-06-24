package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		fmt.Fprintln(writer, drive(reader))
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

const inf = 1 << 60

func solve(a []int) int {
	n := len(a)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
	}

	for d := 3; d <= n; d++ {
		for i := range n {
			j := (i + d - 1) % n
			var best int
			for d1 := 2; d1 < d; d1++ {
				mid := (i + d1 - 1) % n
				cur := a[i] * a[mid] * a[j]
				if d1 > 2 {
					cur += dp[(i+1)%n][(mid-1+n)%n]
				}
				if d1+1 < d {
					cur += dp[(mid+1)%n][(j-1+n)%n]
				}
				best = max(best, cur)
			}
			for d1 := 1; d1 < d; d1++ {
				mid := (i + d1 - 1) % n
				best = max(best, dp[i][mid]+dp[(mid+1)%n][j])
			}
			dp[i][j] = best
		}
	}

	var best int

	for i := range n {
		for j := range n {
			best = max(best, dp[i][j])
		}
	}

	return best
}
