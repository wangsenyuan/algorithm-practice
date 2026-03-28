package main

import "fmt"

func main() {
	var x, k int
	fmt.Scan(&x, &k)
	res := solve(x, k)
	fmt.Println(res)
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

func solve(x int, k int) int {
	if x == 0 {
		return 0
	}
	x %= mod
	// x => (4 * x - 1 )/ 2 = 2 * x - 1/2
	res := pow(2, k+1)
	res = mul(res, x)
	res = sub(res, pow(2, k))
	res = add(res, 1)
	return res
}
