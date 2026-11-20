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
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	mk := 1
	sum := 1
	for sum <= n {
		sum += mk
		mk++
	}
	// k+1 也要能当下标
	mk++

	pref := make([]int32, n+1)
	for i, v := range a {
		pref[i+1] = pref[i] + int32(v)
	}

	// dp[l][d+mk][k&1]：左边已取 l 张，Igor 比 Zhenya 多取 d 张
	// 上一手拿 k 张时，下一手是 k 或 k+1，所以只留这两层，k 从大到小填
	dp := make([][][2]int32, n+1)
	for l := range dp {
		dp[l] = make([][2]int32, 2*mk)
	}

	for k := mk - 1; k >= 1; k-- {
		cur := k & 1
		for l := range dp {
			for j := range dp[l] {
				dp[l][j][cur] = 0
			}
		}
		for rem := k; rem <= n; rem++ {
			for l := 0; l <= n-rem; l++ {
				d := rem - n + 2*l
				if d <= -mk || d >= mk {
					continue
				}
				r := l + rem
				if d <= 0 {
					got := pref[l+k] - pref[l]
					res := got - dp[l+k][d+k+mk][cur]
					if k+1 < mk && l+k+1 <= r {
						got += int32(a[l+k])
						res = max(res, got-dp[l+k+1][d+k+1+mk][(k+1)&1])
					}
					dp[l][d+mk][cur] = res
					continue
				}
				got := pref[r] - pref[r-k]
				res := got - dp[l][d-k+mk][cur]
				if k+1 < mk && r-k-1 >= l {
					got += int32(a[r-k-1])
					res = max(res, got-dp[l][d-k-1+mk][(k+1)&1])
				}
				dp[l][d+mk][cur] = res
			}
		}
	}

	return int(dp[0][mk][1])
}
