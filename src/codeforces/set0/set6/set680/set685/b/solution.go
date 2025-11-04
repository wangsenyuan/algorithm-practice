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

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	P := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &P[i])
	}
	Q := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &Q[i])
	}
	return solve(n, P, Q)
}

func solve(n int, P []int, Q []int) []int {
	g := NewGraph(n, n)

	for i := 2; i <= n; i++ {
		j := P[i-2] - 1
		g.AddEdge(j, i-1)
	}

	chain := make([]int, n)

	for i := range n {
		chain[i] = -1
	}

	ans := make([]int, n)

	pos := make([]int, n)
	var ord int
	sz := make([]int, n)

	ancestor := func(u int, v int) bool {
		return pos[u] <= pos[v] && pos[v] < pos[u]+sz[u]
	}
	big := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		pos[u] = ord
		ord++
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			sz[u] += sz[v]
		}
		big[u] = -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if sz[v]*2 >= sz[u] {
				big[u] = v
				break
			}
		}
	}

	dfs(0)

	find := func(root int, size int, d int) int {

		check := func(i int) bool {
			v := chain[i]
			if v < 0 || !ancestor(root, v) {
				return false
			}
			// v >= 0 and u is ancestor of v
			return sz[v]*2 >= size
		}

		lo, hi := d, n
		for lo < hi {
			mid := (lo + hi) >> 1
			if !check(mid) {
				hi = mid
			} else {
				lo = mid + 1
			}
		}

		if hi == d {
			return -1
		}

		return chain[hi-1]
	}

	var dfs2 func(u int, d int)

	dfs2 = func(u int, d int) {
		ans[u] = u

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != big[u] {
				dfs2(v, d+1)
			}
		}

		// 保证big的链路是最后被添加上去的
		if big[u] >= 0 {
			dfs2(big[u], d+1)
			chain[d+1] = big[u]
			// 肯定能找到一个，因为big[u]就是一个符合条件的节点
			ans[u] = find(u, sz[u], d+1)
		}
	}

	dfs2(0, 0)

	// 稍等，有点混乱了
	// 假设u子树中的centroid是v,
	res := make([]int, len(Q))
	for i, v := range Q {
		res[i] = ans[v-1] + 1
	}

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
