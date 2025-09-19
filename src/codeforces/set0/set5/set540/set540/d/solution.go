package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.9f %.9f %.9f\n", res[0], res[1], res[2])
}

func drive(reader *bufio.Reader) []float64 {
	var r, g, b int
	fmt.Fscan(reader, &r, &g, &b)
	return solve(r, g, b)
}

func solve(r int, g int, b int) []float64 {
	dp := make([][][]float64, r+1)
	for i := range r + 1 {
		dp[i] = make([][]float64, g+1)
		for j := range g + 1 {
			dp[i][j] = make([]float64, b+1)
		}
	}
	dp[r][g][b] = 1
	sum := r + g + b

	for s := sum; s > 1; s-- {
		for i := 0; i <= r && i <= s; i++ {
			for j := 0; j <= g && j <= s-i; j++ {
				k := s - i - j
				if k > b {
					continue
				}
				d1 := i*j + j*k + k*i
				// 从 dp[i][j][k] 转移到下一个状态
				// 一共有s个人， 有两个人2相遇
				// 下面这个转移可不能不大对～
				if i > 0 && j > 0 {
					d2 := i * j
					dp[i][j-1][k] += dp[i][j][k] * float64(d2) / float64(d1)
				}
				if i > 0 && k > 0 {
					d2 := i * k
					dp[i-1][j][k] += dp[i][j][k] * float64(d2) / float64(d1)
				}
				if j > 0 && k > 0 {
					d2 := j * k
					dp[i][j][k-1] += dp[i][j][k] * float64(d2) / float64(d1)
				}
			}
		}
	}
	res := make([]float64, 3)

	for i := 1; i <= r; i++ {
		res[0] += dp[i][0][0]
	}
	for j := 1; j <= g; j++ {
		res[1] += dp[0][j][0]
	}
	for k := 1; k <= b; k++ {
		res[2] += dp[0][0][k]
	}
	return res
}
