package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	fmt.Println(solve(n))
}

const mod = 1000000007

func add(a, b int) int {
	return (a + b) % mod
}

func mul(a, b int) int {
	return (a * b) % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b /= 2
	}
	return res
}

func solve(n int) int {
	// C(2 * n + 2, n + 1) - 1

	F := make([]int, 2*n+3)
	F[0] = 1
	for i := 1; i <= 2*n+2; i++ {
		F[i] = mul(F[i-1], i)
	}

	I := make([]int, 2*n+3)
	I[2*n+2] = pow(F[2*n+2], mod-2)
	for i := 2*n + 1; i >= 0; i-- {
		I[i] = mul(I[i+1], i+1)
	}

	res := F[2*n+2]

	res = mul(res, mul(I[n+1], I[n+1]))

	return add(res, mod-1)
}
