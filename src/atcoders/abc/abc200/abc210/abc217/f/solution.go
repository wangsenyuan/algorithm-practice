package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	pairs := make([][]int, m)
	for i := range m {
		pairs[i] = make([]int, 2)
		fmt.Fscan(reader, &pairs[i][0], &pairs[i][1])
	}
	return solve(n, pairs)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int, pairs [][]int) int {

	C := make([][]int, 2*n+1)
	for i := range 2*n + 1 {
		C[i] = make([]int, 2*n+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = add(C[i-1][j-1], C[i-1][j])
		}
	}

	can := make([][]bool, 2*n)
	for i := range 2 * n {
		can[i] = make([]bool, 2*n)
	}
	for _, cur := range pairs {
		u, v := cur[0]-1, cur[1]-1
		if u > v {
			u, v = v, u
		}
		can[u][v] = true
	}

	dp := make([][]int, 2*n)
	for i := range 2 * n {
		dp[i] = make([]int, 2*n)
		for j := range 2 * n {
			dp[i][j] = -1
		}
	}

	var f func(start int, end int) int
	f = func(start int, end int) (res int) {
		if start > end {
			return 1
		}
		if dp[start][end] != -1 {
			return dp[start][end]
		}
		defer func() {
			dp[start][end] = res
		}()

		for i := start + 1; i <= end; i += 2 {
			if can[start][i] {
				x := f(start+1, i-1)
				y := f(i+1, end)
				cnt1 := (i - start + 1) / 2
				cnt2 := (end - i) / 2
				tmp := mul(x, y)
				tmp = mul(tmp, C[cnt1+cnt2][cnt1])
				res = add(res, tmp)
			}
		}

		return
	}

	return f(0, 2*n-1)
}
