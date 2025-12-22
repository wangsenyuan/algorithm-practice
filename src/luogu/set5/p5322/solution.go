package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var s, n, m int
	fmt.Fscan(reader, &s, &n, &m)
	a := make([][]int, s)
	for i := range s {
		a[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(m, a)
}

func solve(m int, a [][]int) int {
	s := len(a)
	n := len(a[0])

	// T = (1 + s) * s / 2 * n 最大得分 = (101) * 50 * 100 = 505000
	// n * m * s 的复杂性，不大行
	dp := make([]int, m+1)
	ndp := make([]int, m+1)

	buf := make([]int, s)

	for j := range n {
		for i := range s {
			buf[i] = a[i][j]
		}
		slices.Sort(buf)

		for i, cur := range buf {
			w := cur*2 + 1
			// 如果选择b[j] = w 的时候
			for m1 := 0; m1+w <= m; m1++ {
				ndp[m1+w] = max(ndp[m1+w], dp[m1]+(i+1)*(j+1))
			}
		}

		for m1 := range m + 1 {
			dp[m1] = max(dp[m1], ndp[m1])
			if m1 > 0 {
				dp[m1] = max(dp[m1], dp[m1-1])
			}
			ndp[m1] = 0
		}
	}

	return dp[m]
}
