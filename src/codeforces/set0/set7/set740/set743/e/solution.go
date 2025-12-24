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
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const inf = 1 << 60

func solve(a []int) int {
	n := len(a)
	pos := make([][]int, 8)

	for i, v := range a {
		pos[v-1] = append(pos[v-1], i)
	}

	T := 1 << 8

	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, T)
	}

	where := make([]int, 8)

	play := func(x int) int {
		for i := range n + 1 {
			for s := range T {
				dp[i][s] = -inf
			}
		}

		dp[0][0] = 0
		clear(where)

		for i, v := range a {
			v--
			// 如果不使用i
			for s := range T {
				// i+1有可能已经被更新了
				dp[i+1][s] = max(dp[i+1][s], dp[i][s])
			}
			for s := range T {
				if dp[i][s] >= 0 && (s>>v)&1 == 0 {
					if where[v]+x <= len(pos[v]) {
						j := pos[v][where[v]+x-1] + 1
						ns := s | (1 << v)
						dp[j][ns] = max(dp[j][ns], dp[i][s]+x)
					}
					if where[v]+x+1 <= len(pos[v]) {
						j := pos[v][where[v]+x] + 1
						ns := s | (1 << v)
						dp[j][ns] = max(dp[j][ns], dp[i][s]+x+1)
					}
				}
			}
			where[v]++
		}

		return dp[n][T-1]
	}

	var res int

	// 0 次和1次
	for i := range 8 {
		if len(pos[i]) > 0 {
			res++
		}
	}

	// 至少出现1次
	for x := 1; x <= n; x++ {
		tmp := play(x)
		if tmp < 0 {
			break
		}
		res = max(res, tmp)
	}

	return res
}
