package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
}

func process(reader *bufio.Reader) int {
	var n, l, r int
	fmt.Fscan(reader, &n, &l, &r)
	a := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, l, r, a)
}

const inf = 1 << 60

func solve(n int, l int, r int, a []int) int {
	dp := make([]int, n+1)

	que := make([]int, n+1)
	var head, tail int

	for i := n; i >= 0; i-- {
		for tail < head && que[tail] > i+r {
			tail++
		}
		if tail == head {
			dp[i] = a[i]
		} else {
			dp[i] = a[i] + dp[que[tail]]
			if i+r > n {
				dp[i] = max(dp[i], a[i])
			}
		}
		j := i + l - 1
		if j <= n {
			for head > tail && dp[que[head-1]] <= dp[j] {
				head--
			}
			que[head] = j
			head++
		}
	}

	return dp[0]
}
