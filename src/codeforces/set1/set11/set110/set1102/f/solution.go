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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(n, m, a)
}

const inf = 1 << 60

func solve(n int, m int, a [][]int) int {

	N := 1 << n

	dp := make([][][]int, N)
	for i := range N {
		dp[i] = make([][]int, n)
		for j := range n {
			dp[i][j] = make([]int, n)
			for k := range n {
				dp[i][j][k] = -inf
			}
		}
	}

	for i := range n {
		dp[1<<i][i][i] = inf
	}

	get := func(r1 int, r2 int) int {
		res := inf
		for j := range m {
			res = min(res, abs(a[r1][j]-a[r2][j]))
		}
		return res
	}

	fp := make([][]int, n)
	for i := range n {
		fp[i] = make([]int, n)
		for j := range n {
			fp[i][j] = get(i, j)
		}
	}

	for s := 1; s < N; s++ {
		for i := range n {
			if (s>>i)&1 == 1 {
				for j := range n {
					if (s>>j)&1 == 1 {
						// dp[s][i][j] =>
						for k := range n {
							if (s>>k)&1 == 0 {
								ns := s | (1 << k)
								dp[ns][i][k] = max(dp[ns][i][k], min(dp[s][i][j], fp[j][k]))
							}
						}
					}
				}
			}
		}
	}

	var res int

	for i := range n {
		for j := range n {
			tmp := dp[N-1][i][j]
			for k := 1; k < m; k++ {
				// 前一列的最后一个数，和当前列的第一个数
				tmp = min(tmp, abs(a[j][k-1]-a[i][k]))
			}
			res = max(res, tmp)
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
