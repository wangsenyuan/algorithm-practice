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
	a := make([][]float64, n)
	for i := range n {
		a[i] = make([]float64, n)
		for j := range n {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

func solve(a [][]float64) float64 {
	n := len(a)
	if n == 1 {
		return 1.0
	}
	// dp[mask][i] 表示mask，最后胜利的是i时的概率
	N := 1 << n
	dp := make([][]float64, N)
	for i := range N {
		dp[i] = make([]float64, n)
		for j := range n {
			dp[i][j] = -1
		}
	}

	var f func(mask int, i int) float64

	f = func(mask int, i int) (res float64) {
		if dp[mask][i] > -0.5 {
			return dp[mask][i]
		}
		defer func() {
			dp[mask][i] = res
		}()
		if mask == N-1 {
			if i == 0 {
				return 1
			}
			return 0
		}
		for j := range n {
			if (mask>>j)&1 == 0 {
				cur := f(mask|(1<<j), i)*a[i][j] + f(mask|(1<<j), j)*a[j][i]
				res = max(res, cur)
			}
		}
		return
	}

	var res float64
	for i := range n {
		for j := range i {
			mask := (1 << i) | (1 << j)
			cur := f(mask, i)*a[i][j] + f(mask, j)*a[j][i]
			res = max(res, cur)
		}
	}

	return res
}
