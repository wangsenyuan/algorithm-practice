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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]int, t)
	for i := range t {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		a := make([][]int, n)
		for r := range n {
			a[r] = make([]int, m)
			for c := range m {
				fmt.Fscan(reader, &a[r][c])
			}
		}
		res[i] = solve(a)
	}
	return res
}

const inf = 1 << 60

func solve(a [][]int) int {
	n := len(a)
	m := len(a[0])
	dp := make([][]int, n)
	fp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
		fp[i] = make([]int, m)
		for j := range m {
			dp[i][j] = -inf
			fp[i][j] = -inf
		}
	}
	dp[0][0] = a[0][0]
	for i := range n {
		for j := range m {
			if i > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j]+a[i][j])
			}
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1]+a[i][j])
			}
		}
	}

	fp[n-1][m-1] = a[n-1][m-1]

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i < n-1 {
				fp[i][j] = max(fp[i][j], fp[i+1][j]+a[i][j])
			}
			if j < m-1 {
				fp[i][j] = max(fp[i][j], fp[i][j+1]+a[i][j])
			}
		}
	}

	type data struct {
		id  int
		val int
	}
	// 如果在(i, j)处进行翻转, 那么必须知道不通过(i, j) 的最大值
	// best[i][0], best[i][1] 表示第i行, 最好的两列的通过值
	// 要用斜线的部分
	best := make([][]data, n+m)

	for i := range n {
		for j := range m {
			w := dp[i][j] + fp[i][j] - a[i][j]
			cur := data{j, w}
			for k := range len(best[i+j]) {
				if best[i+j][k].val <= w {
					best[i+j][k], cur = cur, best[i+j][k]
				}
			}
			if len(best[i+j]) < 2 {
				best[i+j] = append(best[i+j], cur)
			}
		}
	}

	res := inf

	for i := range n {
		for j := range m {
			// 如果翻转(i, j)
			var sum int
			if i > 0 && j > 0 {
				sum = max(dp[i-1][j], dp[i][j-1])
			} else if i > 0 {
				sum = dp[i-1][j]
			} else if j > 0 {
				sum = dp[i][j-1]
			}

			if i+1 < n && j+1 < m {
				sum += max(fp[i+1][j], fp[i][j+1])
			} else if i+1 < n {
				sum += fp[i+1][j]
			} else if j+1 < m {
				sum += fp[i][j+1]
			}
			sum -= a[i][j]

			if i+j != 0 && i+j != n+m-2 {
				if best[i+j][0].id != j {
					sum = max(sum, best[i+j][0].val)
				} else if len(best[i+j]) > 1 {
					sum = max(sum, best[i+j][1].val)
				}
			}

			res = min(res, sum)
		}
	}

	return res
}
