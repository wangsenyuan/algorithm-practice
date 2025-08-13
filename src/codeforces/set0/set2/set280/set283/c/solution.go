package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n, q, t int
	fmt.Fscan(reader, &n, &q, &t)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, q)
	c := make([]int, q)
	for i := range q {
		fmt.Fscan(reader, &b[i], &c[i])
	}
	return solve(t, a, b, c)
}

const mod = 1_000_000_007

func add(a int, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a int, b int) int {
	return a * b % mod
}

func solve(t int, a []int, b []int, c []int) int {
	// bi > ci
	n := len(a)
	next := make([]int, n)
	prev := make([]int, n)
	for i := range n {
		next[i] = -1
		prev[i] = -1
	}

	for i, u := range b {
		u--
		v := c[i]
		v--
		prev[v] = u
		next[u] = v
	}

	vis := make([]bool, n)

	for u := range n {
		if prev[u] == -1 {
			// a[u] 是这个链条中，出现次数最多的那个
			for u != -1 {
				vis[u] = true
				t -= a[u]
				v := next[u]
				if v == -1 {
					// 最后一个，表示的是，1,2,3....这个新的coins， 可以不用出现
					t += a[u]
				} else {
					a[v] += a[u]
				}
				u = v
			}
			if t < 0 {
				return 0
			}
		}
	}

	for i := range n {
		if !vis[i] {
			// cycle
			return 0
		}
	}

	dp := make([]int, t+1)
	dp[0] = 1
	for u := range n {
		for i := 0; i+a[u] <= t; i++ {
			dp[i+a[u]] = add(dp[i+a[u]], dp[i])
		}
	}

	return dp[t]
}
