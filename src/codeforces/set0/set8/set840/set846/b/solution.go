package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, k, M int
	fmt.Fscan(reader, &n, &k, &M)
	t := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &t[i])
	}
	return solve(n, k, M, t)
}

const inf = 1 << 30

func solve(n int, k int, M int, t []int) int {
	slices.Sort(t)
	s := n * (k + 1)
	dp := make([]int, s+1)
	ndp := make([]int, s+1)
	for i := range s + 1 {
		dp[i] = inf
		ndp[i] = inf
	}
	dp[0] = 0
	for range n {
		for x := range s + 1 {
			var time int
			for j := range k {
				if x+j+1 > s {
					break
				}
				time += t[j]
				ndp[x+j+1] = min(ndp[x+j+1], dp[x]+time)
			}
			if x+k+1 <= s {
				ndp[x+k+1] = min(ndp[x+k+1], dp[x]+time)
			}
		}
		for x := range s + 1 {
			dp[x] = min(dp[x], ndp[x])
			ndp[x] = inf
		}
	}
	for x := s; x >= 0; x-- {
		if dp[x] <= M {
			return x
		}
	}
	return 0
}
