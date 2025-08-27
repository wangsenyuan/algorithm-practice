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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(nums ...int) int {
	res := 1
	for _, num := range nums {
		res = (res * num) % mod
	}
	return res
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func inv(a int) int {
	return pow(a, mod-2)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	segments := make([][]int, n)
	for i := range segments {
		var l, r, p, q int
		fmt.Fscan(reader, &l, &r, &p, &q)
		segments[i] = []int{l, r, p, q}
	}
	return solve(m, segments)
}

func solve(m int, segments [][]int) int {
	dp := make([]int, m+1)
	dp[0] = 1
	at := make([][]int, m+1)
	for i, seg := range segments {
		r := seg[1]
		at[r] = append(at[r], i)
		p, q := seg[2], seg[3]
		dp[0] = mul(dp[0], q-p, inv(q))
	}

	for i := 1; i <= m; i++ {
		for _, j := range at[i] {
			l, p, q := segments[j][0], segments[j][2], segments[j][3]
			dp[i] = add(dp[i], mul(dp[l-1], p, inv(q-p)))
		}
	}

	return dp[m]
}
