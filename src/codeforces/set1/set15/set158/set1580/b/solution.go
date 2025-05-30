package main

import "fmt"

func main() {
	var n, m, k, p int
	fmt.Scanf("%d %d %d %d", &n, &m, &k, &p)
	fmt.Println(solve(n, m, k, p))
}

func solve(n int, m int, k int, p int) int {
	add := func(a, b int) int {
		a %= p
		b %= p
		a += b
		if a >= p {
			a -= p
		}
		return a
	}

	mul := func(num ...int) int {
		res := 1
		for _, x := range num {
			res *= x
			res %= p
		}
		return res
	}

	P := make([]int, n+1)
	P[0] = 1
	for i := 1; i <= n; i++ {
		P[i] = mul(P[i-1], i)
	}

	C := make([][]int, n+1)
	for i := range C {
		C[i] = make([]int, i+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = add(C[i-1][j-1], C[i-1][j])
		}
	}

	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
			for x := range dp[i][j] {
				dp[i][j][x] = -1
			}
		}
	}

	var dfs func(dep int, sz int, need int) int

	dfs = func(dep int, sz int, need int) int {
		if dep < 0 {
			if need > 0 {
				// 高度超过了m，不能再安排新的了
				return 0
			}
			return P[sz]
		}
		if sz == 0 {
			return 1
		}
		if dp[dep][sz][need] != -1 {
			return dp[dep][sz][need]
		}

		p := &dp[dep][sz][need]

		if dep == 0 {
			// 找到了一个
			need--
		}
		var res int

		for lsz := range sz {
			for lneed := max(need-(sz-1-lsz), 0); lneed <= min(lsz, need); lneed++ {
				tmp := dfs(dep-1, lsz, lneed)
				if tmp == 0 {
					continue
				}
				tmp1 := dfs(dep-1, sz-1-lsz, need-lneed)
				res = add(res, mul(tmp, tmp1, C[sz-1][lsz]))
			}
		}
		*p = res
		return res
	}

	return dfs(m-1, n, k)
}
