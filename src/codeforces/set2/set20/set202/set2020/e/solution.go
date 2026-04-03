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
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(a, p)
}

const mod = 1e9 + 7

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

func solve(a []int, p []int) int {
	dp := make([]int, 1024)
	dp[0] = 1

	p10_inv := pow(10000, mod-2)

	ndp := make([]int, 1024)

	for i, v := range a {
		for x := range 1024 {
			y := x ^ v
			ndp[y] = add(ndp[y], mul(dp[x], mul(p[i], p10_inv)))
			ndp[x] = add(ndp[x], mul(dp[x], mul(10000-p[i], p10_inv)))
		}
		copy(dp, ndp)
		clear(ndp)
	}

	var res int
	for i := range 1024 {
		res = add(res, mul(mul(i, i), dp[i]))
	}

	return res
}
