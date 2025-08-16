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
	var n, W, B, X int
	fmt.Fscan(reader, &n, &W, &B, &X)
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	cost := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &cost[i])
	}
	return solve(W, X, B, c, cost)
}

const inf = 1 << 60

func solve(W int, X int, B int, c []int, cost []int) int {
	var sum_c int
	for _, v := range c {
		sum_c += v
	}
	ndp := make([]int, sum_c+1)

	dp := make([]int, sum_c+1)
	for i := range dp {
		dp[i] = -inf
		ndp[i] = -inf
	}
	dp[0] = W
	ndp[0] = W

	for i, v := range c {
		u := cost[i]
		for k := 0; k <= v; k++ {
			for j := k; j <= sum_c; j++ {
				if dp[j-k] >= 0 {
					ndp[j] = max(ndp[j], min(dp[j-k]+X, W+(j-k)*B)-u*k)
				}
			}
		}
		copy(dp, ndp)
	}
	for j := sum_c; j >= 0; j-- {
		if dp[j] >= 0 {
			return j
		}
	}
	return 0
}
