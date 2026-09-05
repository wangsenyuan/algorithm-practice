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

func drive(reader *bufio.Reader) int64 {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(k, a, b)
}

const inf = 1 << 60

func solve(k int, a, b []int) int64 {
	if k&1 == 1 {
		return solve1(a, b)
	}
	var sum int
	best := -inf
	for _, v := range a {
		sum += v
		best = max(best, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return int64(best)
}

func solve1(a []int, b []int) int64 {
	// a[i] + b[i]
	// dp[0] 表示在没有操作时的最优解
	// dp[1] 表示已经有操作时的最优解
	dp := []int{-inf, -inf}
	sum := make([]int, 2)

	for i := range a {
		s0 := sum[0] + a[i]
		s1 := max(sum[1]+a[i], sum[0]+a[i]+b[i])
		dp[0] = max(dp[0], s0)
		dp[1] = max(dp[1], s1)
		sum[0] = max(0, s0)
		sum[1] = max(0, s1)
	}

	return int64(dp[1])
}
