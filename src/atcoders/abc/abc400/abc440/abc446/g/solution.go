package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
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

func solve(a []int) int {
	n := len(a)
	// dp[i] = 以i为block end的方案数
	dp := make([]int, n+1)
	fp := make([]int, n+1)
	pos := make([][]int, n+1)

	dp[0] = 1
	fp[0] = 1
	for i := 1; i <= n; i++ {
		v := a[i-1]
		pos[v] = append(pos[v], i)
		if len(pos[v]) >= v {
			r := pos[v][len(pos[v])-v]
			dp[i] = fp[r-1]
			if len(pos[v]) > v {
				l := pos[v][len(pos[v])-v-1]
				dp[i] = sub(dp[i], fp[l])
			}
		}
		fp[i] = add(fp[i-1], dp[i])
	}

	var res int
	for i := range n {
		res = add(res, dp[i+1])
	}

	return res
}
