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

	res := drive(reader)
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

const X = 2000010

var lpf [X]int

var fs [X][]int

func init() {
	var primes []int
	for i := 2; i < X; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= X {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	for i := 2; i < X; i++ {
		u := i
		for u > 1 {
			j := lpf[u]
			fs[i] = append(fs[i], j)
			for u%j == 0 {
				u /= j
			}
		}
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		var t int
		fmt.Fscan(reader, &t)
		queries[i] = make([]int, t+1)
		queries[i][0] = t
		for j := 1; j <= t; j++ {
			fmt.Fscan(reader, &queries[i][j])
		}
	}
	return solve(a, edges, queries)
}

func solve(a []int, edges [][]int, queries [][]int) []int {
	n := len(a)
	adj := make([][]int, n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	fa := make([]int, n)
	dep := make([]int, n)

	var dfs func(p int, u int)

	stacks := make([][]int, X)

	update := func(u int) {
		for _, f := range fs[a[u]] {
			if len(stacks[f]) > 0 {
				v := stacks[f][len(stacks[f])-1]
				if fa[u] < 0 || dep[fa[u]] < dep[v] {
					fa[u] = v
				}
			}
			stacks[f] = append(stacks[f], u)
		}
	}

	revert := func(u int) {
		for _, f := range fs[a[u]] {
			stacks[f] = stacks[f][:len(stacks[f])-1]
		}
	}

	dfs = func(p int, u int) {
		fa[u] = -1

		update(u)

		for _, v := range adj[u] {
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}

		revert(u)
	}

	dfs(-1, 0)

	var ans []int

	for _, q := range queries {
		if q[0] == 1 {
			v := q[1] - 1
			if fa[v] >= 0 {
				ans = append(ans, fa[v]+1)
			} else {
				ans = append(ans, -1)
			}
		} else {
			v, w := q[1]-1, q[2]
			a[v] = w
			dfs(-1, 0)
		}
	}

	return ans
}
