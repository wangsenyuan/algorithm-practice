package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	res := solve(n, k)
	fmt.Println(res)
}

const mod = 1000000007

func add(a int, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a int, b int) int {
	return a * b % mod
}

func pow(a int, b int) int {
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

func solve(n int, k int) int {
	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(F[i-1], i)
	}
	I := make([]int, n+1)
	I[n] = inverse(F[n])
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(I[i+1], i+1)
	}
	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return mul(F[n], mul(I[r], I[n-r]))
	}

	var res int
	for sz := 1; sz <= k; sz++ {
		// 选择sz-1个点，和1组成一个cycle
		cur := nCr(k-1, sz-1)
		// 排列cycle
		cur = mul(cur, F[sz-1])
	
		if sz < k {
			cur = mul(cur, mul(sz, pow(k, k-sz-1)))
		}
		res = add(res, cur)
	}

	return mul(res, pow(n-k, n-k))
}
