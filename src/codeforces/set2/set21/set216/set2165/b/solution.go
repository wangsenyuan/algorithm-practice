package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
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
	return a * b % mod
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	slices.Sort(a)
	var mx int
	n := len(a)
	var freq []int
	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		mx = max(mx, i-j)
		freq = append(freq, i-j)
	}
	// dp[s] = sum of prod(cnt[x]) over supports T with sum(cnt)=s (0/1 knapsack on distinct values).
	dp := make([]int, n+1)
	dp[0] = 1
	for _, c := range freq {
		for s := n - c; s >= 0; s-- {
			if dp[s] != 0 {
				dp[s+c] = add(dp[s+c], mul(dp[s], c))
			}
		}
	}

	var res int
	for i := mx; i <= n; i++ {
		res = add(res, dp[i])
	}
	return res
}
