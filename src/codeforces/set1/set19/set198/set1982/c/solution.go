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
	var n, L, R int
	fmt.Fscan(reader, &n, &L, &R)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(L, R, a)
}

func solve(L int, R int, a []int) int {
	n := len(a)
	var score int

	var sum int
	var j int

	for i := range n {
		sum += a[i]
		if sum < L {
			continue
		}
		for sum > R {
			sum -= a[j]
			j++
		}
		if sum >= L {
			score++
			j = i + 1
			sum = 0
		}
	}
	return score
}

func solve1(L int, R int, a []int) int {
	n := len(a)
	sum := make([]int, n+1)

	dp := make([]int, n+1)
	var j int
	for i := range n {
		dp[i+1] = dp[i]
		sum[i+1] = sum[i] + a[i]
		for sum[i+1]-sum[j] >= L {
			j++
		}
		if j > 0 && sum[i+1]-sum[j-1] >= L && sum[i+1]-sum[j-1] <= R {
			dp[i+1] = max(dp[i+1], dp[j-1]+1)
		}
	}

	return dp[n]
}
