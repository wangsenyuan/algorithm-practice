package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	return solve(n, s)
}

func solve(n int, s string) []int {
	// TODO: solve by hand first.
	adj := make([][]int, n)
	rev := make([][]int, n)

	for i := range n - 1 {
		if s[i] == 'R' {
			adj[i] = append(adj[i], i+1)
			rev[i+1] = append(rev[i+1], i)
		} else {
			adj[i+1] = append(adj[i+1], i)
			rev[i] = append(rev[i], i+1)
		}
	}
	deg := make([]int, n)
	que := make([]int, n)
	// deg[i] <= 2 (最多只有两条边进来, 且这种情况下, 没有出边)
	bfs := func(adj [][]int) []int {
		clear(deg)
		for u := range n {
			// adj[u]最多也有只有两个 (这个时候它就是一个起点)
			for _, v := range adj[u] {
				deg[v]++
			}
		}
		dp := make([]int, n)
		var head, tail int
		for u := range n {
			if deg[u] == 0 {
				que[head] = u
				head++
				dp[u] = 1
			}
		}

		for tail < head {
			u := que[tail]
			tail++
			for _, v := range adj[u] {
				deg[v]--
				dp[v] += dp[u]
				if deg[v] == 0 {
					dp[v]++
					que[head] = v
					head++
				}
			}
		}

		return dp
	}

	// dp[u] 表示左边需要多少空间(包括自己)
	dp := bfs(adj)
	// fp[u]表示右边需要多少空间
	fp := bfs(rev)

	ans := make([]int, n+1)
	for u := range n {
		l := dp[u] - 1
		r := n - fp[u]
		ans[l]++
		ans[r+1]--
	}

	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}

	return ans[:n]
}
