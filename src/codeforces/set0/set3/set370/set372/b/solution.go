package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	drive(reader, writer)
}

func drive(reader *bufio.Reader, writer *bufio.Writer) {
	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)
	reader.ReadString('\n')
	a := make([]string, n)
	for i := range n {
		s, _ := reader.ReadString('\n')
		a[i] = strings.TrimSpace(s)
	}
	solve(a, q, reader, writer)
}

func solve(a []string, q int, reader *bufio.Reader, writer *bufio.Writer) {
	n := len(a)
	m := len(a[0])

	sum := make([][]int16, n+1)
	for i := range n + 1 {
		sum[i] = make([]int16, m+1)
	}
	for i := range n {
		for j := range m {
			if a[i][j] == '0' {
				sum[i+1][j+1] = 1
			}
			sum[i+1][j+1] += sum[i+1][j]
			sum[i+1][j+1] += sum[i][j+1]
			sum[i+1][j+1] -= sum[i][j]
		}
	}

	dp := make([][][][]int32, n)
	for i := range n {
		dp[i] = make([][][]int32, m)
		for j := range m {
			dp[i][j] = make([][]int32, n)
			for u := range n {
				dp[i][j][u] = make([]int32, m)
			}
		}
	}

	get := func(r1 int, c1 int, r2 int, c2 int) int {
		return int(sum[r2+1][c2+1] - sum[r2+1][c1] - sum[r1][c2+1] + sum[r1][c1])
	}

	for r1 := range n {
		for c1 := range m {
			for r2 := r1; r2 < n; r2++ {
				for c2 := c1; c2 < m; c2++ {
					if get(r1, c1, r2, c2) == (r2-r1+1)*(c2-c1+1) {
						dp[r1][c1][r2][c2] = 1
					}
					if r2 > r1 {
						dp[r1][c1][r2][c2] += dp[r1][c1][r2-1][c2]
					}
					if c2 > c1 {
						dp[r1][c1][r2][c2] += dp[r1][c1][r2][c2-1]
					}
					if r2 > r1 && c2 > c1 {
						dp[r1][c1][r2][c2] -= dp[r1][c1][r2-1][c2-1]
					}
				}
			}
		}
	}

	fp := make([][][][]int32, n)
	for i := range n {
		fp[i] = make([][][]int32, m)
		for j := range m {
			fp[i][j] = make([][]int32, n)
			for u := range n {
				fp[i][j][u] = make([]int32, m)
			}
		}
	}

	for r2 := n - 1; r2 >= 0; r2-- {
		for c2 := m - 1; c2 >= 0; c2-- {
			for r1 := r2; r1 >= 0; r1-- {
				for c1 := c2; c1 >= 0; c1-- {
					fp[r1][c1][r2][c2] = dp[r1][c1][r2][c2]
					if r1 < r2 {
						fp[r1][c1][r2][c2] += fp[r1+1][c1][r2][c2]
					}
					if c1 < c2 {
						fp[r1][c1][r2][c2] += fp[r1][c1+1][r2][c2]
					}
					if r1 < r2 && c1 < c2 {
						fp[r1][c1][r2][c2] -= fp[r1+1][c1+1][r2][c2]
					}
				}
			}
		}
	}

	for range q {
		var r1, c1, r2, c2 int
		fmt.Fscan(reader, &r1, &c1, &r2, &c2)
		ans := fp[r1-1][c1-1][r2-1][c2-1]
		fmt.Fprintln(writer, ans)
	}
}
