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

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(s, a)
}

const inf = 1 << 60

func solve(s string, a []int) int64 {
	n := len(s)

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	var f func(start int, end int, prefix int) (res int)
	f = func(start int, end int, prefix int) (res int) {
		if start >= end {
			return 0
		}
		if start+1 == end {
			return a[prefix-1]
		}
		if dp[start][end][prefix] > 0 {
			return dp[start][end][prefix]
		}

		defer func() {
			dp[start][end][prefix] = res
		}()

		res = a[prefix-1] + f(start+1, end, 1)

		for i := start + 1; i < end; i++ {
			if s[i] == s[start] {
				res = max(res, f(start+1, i, 1)+f(i, end, prefix+1))
			}
		}
		return
	}

	return int64(f(0, n, 1))
}
