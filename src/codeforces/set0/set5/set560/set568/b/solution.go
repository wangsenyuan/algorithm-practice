package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solve(n))
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

func solve(n int) int {
	if n == 1 {
		return 1
	}

	C := make([][]int, n+1)
	for i := range n + 1 {
		C[i] = make([]int, i+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = add(C[i-1][j-1], C[i-1][j])
		}
	}

	B := make([]int, n+1)
	B[0] = 1
	for i := 1; i <= n; i++ {
		for k := range i {
			B[i] = add(B[i], mul(C[i-1][k], B[k]))
		}
	}

	var res int

	for j := 1; j <= n; j++ {
		cur := mul(C[n][j], B[n-j])
		res = add(res, cur)
	}

	return res
}
