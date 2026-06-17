package main

import "fmt"

func main() {
	var tc int
	fmt.Scan(&tc)

	for range tc {
		var n, m int
		fmt.Scan(&n, &m)
		res := solve(n, m)
		fmt.Println(res)
	}
}

const mod = 998244353

func add(a, b int) int {
	a %= mod
	b %= mod
	return (a + b) % mod
}

func mul(a, b int) int {
	a %= mod
	b %= mod
	return a * b % mod
}

func solve(n int, m int) int {
	var res int

	n1 := uint(n)

	for d := uint(1); d <= n1; {
		// n < d * 10 取n
		var nd uint
		if n1/10 <= d && n1 < d*10 {
			nd = n1
		} else {
			nd = d*10 - 1
		}

		d1 := uint(d) * 10
		c := gcd(d1-1, uint(m))
		m1 := uint(m) / c
		x := n / int(m1)
		y := int(nd - d + 1)
		res = add(res, mul(x, y))
		d = nd + 1
	}

	return res
}

func gcd(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
