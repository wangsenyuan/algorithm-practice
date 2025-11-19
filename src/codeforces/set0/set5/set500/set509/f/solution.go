package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(b)
}

const mod = 1e9 + 7

func add(a, b int) int {
	return (a + b) % mod
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(b []int) int {

	n := len(b)

	dp := make([][]int, n)
	fp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		fp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = -1
			fp[i][j] = -1
		}
	}
	var f func(l int, r int) int
	var g func(l int, r int) int

	// g应该就是组成森林的计数
	g = func(l int, r int) (res int) {
		// l > 0
		if fp[l][r] != -1 {
			return fp[l][r]
		}
		defer func() {
			fp[l][r] = res
		}()

		// 组成一棵树的计数
		res = f(l, r)

		for pos := l + 1; pos <= r; pos++ {
			if b[l] < b[pos] {
				res = add(res, mul(f(l, pos-1), g(pos, r)))
			}
		}
		return
	}

	f = func(l int, r int) int {
		if l == r {
			return 1
		}
		if dp[l][r] != -1 {
			return dp[l][r]
		}
		dp[l][r] = g(l+1, r)
		return dp[l][r]
	}

	return f(0, n-1)
}
