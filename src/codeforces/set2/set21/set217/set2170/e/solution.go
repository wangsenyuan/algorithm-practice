package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
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

func inverse(a int) int {
	return pow(a, mod-2)
}

func div(a, b int) int {
	return mul(a, inverse(b))
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	constraints := make([][]int, m)
	for i := range m {
		constraints[i] = make([]int, 2)
		fmt.Fscan(reader, &constraints[i][0], &constraints[i][1])
	}
	return solve(n, constraints)
}

func solve(n int, constraints [][]int) int {
	L := make([]int, n)
	for i := range n {
		L[i] = -1
	}
	for _, cur := range constraints {
		l, r := cur[0]-1, cur[1]-1
		L[r] = max(L[r], l)
	}
	for i := 1; i < n; i++ {
		L[i] = max(L[i], L[i-1])
	}

	// dp[i] = s[i-1] != s[i]的方案数
	// fp[i] = sum(dp[0...i])
	dp := 2
	fp := make([]int, n+2)
	fp[1] = 2
	for i := 1; i <= n; i++ {
		dp = sub(fp[i], fp[L[i-1]+1])
		fp[i+1] = add(fp[i], dp)
	}

	return dp
}
