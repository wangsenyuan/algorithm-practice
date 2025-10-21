package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.10f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	segs := make([][]int, n)
	for i := 0; i < n; i++ {
		segs[i] = make([]int, 2)
		fmt.Fscan(reader, &segs[i][0], &segs[i][1])
	}
	return solve(n, segs)
}

func solve(n int, segs [][]int) float64 {
	calc := func(x int) float64 {
		dp := make([][]float64, 2)
		ndp := make([][]float64, 2)
		for i := range 2 {
			dp[i] = make([]float64, 3)
			ndp[i] = make([]float64, 3)
		}
		dp[0][0] = 1.0

		for _, cur := range segs {
			l, r := cur[0], cur[1]
			m := r - l + 1
			for i := range 2 {
				for j := range 3 {
					if i == 0 && x < r {
						// 可以使用一个更大的数
						ndp[i+1][j] += dp[i][j] * float64(min(m, r-x)) / float64(m)
					}
					// 如果bid了x
					if x >= l && x <= r {
						ndp[i][min(j+1, 2)] += dp[i][j] / float64(m)
					}
					// 如果bid了 < x
					if x > l {
						ndp[i][j] += dp[i][j] * float64(min(m, x-l)) / float64(m)
					}
				}
			}
			for i := range 2 {
				for j := range 3 {
					dp[i][j] = ndp[i][j]
					ndp[i][j] = 0
				}
			}
		}

		res := dp[0][2] + dp[1][2] + dp[1][1]

		return res * float64(x)
	}

	var res float64
	for x := 1; x <= 10000; x++ {
		res += calc(x)
	}
	return res
}
