package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

const mod = 998244353

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

func solve(n int, edges [][]int) int {
	adj := make([][]int, n)
	for i := range n {
		adj[i] = make([]int, n)
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u][v]++
		adj[v][u]++
	}

	dp := make([][]int, 1<<n)
	for i := range 1 << n {
		dp[i] = make([]int, n)
	}
	for i := range n {
		dp[1<<i][i] = 1
	}

	var res int

	for mask := range 1 << n {
		st := bits.TrailingZeros(uint(mask))
		for v, fv := range dp[mask] {
			if fv > 0 {
				res = add(res, mul(fv, adj[v][st]))
				for w := st + 1; w < n; w++ {
					if (mask>>w)&1 == 0 {
						dp[mask|(1<<w)][w] = add(dp[mask|(1<<w)][w], mul(fv, adj[v][w]))
					}
				}
			}
		}
	}

	res = sub(res, len(edges))

	return mul(res, inverse(2))
}
