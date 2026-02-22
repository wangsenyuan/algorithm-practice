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
	res := drive(reader)
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		if u != v {
			adj[u] = append(adj[u], v)
		}
	}

	for u := range n {
		slices.Sort(adj[u])
		adj[u] = slices.Compact(adj[u])
	}

	vis := make([]bool, n)
	can := make([]bool, n)
	var cnt, cnt2 int

	var dfs func(u int, mx int)
	dfs = func(u int, mx int) {
		vis[u] = true
		cnt2++
		for _, v := range adj[u] {
			if !vis[v] && v <= mx {
				dfs(v, mx)
			}
			if !can[v] {
				can[v] = true
				cnt++
			}
		}
	}

	ans := make([]int, n)

	for i := range n {
		if i == 0 {
			can[i] = true
			cnt++
		}
		if can[i] {
			dfs(i, i)
		}
		ans[i] = -1
		if cnt2 == i+1 {
			ans[i] = cnt - cnt2
		}
	}

	return ans
}
