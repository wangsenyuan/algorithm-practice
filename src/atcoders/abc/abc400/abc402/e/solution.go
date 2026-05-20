package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.10f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, x int
	fmt.Fscan(reader, &n, &x)
	problems := make([][]int, n)
	for i := range n {
		problems[i] = make([]int, 3)
		fmt.Fscan(reader, &problems[i][0], &problems[i][1], &problems[i][2])
	}
	return solve(x, problems)
}

const inf = 1 << 30

func solve(x int, problems [][]int) float64 {
	n := len(problems)

	dp := make([][]float64, 1<<n)
	for i := range 1 << n {
		dp[i] = make([]float64, x+1)
	}

	T := 1 << n

	for w := range x + 1 {
		for mask := range T {
			for s := mask; s > 0; s = (s - 1) & mask {
				i := bits.TrailingZeros(uint(s))
				s, c := problems[i][0], problems[i][1]
				p := float64(problems[i][2]) / 100
				if c <= w {
					v1 := p * (dp[mask^(1<<i)][w-c] + float64(s))
					v2 := (1 - p) * dp[mask][w-c]
					dp[mask][w] = max(dp[mask][w], v1+v2)
				}
			}
		}
	}

	return dp[T-1][x]
}
