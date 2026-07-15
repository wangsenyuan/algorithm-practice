package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for _, v := range drive(reader) {
		fmt.Println(v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]int, t)
	for i := range t {
		var n int
		var k int
		fmt.Fscan(reader, &n, &k)
		res[i] = solve(n, k)
	}
	return res
}

const mod = 998244353

func add(a int, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int, k int) int {
	dp := make([][]int, k+1)
	ndp := make([][]int, k+1)
	fp := make([][]int, k+1)
	nfp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, 2)
		ndp[i] = make([]int, 2)
		fp[i] = make([]int, 2)
		nfp[i] = make([]int, 2)
	}
	// dp[i][0/1]表示有i个popcount的情况下, 和n 不相等(0)或者相等时的数量

	var ds []int
	for i := n; i > 0; i >>= 1 {
		ds = append(ds, i&1)
	}

	slices.Reverse(ds)

	dp[0][1] = 1

	pw := make([]int, len(ds)+1)
	pw[0] = 1
	for i := 1; i <= len(ds); i++ {
		pw[i] = add(pw[i-1], pw[i-1])
	}

	m := len(ds)
	for p, v := range ds {
		for i := range p + 1 {
			for e := range 2 {
				// 现在这里再放置一个新的数
				for d := range 2 {
					if e == 1 && d > v {
						continue
					}
					ne := e
					if e == 1 && d < v {
						ne = 0
					}
					ni := i + d
					if ni <= k {
						ndp[ni][ne] = add(ndp[ni][ne], dp[i][e])
						nfp[ni][ne] = add(nfp[ni][ne], fp[i][e])
						if d == 1 {
							nfp[ni][ne] = add(nfp[ni][ne], mul(pw[m-p-1], dp[i][e]))
						}
					}
				}
			}
		}
		dp, ndp = ndp, dp
		fp, nfp = nfp, fp
		// dp[0][0] = add(dp[0][0], 1)
		for i := range k + 1 {
			for e := range 2 {
				ndp[i][e] = 0
				nfp[i][e] = 0
			}
		}
	}

	return add(fp[k][0], fp[k][1])
}
