package p3855

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

func add(a, b int) int {
	return (a + b) % mod
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func div(a int, b int) int {
	return mul(a, pow(b, mod-2))
}

func sumOfNumbers(l int, r int, k int) int {
	// k 很大
	// 除去当前位， 还有有 pow(r - l + 1, k-1) 个数字
	t := pow(r-l+1, k-1)
	s := (r + l) * (r - l + 1) / 2
	// pow(10, 0) + pow(10, 1) + ... pow(10, k - 1)
	// pow()
	w := sub(pow(10, k), 1)
	res := div(w, 9)

	return mul(t, mul(res, s))
}
