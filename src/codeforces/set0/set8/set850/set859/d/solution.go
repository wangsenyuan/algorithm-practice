package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	N := 1 << n
	a := make([][]int, N)
	for i := range N {
		a[i] = make([]int, N)
		for j := range N {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(n, a)
}

func solve(n int, a [][]int) float64 {
	N := len(a)

	p := make([][]float64, n+1)
	for i := range n + 1 {
		p[i] = make([]float64, N)
	}

	for i := range N {
		p[0][i] = 1.0
	}

	for r := 1; r <= n; r++ {
		step := 1 << (r - 1)
		for t := range N {
			var sum float64
			for u := range N {
				if (u/step)^1 == (t / step) {
					sum += p[r-1][u] * float64(a[t][u]) / 100
				}
			}
			p[r][t] = p[r-1][t] * sum
		}
	}

	dp := make([][]float64, n+1)
	for i := range n + 1 {
		dp[i] = make([]float64, N)
	}

	for r := 1; r <= n; r++ {
		// r is round
		step := 1 << (r - 1)
		for t := range N {
			for u := range N {
				if (u/step)^1 == (t / step) {
					tmp := dp[r-1][t] + dp[r-1][u] + p[r][t]*float64(step)
					dp[r][t] = max(dp[r][t], tmp)
				}
			}
		}
	}

	var res float64
	for i := range N {
		res = max(res, dp[n][i])
	}

	return res
}
