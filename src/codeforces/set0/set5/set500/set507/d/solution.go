package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k, m int
	fmt.Fscan(reader, &n, &k, &m)
	fmt.Println(solve(n, k, m))
}

func solve(n int, k int, m int) int {
	if m == 1 {
		// 不管结果是多少
		return 0
	}

	dp := make([][][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([][]int, k)
		for j := range k {
			dp[i][j] = make([]int, 2)
			for s := range 2 {
				dp[i][j][s] = -1
			}
		}
	}

	pw := make([]int, n)
	pw[0] = 1
	pw2 := make([]int, n)
	pw2[0] = 1
	for i := 1; i < n; i++ {
		pw[i] = pw[i-1] * 10 % m
		pw2[i] = pw2[i-1] * 10 % k
	}

	var f func(ind int, rem int, change int) int

	f = func(ind int, rem int, change int) (res int) {
		if rem == 0 && change == 1 {
			if ind == n {
				return 1
			}
			return pw[n-ind-1] * 9 % m
		}
		if ind == n {
			return 0
		}

		if dp[ind][rem][change] != -1 {
			return dp[ind][rem][change]
		}

		defer func() {
			dp[ind][rem][change] = res
		}()

		for i := range 10 {
			newChange := change
			if i != 0 {
				newChange = 1
			}
			res += f(ind+1, (pw2[ind]%k*i+rem)%k, newChange)
			res %= m
		}

		return
	}

	return f(0, 0, 0)
}
