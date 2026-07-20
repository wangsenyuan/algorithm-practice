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

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans[0], ans[1])
	}
}

func drive(reader *bufio.Reader) [][]int {
	var t int
	fmt.Fscan(reader, &t)
	res := make([][]int, t)
	for i := range t {
		var a, b, k int
		fmt.Fscan(reader, &a, &b, &k)
		res[i] = solve(a, b, k)
	}
	return res
}

const mod = 1e9 + 7

func mul(a, b int) int {
	return a * b % mod
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

func div(a, b int) int {
	return mul(a, pow(b, mod-2))
}

func solve(a, b, k int) []int {
	n := k*(a-1) + 1

	m := mul(b-1, k)

	// C(n, a)
	f1 := 1
	f2 := 1
	if a > n-a {
		a = n - a
	}
	for i := range a {
		f1 = mul(f1, (n-i)%mod)
		f2 = mul(f2, (i+1)%mod)
	}

	m = mul(m, div(f1, f2))

	m++
	if m >= mod {
		m -= mod
	}
	return []int{n % mod, m}
}
