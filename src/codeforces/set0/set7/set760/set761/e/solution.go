package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}

	res := solve(n, edges)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}

	buf.WriteTo(os.Stdout)
}

const inf = 1e18

func solve(n int, edges [][]int) [][]int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	for i := range n {
		if len(adj[i]) > 4 {
			return nil
		}
	}

	ans := make([][]int, n)
	ans[0] = []int{0, 0}

	pair := []int{2, 3, 0, 1}

	dd := []int{-1, 0, 1, 0, -1}

	var dfs func(p int, u int, avoid int, edgeLen int)
	dfs = func(p int, u int, avoid int, edgeLen int) {
		var i int
		for _, v := range adj[u] {
			if p == v {
				continue
			}
			if i == avoid {
				i++
			}
			ans[v] = []int{ans[u][0] + dd[i]*(1<<edgeLen), ans[u][1] + dd[i+1]*(1<<edgeLen)}
			dfs(u, v, pair[i], edgeLen-1)
			i++
		}
	}

	dfs(-1, 0, -1, n)

	return ans
}
