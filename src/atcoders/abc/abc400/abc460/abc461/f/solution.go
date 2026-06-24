package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
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

const N = 15

var F [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(F[i-1], i)
	}
}

func solve(n int) int {
	if n == 1 {
		return 1
	}
	var factors []int
	for i := 1; i <= n/i; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if i*i != n {
				factors = append(factors, n/i)
			}
		}
	}

	sort.Ints(factors)

	m := len(factors)
	dp := make([][]int, N)
	dp1 := make([][]int, N)
	ndp := make([][]int, N)
	ndp1 := make([][]int, N)
	for i := range N {
		dp[i] = make([]int, m+1)
		dp1[i] = make([]int, m+1)
		ndp[i] = make([]int, m+1)
		ndp1[i] = make([]int, m+1)
	}

	dp[0][0] = 1
	dp1[0][0] = 0

	for _, a := range factors {
		for j := range N {
			for k := range m {
				// not take a
				ndp[j][k] = add(ndp[j][k], dp[j][k])
				ndp1[j][k] = add(ndp1[j][k], dp1[j][k])
				// take a
				if j > 0 && factors[k]%a == 0 {
					k1 := sort.SearchInts(factors, factors[k]/a)
					if k1 < m && factors[k1] == factors[k]/a {
						ndp[j][k] = add(ndp[j][k], dp[j-1][k1])
						ndp1[j][k] = add(ndp1[j][k], add(dp1[j-1][k1], mul(dp[j-1][k1], a)))
					}
				}
			}
		}
		for j := range N {
			for k := range m {
				dp[j][k] = ndp[j][k]
				dp1[j][k] = ndp1[j][k]
				ndp[j][k] = 0
				ndp1[j][k] = 0
			}
		}
	}

	var res int
	for j := 1; j < N; j++ {
		res = add(res, mul(dp1[j][m-1], F[j]))
	}

	return res
}
