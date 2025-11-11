package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var p, k int
	fmt.Fscan(reader, &p, &k)
	fmt.Println(solve(p, k))
}

const mod = 1e9 + 7

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

func solve(p int, k int) int {
	if k == 0 {
		return pow(p, p-1)
	}
	if k == 1 {
		return pow(p, p)
	}
	// k >= 2
	m := 1
	x := k % p
	for x != 1 {
		x = x * k % p
		m++
	}
	return pow(p, (p-1)/m)
}
