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

func drive(reader *bufio.Reader) string {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	m := n * k
	edges := make([][]int, m-1)
	for i := range m - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, k, edges)
}

func solve(n, k int, edges [][]int) string {
	if k == 1 {
		return "Yes"
	}
	adj := make([][]int, n*k)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		var todo []int
		for _, v := range adj[u] {
			if p != v {
				tmp := dfs(u, v)
				if tmp < 0 {
					return -1
				}
				if tmp > 0 {
					todo = append(todo, tmp)
				}
				if len(todo) > 2 {
					return -1
				}
			}
		}
		if len(todo) == 0 {
			return 1
		}
		if len(todo) == 1 {
			w := todo[0] + 1
			return w % k
		}
		// len(todo) == 2
		l := todo[0]
		r := todo[1]
		if (l+1+r)%k != 0 {
			return -1
		}
		return 0
	}

	res := dfs(-1, 0)

	if res == 0 {
		return "Yes"
	}
	return "No"
}
