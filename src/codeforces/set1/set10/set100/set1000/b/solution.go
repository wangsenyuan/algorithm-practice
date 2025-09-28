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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(m, a)
}

// const inf = 1 << 60
// 0 1 2 3 4

func solve(m int, a []int) int {
	n := len(a)

	// dp[i]表示从i开始是on的状态
	dp := make([]int, n+2)

	a = append(a, m)
	for i := n - 1; i >= 0; i-- {
		// 交替和
		dp[i] = dp[i+2] + a[i+1] - a[i]
	}

	best := a[0] + dp[1]
	if a[0] != 1 {
		// 在位置1处插入一个指令
		best = max(best, dp[0]+a[0]-1)
	}
	sum := a[0]
	for i := 0; i < n; i++ {
		// 在为止a[i] + 1 处插入一个指令
		if i == n-1 || a[i]+1 < a[i+1] {
			j := m
			if i+1 < n {
				j = a[i+1]
			}
			best = max(best, sum+dp[i+2]+j-a[i]-1)
		}
		if i&1 == 1 && i+1 < n {
			sum += a[i+1] - a[i]
		}
	}

	return best
}
