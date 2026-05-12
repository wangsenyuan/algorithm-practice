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
		res := drive(reader)
		if res {
			fmt.Fprintln(writer, "Yes")
		} else {
			fmt.Fprintln(writer, "No")
		}
	}
}

func drive(reader *bufio.Reader) bool {
	var n, k, v int
	fmt.Fscan(reader, &n, &k, &v)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		edges[i] = []int{a, b}
	}
	return solve(n, edges, k, v)
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 30

func solve(n int, edges [][]int, k int, root int) bool {
	root--
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		first, second := inf, inf
		for _, v := range adj[u] {
			if p != v {
				tmp := dfs(u, v)
				if tmp <= first {
					second = first
					first = tmp
				} else if tmp <= second {
					second = tmp
				}
			}
		}
		if first == inf {
			return 0
		}
		if second == inf {
			return first + 1
		}
		if first+second < k {
			return 0
		}
		return first + 1
	}

	return dfs(-1, root) == 0
}
