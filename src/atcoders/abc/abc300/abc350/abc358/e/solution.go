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
	var k int
	fmt.Fscan(reader, &k)
	c := make([]int, 26)
	for i := range c {
		fmt.Fscan(reader, &c[i])
	}
	return solve(k, c)
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

const X = 1010

var F [X]int
var I [X]int

func init() {
	F[0] = 1
	for i := 1; i < X; i++ {
		F[i] = mul(F[i-1], i)
	}
	I[X-1] = pow(F[X-1], mod-2)
	for i := X - 2; i >= 0; i-- {
		I[i] = mul(I[i+1], i+1)
	}
}

func nCr(n int, r int) int {
	if n < r || r < 0 {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(k int, c []int) int {
	dp := make([]int, k+1)
	fp := make([]int, k+1)
	dp[0] = 1

	for _, x := range c {
		for d, v := range dp {
			if v == 0 {
				continue
			}
			for i := 0; i <= x && i+d <= k; i++ {
				fp[i+d] = add(fp[i+d], mul(v, nCr(i+d, i)))
			}
		}
		copy(dp, fp)
		clear(fp)
	}

	var ans int
	for d := 1; d <= k; d++ {
		ans = add(ans, dp[d])
	}

	return ans
}
