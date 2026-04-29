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
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	p := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &p[i])
	}
	return solve(n, k, p)
}

func solve(n int, k int, p []int) int {
	adj := make([][]int, n)
	for i := 1; i < n; i++ {
		adj[p[i-1]-1] = append(adj[p[i-1]-1], i)
	}

	var cnt []int

	mostDep := n

	var dfs func(u int, d int)
	dfs = func(u int, d int) {
		if len(adj[u]) == 0 {
			mostDep = min(mostDep, d)
		}
		if len(cnt) == d {
			cnt = append(cnt, 0)
		}
		cnt[d]++
		for _, v := range adj[u] {
			dfs(v, d+1)
		}
	}

	dfs(0, 0)

	dp := make([]bool, n+1)
	dp[0] = true

	var sum int
	for i := range mostDep + 1 {
		v := cnt[i]
		sum += v
		for w := sum; w >= v; w-- {
			if dp[w-v] {
				dp[w] = true
			}
		}
	}

	for x := 1; x <= sum; x++ {
		if dp[x] {
			y := sum - x
			if x <= k && y <= n-k || x <= n-k && y <= k {
				return mostDep + 1
			}
		}
	}

	return mostDep
}

const inf = 1 << 30
