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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	// m := len(edges)
	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
	}
	pair := make([]int, n+1)
	marked := make([]int, n+1)

	check := func(c int) int {
		clear(pair)
		clear(marked)

		var id int

		var argument func(u int) bool
		argument = func(u int) bool {
			if marked[u] == id {
				return false
			}
			marked[u] = id
			for _, v := range adj[u] {
				if v == c {
					continue
				}
				if pair[v] == 0 || argument(pair[v]) {
					pair[v] = u
					return true
				}
			}
			return false
		}

		cnt := []int{0, 0}
		for _, e := range edges {
			if e[0] == c || e[1] == c {
				cnt[0]++
			} else {
				cnt[1]++
			}
		}

		for i := 1; i <= n; i++ {
			if i == c {
				continue
			}
			id++
			argument(i)
		}

		var sz int
		for i := 1; i <= n; i++ {
			if pair[i] > 0 {
				sz++
			}
		}

		return 2*(n-1) + 1 - cnt[0] + cnt[1] - sz + (n - 1) - sz
	}

	res := 1 << 60
	for i := 1; i <= n; i++ {
		res = min(res, check(i))
	}

	return res
}
