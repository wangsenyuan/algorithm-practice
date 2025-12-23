package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &p[i])
	}
	color := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &color[i])
	}
	return solve(n, p, color)
}

const mod = 1e9 + 7

func add(a, b int) int {
	return (a + b) % mod
}

func mul(a, b int) int {
	return (a * b) % mod
}

func solve(n int, p []int, color []int) int {
	adj := make([][]int, n)
	for i, v := range p {
		adj[v] = append(adj[v], i+1)
	}

	dp := make([][2]int, n)
	var dfs func(u int)
	dfs = func(u int) {
		// dp[u][0] 表示和u关联的都是白色节点
		// dp[u][1] 表示u和一个黑色节点关联
		dp[u][0] = 1
		dp[u][1] = 0
		for _, v := range adj[u] {
			dfs(v)
			// u已经关联了黑色节点, 只能使用v的白色节点
			dp[u][1] = mul(dp[u][1], dp[v][0])
			// 或者目前u关联到白色节点，和v的黑色关联
			dp[u][1] = add(dp[u][1], mul(dp[u][0], dp[v][1]))
			dp[u][0] = mul(dp[u][0], dp[v][0])
		}

		if color[u] == 1 {
			// 如果u是黑色节点，那么它只能使用原来的那些白色方案
			dp[u][1] = dp[u][0]
		} else {
			// u是白色节点，那么需要把dp[u][1]的方案也计算进去
			dp[u][0] = add(dp[u][0], dp[u][1])
		}
	}

	dfs(0)

	return dp[0][1]
}
