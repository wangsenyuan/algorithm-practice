package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	res := solve(s)
	fmt.Println(res)
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

func solve(s string) int {
	n := len(s)
	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}
	I := make([]int, n+1)
	I[n] = pow(F[n], mod-2)
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return mul(F[n], mul(I[r], I[n-r]))
	}

	suf := make([]int, n+1)
	play := func(x int) int {
		for i := n - 1; i >= 0; i-- {
			suf[i] = suf[i+1]
			if int(s[i]-'0') == x {
				suf[i]++
			}
		}
		var pref int
		var ans int
		for i := range n {
			if int(s[i]-'0') == x-1 {
				pref++
				cur := nCr(pref-1+suf[i], pref)
				ans = add(ans, cur)
			}

		}
		return ans
	}

	var res int
	for i := 1; i <= 9; i++ {
		res = add(res, play(i))
	}
	return res
}
