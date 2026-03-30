package p3881

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

func countVisiblePeople(n int, pos int, k int) int {
	F := make([]int, n)
	F[0] = 1
	for i := 1; i < n; i++ {
		F[i] = mul(F[i-1], i)
	}
	I := make([]int, n)
	I[n-1] = pow(F[n-1], mod-2)
	for i := n - 2; i >= 0; i-- {
		I[i] = mul(I[i+1], i+1)
	}
	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return mul(F[n], mul(I[r], I[n-r]))
	}

	return mul(nCr(n-1, k), 2)
}
