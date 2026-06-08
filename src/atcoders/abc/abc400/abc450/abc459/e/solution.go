package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &p[i])
	}
	C := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &C[i])
	}

	D := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &D[i])
	}

	return solve(n, p, C, D)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return (a % mod) * (b % mod) % mod
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

func solve(n int, P []int, C []int, D []int) int {

	adj := make([][]int, n)
	for i := 1; i < n; i++ {
		adj[P[i-1]-1] = append(adj[P[i-1]-1], i)
	}

	calc := func(w int, d int) int {
		// nCr(w + d, d)
		res := 1
		div := 1
		for i := 1; i <= d; i++ {
			res = mul(res, w+i)
			div = mul(div, i)
		}
		res = mul(res, inverse(div))
		return res
	}

	var dfs func(u int) int
	dfs = func(u int) int {
		d := D[u]
		res := 1
		for _, v := range adj[u] {
			res = mul(res, dfs(v))
			if res == 0 {
				return 0
			}
			D[u] += D[v]
			C[u] += C[v]
		}
		if C[u] < D[u] {
			return 0
		}
		return mul(res, calc(C[u]-D[u], d))
	}

	return dfs(0)
}
