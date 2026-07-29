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
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

const inf = 1 << 60

func solve(s string) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, n+1)
		for j := range n + 1 {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0

	for i := range n {
		ndp := make([][]int, n+1)
		for j := range n + 1 {
			ndp[j] = make([]int, n+1)
			for k := range n + 1 {
				ndp[j][k] = inf
			}
		}

		if s[i] != 'T' {
			// F or N
			for f := range n {
				for s := range n {
					ndp[f+1][s+1] = min(ndp[f+1][s+1], max(dp[f][s], s+1))
				}
			}
		}

		if s[i] != 'F' {
			// T or N
			for f := range n + 1 {
				for s := range n + 1 {
					ndp[f][max(s-1, 0)] = min(ndp[f][max(s-1, 0)], dp[f][s])
				}
			}
		}
		dp = ndp
	}

	var ans int
	for f := range n + 1 {
		for s := range n + 1 {
			ans = max(ans, f-dp[f][s])
		}
	}
	return ans
}
