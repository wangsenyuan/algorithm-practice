package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (n int, edges [][]int, conditions [][]int, res []int) {
	var q int
	fmt.Fscan(reader, &n, &q)
	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	conditions = make([][]int, q)
	for i := range q {
		conditions[i] = make([]int, 3)
		fmt.Fscan(reader, &conditions[i][0], &conditions[i][1], &conditions[i][2])
	}
	res = solve(n, edges, conditions)
	return
}

func solve(n int, edges [][]int, conditions [][]int) []int {
	adj := make([][]int, n)
	for i, cur := range conditions {
		u, v := cur[0]-1, cur[1]-1
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}

	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
	}

	dp := make([]int, n)
	for i := range n {
		dp[i] = -1
	}

	var dfs func(u int, w int, cnt *int) bool
	dfs = func(u int, w int, cnt *int) bool {
		if dp[u] >= 0 {
			return dp[u] == w
		}
		if deg[u]%2 == 1 {
			*cnt++
		}
		dp[u] = w
		for _, i := range adj[u] {
			cur := conditions[i]
			v := (cur[1] - 1) ^ (cur[0] - 1) ^ u
			if !dfs(v, cur[2]^w, cnt) {
				return false
			}
		}
		return true
	}

	oddComp := -1

	for i := range n {
		if dp[i] < 0 {
			var cnt int
			if !dfs(i, 0, &cnt) {
				return nil
			}
			if cnt%2 == 1 {
				oddComp = i
			}
		}
	}

	var sum int
	for i := range n {
		if deg[i]%2 == 1 {
			sum ^= dp[i]
		}
	}

	marked := make([]bool, n)
	var dfs2 func(u int, s int)
	dfs2 = func(u int, s int) {
		if marked[u] {
			return
		}
		marked[u] = true
		dp[u] ^= s
		for _, i := range adj[u] {
			cur := conditions[i]
			v := (cur[1] - 1) ^ (cur[0] - 1) ^ u
			dfs2(v, s)
		}
	}

	if oddComp >= 0 {
		dfs2(oddComp, sum)
	}

	res := make([]int, n-1)
	for i, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		res[i] = dp[u] ^ dp[v]
	}

	return res
}
