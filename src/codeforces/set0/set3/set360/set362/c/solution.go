package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := drive(reader)
	fmt.Println(ans[0], ans[1])
}

func drive(reader *bufio.Reader) (ans []int) {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}
	ans = solve(p)
	return
}

func solve(p []int) []int {
	n := len(p)

	var inv int

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n+1)
	}

	for i := range n {
		var u, v int
		for j := i + 1; j < n; j++ {
			if p[i] > p[j] {
				v++
				inv++
			} else {
				u++
			}
			// 交换i,j的时候，会使的inv减少v
			dp[i][j] -= v
			// 如果交换(i, j+1), 会使的增加u
			dp[i][j+1] += u
		}
		u = 0
		v = 0
		for j := i - 1; j >= 0; j-- {
			if p[j] < p[i] {
				v++
			} else {
				u++
			}
			// 如果将i交换到j的位置， 会增加v的inv
			// 为啥要-1/+1， 这个有点神奇的
			dp[j][i] += v
			if j > 0 {
				dp[j-1][i] -= u
			}
		}
	}

	var best, cnt int
	for i := range n {
		for j := i + 1; j < n; j++ {
			if dp[i][j] < best {
				best = dp[i][j]
				cnt = 1
			} else if dp[i][j] == best {
				cnt++
			}
		}
	}

	return []int{inv + best, cnt}
}
