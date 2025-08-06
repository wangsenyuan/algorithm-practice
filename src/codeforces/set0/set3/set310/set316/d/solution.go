package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

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

func solve(a []int) int {
	n := len(a)

	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(F[i-1], i)
	}

	I := make([]int, n+1)
	I[0] = 1
	I[1] = 1
	for i := 2; i <= n; i++ {
		I[i] = add(I[i-1], mul(I[i-2], i-1))
	}

	var m int
	for _, v := range a {
		if v == 1 {
			m++
		}
	}

	// I[m] * n! / m!
	ans := mul(I[m], F[n])

	ans = mul(ans, pow(F[m], mod-2))

	return ans
}
