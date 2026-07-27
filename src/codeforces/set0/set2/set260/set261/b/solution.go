package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%.10f\n", drive(reader))
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var p int
	fmt.Fscan(reader, &p)
	return solve(a, p)
}

func solve(a []int, p int) float64 {
	n := len(a)

	// dp[j][s] = # of j-subsets with sum s
	dp := make([][]int64, n+1)
	for j := range dp {
		dp[j] = make([]int64, p+1)
	}
	dp[0][0] = 1

	for _, v := range a {
		for j := n - 1; j >= 0; j-- {
			for s := 0; s+v <= p; s++ {
				if dp[j][s] != 0 {
					dp[j+1][s+v] += dp[j][s]
				}
			}
		}
	}

	// E[X] = Σ_k Pr(X >= k) = Σ_k (# k-subsets with sum ≤ p) / C(n, k)
	var ans float64
	for k := 1; k <= n; k++ {
		var ways int64
		for s := 0; s <= p; s++ {
			ways += dp[k][s]
		}
		ans += float64(ways) / comb(n, k)
	}
	return ans
}

func comb(n, k int) float64 {
	if k < 0 || k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	res := 1.0
	for i := 1; i <= k; i++ {
		res = res * float64(n-k+i) / float64(i)
	}
	return res
}
