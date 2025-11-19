package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	fmt.Println(solve(n))
}

const mod = 1e9 + 7

func add(a, b int) int {
	return (a + b) % mod
}

func sub(a, b int) int {
	return (a - b + mod) % mod
}
func mul(a, b int) int {
	return (a * b) % mod
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

func inverse(a int) int {
	return pow(a, mod-2)
}

func solve(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 4
	}
	F := make([]int, 2*n+1)
	F[0] = 1
	for i := 1; i <= 2*n; i++ {
		F[i] = mul(F[i-1], i)
	}
	I := make([]int, 2*n+1)
	I[2*n] = inverse(F[2*n])
	for i := 2*n - 1; i >= 0; i-- {
		I[i] = mul(I[i+1], i+1)
	}

	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return mul(F[n], mul(I[r], I[n-r]))
	}

	var res int
	for i := 1; i <= n; i++ {
		cur := nCr(n-1+i, i)
		cur = sub(cur, nCr(n-2+i, i))
		res = add(res, cur)
	}
	// iiii这样结构的有n个
	res = sub(res, n)
	res = mul(res, 2)
	res = add(res, n)
	return res
}
