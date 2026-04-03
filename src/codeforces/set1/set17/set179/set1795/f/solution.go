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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	var k int
	fmt.Fscan(reader, &k)
	a := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, edges, a)
}

func solve(n int, edges [][]int, a []int) int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	return solve1(n, adj, a)
}

func solve1(n int, adj [][]int, a []int) int {
	pos := make([]int, n)
	for i := range n {
		pos[i] = -1
	}
	for i, v := range a {
		pos[v-1] = i
	}

	k := len(a)
	root := a[0] - 1
	parent := make([]int, n)
	order := make([]int, 0, n)
	for i := range n {
		parent[i] = -2
	}
	parent[root] = -1
	stack := []int{root}
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, u)
		for _, v := range adj[u] {
			if parent[v] < -1 {
				parent[v] = u
				stack = append(stack, v)
			}
		}
	}

	dp := make([]int, n)

	check := func(mid int) bool {
		for i := len(order) - 1; i >= 0; i-- {
			u := order[i]
			var cnt int
			var need int
			var room int

			for _, v := range adj[u] {
				if v == parent[u] {
					continue
				}
				tmp := dp[v]
				if tmp > 0 {
					need = max(need, tmp)
					cnt++
					if cnt == 2 || pos[u] >= 0 {
						return true
					}
				} else if tmp < 0 {
					room = min(room, tmp)
				}
			}

			if pos[u] >= 0 {
				r := mid % k
				if pos[u] < r {
					need = (mid + k - 1) / k
				} else {
					need = mid / k
				}
				if need <= -room {
					dp[u] = 0
				} else {
					dp[u] = need
				}
				continue
			}

			if need == 0 {
				dp[u] = room - 1
				continue
			}

			need--
			if need <= -room {
				dp[u] = 0
			} else {
				dp[u] = need
			}
		}
		return dp[root] != 0
	}

	return sort.Search(n, check) - 1
}
