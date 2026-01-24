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
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	// store adjacency as int32 to save memory
	adj := make([][]int32, n)
	for i := 0; i < n-1; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		x--
		y--
		adj[x] = append(adj[x], int32(y))
		adj[y] = append(adj[y], int32(x))
	}
	return solve(n, adj)
}

func solve(n int, adj [][]int32) []int {
	ans := make([]int, n)

	// depth-count arrays stored as int32 to reduce memory footprint
	merge := func(a []int32, b []int32) ([]int32, int) {
		if len(a) > len(b) {
			a, b = b, a
		}
		// len(a) <= len(b)
		// 后对齐的
		res := len(b) - 1
		for i := len(a) - 1; i >= 0; i-- {
			j := len(b) - 1 - (len(a) - 1 - i)
			b[j] += a[i]
			if b[j] > b[res] {
				res = j
			}
		}
		return b, len(b) - 1 - res
	}

	// iterative DFS using an explicit stack to avoid recursion
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}
	order := make([]int, 0, n)
	stack := []int{0}
	parent[0] = -1

	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, u)
		for _, vv := range adj[u] {
			v := int(vv)
			if v == parent[u] {
				continue
			}
			parent[v] = u
			stack = append(stack, v)
		}
	}

	// dp[u] stores the depth-count array for subtree rooted at u
	dp := make([][]int32, n)

	// process nodes in reverse order to simulate post-order
	for i := n - 1; i >= 0; i-- {
		u := order[i]
		var a []int32
		ans[u] = -1

		for _, vv := range adj[u] {
			v := int(vv)
			if v == parent[u] {
				continue
			}
			b := dp[v]
			// choose better child as base
			if ans[u] == -1 ||
				int(a[len(a)-1-ans[u]]) < int(b[len(b)-1-ans[v]]) ||
				(a[len(a)-1-ans[u]] == b[len(b)-1-ans[v]] && ans[u] > ans[v]) {
				ans[u] = ans[v]
				a, b = b, a
			}
			c, x := merge(a, b)
			if len(a) == 0 ||
				int(a[len(a)-1-ans[u]]) < int(c[len(c)-1-x]) ||
				(a[len(a)-1-ans[u]] == c[len(c)-1-x] && ans[u] > x) {
				ans[u] = x
			}
			a = c
		}

		if ans[u] != -1 && a[len(a)-1-ans[u]] == 1 {
			ans[u] = 0
		} else {
			ans[u]++
		}

		a = append(a, 1)
		dp[u] = a
	}

	return ans
}
