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
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)

	pos := make([][]int, n+1)
	for i, v := range a {
		pos[v] = append(pos[v], i)
	}

	dp := make([]int, n)
	best := make([]int, n+1)

	for v := 1; v <= n; v++ {
		best[v] = best[v-1]
		k1 := len(pos[v-1])
		k := len(pos[v])

		var w int
		for i, j := k-1, k1; i >= 0; i-- {
			for j > 0 && pos[v-1][j-1] > pos[v][i] {
				w = max(w, dp[pos[v-1][j-1]])
				j--
			}
			dp[pos[v][i]] = w + 1
			if v >= 2 {
				// 考虑不选择v-1, 从i开始选择所有的v
				dp[pos[v][i]] = max(dp[pos[v][i]], best[v-2]+k-i)
			}
			w = max(w, dp[pos[v][i]])
		}
		best[v] = max(best[v], w)
	}

	return n - best[n]
}
