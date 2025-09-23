package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

const mod = 1000000007

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

func sub(a, b int) int {
	return add(a, mod-b)
}

func inverse(a int) int {
	return pow(a, mod-2)
}

func solve(A int, B int, n int, x int) int {
	if A == 0 {
		return B
	}

	if B == 0 {
		return mul(pow(A, n), x)
	}
	if A == 1 {
		return add(x, mul(B, n%mod))
	}
	res1 := mul(pow(A, n), x)
	// A > 1
	res2 := mul(sub(pow(A, n), 1), inverse(A-1))
	res2 = mul(res2, B)

	return add(res1, res2)
}

func drive(reader *bufio.Reader) int {
	var A, B, n, x int
	fmt.Fscan(reader, &A, &B, &n, &x)
	return solve(A, B, n, x)
}
