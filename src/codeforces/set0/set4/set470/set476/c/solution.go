package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscan(reader, &a, &b)
	res := solve(a, b)
	fmt.Println(res)
}

const mod = 1_000_000_007

func solve(a, b int) int {
	if b == 1 {
		return 0
	}
	// Nice: x % b != 0 and (x/b) / (x%b) = k in [1,a]
	// Let r = x%b, q = x/b. Then q = k*r, so x = q*b + r = r*(k*b + 1)
	// Sum over r in [1,b-1], k in [1,a] of r*(k*b+1)
	// = sum_r * sum_k(k*b+1) = (b-1)*b/2 * a*(b*(a+1)/2 + 1)
	m := int64(mod)
	inv2 := modPow(2, m-2, m)

	sumR := modMul(modMul(int64(b-1), int64(b), m), inv2, m)
	sumKPart := modMul(modMul(int64(b), int64(a), m), int64(a+1), m)
	sumKPart = modMul(sumKPart, inv2, m)
	sumKPart = (sumKPart + int64(a)) % m

	return int(modMul(sumR, sumKPart, m))
}

func modMul(a, b, m int64) int64 {
	return (a % m) * (b % m) % m
}

func modPow(base, exp, m int64) int64 {
	var res int64 = 1
	base %= m
	for exp > 0 {
		if exp&1 == 1 {
			res = res * base % m
		}
		base = base * base % m
		exp >>= 1
	}
	return res
}