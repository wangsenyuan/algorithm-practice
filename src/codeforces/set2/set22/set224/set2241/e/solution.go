package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int64 {
	var tc int
	fmt.Fscan(reader, &tc)
	res := make([]int64, tc)
	for i := range tc {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[j])
		}
		edges := make([][]int, n-1)
		for j := range n - 1 {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			edges[j] = []int{u, v}
		}
		res[i] = solve(a, edges)
	}
	return res
}

func nC3(n int) int {
	return n * (n - 1) * (n - 2) / 6
}

func nC2(n int) int {
	return n * (n - 1) / 2
}

func solve(a []int, edges [][]int) int64 {
	check := func(w int) bool {
		// w * w * w = 是一个平方数
		// 那么w必须也是一个平方数?
		if w == 1 {
			return true
		}
		v := sort.Search(w, func(v int) bool {
			return v*v >= w
		})
		return v*v == w
	}

	n := len(a)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var ans int

	sz := make([]int, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u] = 1

		for _, v := range adj[u] {
			if p != v {
				dfs(u, v)

				sz[u] += sz[v]
			}
		}
		if check(a[u]) {
			var sum, pairs, triples int
			for _, v := range adj[u] {
				s := sz[v]
				if p == v {
					s = n - sz[u]
				}
				triples += s * pairs
				pairs += s * sum
				sum += s
			}

			ans += triples
			ans += pairs
		}
	}

	dfs(-1, 0)

	return int64(ans)
}
