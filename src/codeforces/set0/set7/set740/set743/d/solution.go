package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ok, ans := drive(reader)
	if !ok {
		fmt.Println("Impossible")
		return
	}
	fmt.Println(ans)
}

func drive(reader *bufio.Reader) (bool, int) {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, a, edges)
}

const inf = 1 << 60

func solve(n int, a []int, edges [][]int) (bool, int) {

	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	res := -inf

	val := make([]int, n)

	var dfs func(p int, u int) int

	dfs = func(p int, u int) int {
		val[u] = a[u]
		first, second := -inf, -inf

		for _, v := range adj[u] {
			if p != v {
				tmp := dfs(u, v)
				if tmp >= first {
					second = first
					first = tmp
				} else if tmp >= second {
					second = tmp
				}
				if second > -inf {
					res = max(res, first+second)
				}
				val[u] += val[v]
			}
		}
		return max(val[u], first)
	}

	dfs(-1, 0)

	if res > -inf {
		return true, res
	}
	return false, 0
}
