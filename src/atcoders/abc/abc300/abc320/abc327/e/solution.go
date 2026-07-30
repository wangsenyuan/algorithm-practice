package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.15f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(p)
}

func solve(p []int) float64 {
	n := len(p)

	// fp[i] = pref sum 0.9 ^^ (k - i)
	fp := make([]float64, n+1)

	fp[0] = 0
	for k := 1; k <= n; k++ {
		fp[k] = fp[k-1]*0.9 + 1.0
	}

	dp := make([]float64, n+1)
	for i := range dp {
		dp[i] = -1e16
	}

	dp[0] = 0
	for i, v := range p {
		for k := i + 1; k > 0; k-- {
			tmp := (dp[k-1]*fp[k-1])*0.9 + float64(v)
			tmp /= fp[k]
			dp[k] = max(dp[k], tmp)
		}
	}

	var best float64 = -1e16
	for k := 1; k <= n; k++ {
		cur := dp[k] - 1200/math.Sqrt(float64(k))
		best = max(best, cur)
	}

	return best
}
