package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, 2)
		fmt.Fscan(reader, &a[i][0], &a[i][1])
	}
	return solve(n, a)
}

const inf = 1 << 30

func solve(n int, a [][]int) int {
	// n 行2列
	offset := 5 * n

	dp := make([]int, 2*offset+1)
	for i := range 2*offset + 1 {
		dp[i] = inf
	}
	dp[offset] = 0
	ndp := make([]int, 2*offset+1)

	for _, cur := range a {
		for j := range 2*offset + 1 {
			ndp[j] = inf
		}
		x, y := cur[0], cur[1]
		diff := x - y
		for i := -offset; i <= offset; i++ {
			if i+diff >= -offset && i+diff <= offset {
				ndp[i+diff+offset] = min(ndp[i+diff+offset], dp[i+offset])
			}
			if i-diff >= -offset && i-diff <= offset {
				ndp[i-diff+offset] = min(ndp[i-diff+offset], dp[i+offset]+1)
			}
		}
		copy(dp, ndp)
	}

	for d := 0; d <= offset; d++ {
		w := min(dp[d+offset], dp[-d+offset])
		if w < inf {
			return w
		}
	}

	return -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
