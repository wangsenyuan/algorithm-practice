package main

import (
	"bufio"
	"fmt"
	"os"
)

const inf int = 1 << 60

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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}

	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	dsu := make([]int, n+len(edges))
	for i := range dsu {
		dsu[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if dsu[x] != x {
			dsu[x] = find(dsu[x])
		}
		return dsu[x]
	}

	adj := make([][]int, n+len(edges))
	val := make([]int, n+len(edges))
	var tot int
	// process edges sequentially
	id := n
	deg := make([]int, n)
	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		deg[u]++
		deg[v]++
		tot += w
		ru := find(u)
		rv := find(v)
		val[id] = w
		if ru != rv {
			// Merge two different components: create one new internal node
			// whose children are the current roots of both components.
			adj[id] = append(adj[id], ru)
			adj[id] = append(adj[id], rv)
			dsu[ru] = id
			dsu[rv] = id
		} else {
			// Edge inside one component: create a unary wrapper, not two copies
			// of the same child. This preserves the editorial's "ancestor can
			// lower the transfer cost" behavior.
			adj[id] = append(adj[id], ru)
			dsu[ru] = id
		}
		id++
	}

	var dfs func(u int, x int) int
	dfs = func(u int, x int) int {
		var cnt int
		x = min(x, val[u])
		for _, v := range adj[u] {
			cnt += dfs(v, x)
		}

		if u < n && deg[u]%2 == 1 {
			cnt++
		}
		tot += (cnt / 2) * x
		return cnt & 1
	}

	dfs(id-1, inf)

	return tot
}
