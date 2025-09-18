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

const mod = 1e9 + 7

func mul(a, b int) int {
	return a * b % mod
}

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	return solve(n, m, k)
}

func solve(n int, m int, k int) int {
	C := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		C[i] = make([]int, n+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = add(C[i-1][j-1], C[i-1][j])
		}
	}
	k = min(k, n*m-k)
	// 第1列要和第n+1列相同数字相同，但是位置可以不同
	dp := make([]int, k+1)
	dp[0] = 1
	ndp := make([]int, k+1)
	t := m % n
	for c := range n {
		clear(ndp)
		for x := 0; x <= k && x <= n; x++ {
			s1 := C[n][x]
			w := m / n
			if t > 0 && c < t {
				w++
			}
			// 每列有s1种选择，共有w列
			s2 := pow(s1, w)
			for y := 0; x+y <= k; y++ {
				ndp[x+y] = add(ndp[x+y], mul(dp[y], s2))
			}
		}
		copy(dp, ndp)
	}

	return dp[k]
}
