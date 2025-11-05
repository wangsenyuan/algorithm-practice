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

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	return solve(n, m)
}

const mod = 1e9 + 7

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

func solve(n int, m int) int {
	res := mul(n%mod, m%mod)

	m = min(m, n)

	sum := func(l int, r int) int {
		res := mul((l+r)%mod, (r-l+1)%mod)
		return mul(res, (mod+1)/2)
	}

	var ans int
	min_val := m
	for i := 1; i*i <= n; i++ {
		lf := n / (i + 1)
		rg := min(m, n/i)
		if lf >= rg {
			continue
		}
		min_val = lf
		ans = add(ans, mul(i, sum(lf+1, rg)))
	}

	for i := 1; i <= min_val; i++ {
		ans = add(ans, mul((n/i)%mod, i))
	}

	return sub(res, ans)
}
