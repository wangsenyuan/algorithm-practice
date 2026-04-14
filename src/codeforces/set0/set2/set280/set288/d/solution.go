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
	fmt.Fprintln(writer, drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)

	adj := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	sz := make([]int64, n+1)
	par := make([]int, n+1)
	order := make([]int, 0, n)

	// BFS to get parent and topological order
	visited := make([]bool, n+1)
	visited[1] = true
	queue := []int{1}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)
		for _, u := range adj[v] {
			if !visited[u] {
				visited[u] = true
				par[u] = v
				queue = append(queue, u)
			}
		}
	}

	// Compute subtree sizes bottom-up
	for i := 1; i <= n; i++ {
		sz[i] = 1
	}
	for i := len(order) - 1; i >= 0; i-- {
		v := order[i]
		if par[v] != 0 {
			sz[par[v]] += sz[v]
		}
	}

	// f(x) = C(x, 2) = x*(x-1)/2
	f := func(x int64) int64 {
		if x < 2 {
			return 0
		}
		return x * (x - 1) / 2
	}

	// Sv[v] = sum of f(sz[c]) over children c of v
	Sv := make([]int64, n+1)
	for i := len(order) - 1; i >= 0; i-- {
		v := order[i]
		if par[v] != 0 {
			Sv[par[v]] += f(sz[v])
		}
	}

	N := int64(n)
	var ans int64

	for _, v := range order {
		p := par[v]
		// Sum1: paths with LCA = v contribute f(N-sz[v]) per path
		ans += f(N-sz[v]) * (f(sz[v]) - Sv[v])

		// Sum2: for each child c of v, paths through v but not into c's subtree
		for _, u := range adj[v] {
			if u == p {
				continue
			}
			// u is a child of v
			ans += f(sz[u]) * (f(N-sz[u]) + f(sz[u]) - Sv[v] - f(N-sz[v]))
		}
	}

	return ans
}
