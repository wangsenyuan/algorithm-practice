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
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
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
	if n == 1 {
		return a[0]
	}

	var res int
	for i := range n {
		if i&1 == 0 {
			res += a[i]
		} else {
			res -= a[i]
		}
	}
	var best int
	if n&1 == 1 {
		best = n - 1
	} else {
		best = n - 2
	}

	dp := []int{-inf, -inf}

	for i := range n {
		if i&1 == 0 {
			// a[i] 和 a[j]交换， 且j&1 == 1
			best = max(best, -2*a[i]+i+dp[1])
			// 如果把i换到j位，要从+变成-
			dp[0] = max(dp[0], -2*a[i]-i)
		} else {
			best = max(best, 2*a[i]+i+dp[0])
			dp[1] = max(dp[1], 2*a[i]-i)
		}
	}

	return res + best
}

const inf = 1 << 60
