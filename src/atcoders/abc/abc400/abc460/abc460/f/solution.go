package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for _, ans := range drive(reader) {
		fmt.Println(ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([]int, q)
	for i := range q {
		fmt.Fscan(reader, &queries[i])
	}
	return solve(n, edges, queries)
}

func solve(n int, edges [][]int, queries []int) []int {
	adj := make([][]int, n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// pos := make([]int, n)
	// at := make([]int, n)
	dep := make([]int, n)
	first := make([]int, n)
	euler := make([]int, 0, 2*n-1)

	// var ptr int
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		// pos[u] = ptr
		// at[ptr] = u
		// ptr++
		first[u] = len(euler)
		euler = append(euler, u)
		for _, v := range adj[u] {
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
				euler = append(euler, u)
			}
		}
	}

	dfs(0, 0)

	m := len(euler)
	h := bits.Len(uint(m))
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, h)
		dp[i][0] = euler[i]
	}

	for j := 1; j < h; j++ {
		for i := 0; i+(1<<j) <= m; i++ {
			dp[i][j] = dp[i][j-1]
			k := dp[i+(1<<(j-1))][j-1]
			if dep[dp[i][j]] > dep[k] {
				dp[i][j] = k
			}
		}
	}

	getDist := func(u int, v int) int {
		l, r := first[u], first[v]
		if l > r {
			l, r = r, l
		}

		d := r - l + 1
		w := bits.Len(uint(d)) - 1
		p := dp[l][w]
		q := dp[r-(1<<w)+1][w]
		if dep[p] > dep[q] {
			p = q
		}

		return dep[u] + dep[v] - 2*dep[p]
	}

	merge := func(a state, b state) state {
		c := a
		if c.dist < b.dist {
			c = b
		}

		if len(a.nodes) == 2 && len(b.nodes) == 2 {
			for _, u := range a.nodes {
				for _, v := range b.nodes {
					d := getDist(u, v)
					if d > c.dist {
						c.dist = d
						c.nodes = []int{u, v}
					}
				}
			}
		}

		return c
	}

	tr := make([]state, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			tr[i] = state{[]int{l, l}, 0}
			return
		}
		mid := (l + r) >> 1
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		tr[i] = merge(tr[i*2+1], tr[i*2+2])
	}
	build(0, 0, n-1)

	var flip func(i int, l int, r int, pos int)
	flip = func(i int, l int, r int, pos int) {
		if l == r {
			if len(tr[i].nodes) == 0 {
				tr[i].nodes = []int{l, l}
				tr[i].dist = 0
			} else {
				tr[i].nodes = nil
				tr[i].dist = -inf
			}
			return
		}
		mid := (l + r) >> 1
		if pos <= mid {
			flip(i*2+1, l, mid, pos)
		} else {
			flip(i*2+2, mid+1, r, pos)
		}

		tr[i] = merge(tr[i*2+1], tr[i*2+2])
	}

	ans := make([]int, len(queries))

	for i, x := range queries {
		flip(0, 0, n-1, x-1)
		ans[i] = tr[0].dist
	}

	return ans
}

type state struct {
	nodes []int
	dist  int
}

const inf = 1 << 60
