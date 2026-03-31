package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for tc > 0 {
		tc--
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(m, a)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(m int, a []int) int {
	n := len(a)
	if a[0] == 0 {
		a[0] = 1
	}
	if a[0] > 1 {
		return 0
	}
	// dp[x]到i为止，且a[i] = x时的方案数

	sets := make([][]int, m+1)
	for w := 1; w <= m; w++ {
		for v := w + 1; v <= min(m, 2*w); v++ {
			diff := v - w
			if gcd(v, w) == diff {
				sets[v] = append(sets[v], w)
			}
		}
	}

	dp := make([]int, m+1)
	dp[1] = 1
	for i := 1; i < n; i++ {
		v := a[i]
		if v > 0 {
			var sum int
			for _, w := range sets[v] {
				sum = add(sum, dp[w])
			}
			clear(dp)
			dp[v] = sum
		} else {
			for v := m; v > 1; v-- {
				var sum int
				for _, w := range sets[v] {
					sum = add(sum, dp[w])
				}
				dp[v] = sum
			}
		}
		dp[1] = 0
	}
	var res int
	for _, v := range dp {
		res = add(res, v)
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
