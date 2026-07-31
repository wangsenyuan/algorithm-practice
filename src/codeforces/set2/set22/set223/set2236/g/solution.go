package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		for _, v := range drive(reader) {
			fmt.Fprintln(writer, v)
		}
	}
}

func drive(reader *bufio.Reader) []int64 {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	edges := make([][]int, n-1)
	for i := range n - 1 {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	queries := make([][]int, q)
	for i := range q {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		queries[i] = []int{x, y}
	}
	return solve(a, edges, queries)
}

func solve(a []int, edges [][]int, queries [][]int) []int64 {
	n := len(a)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	h := bits.Len(uint(n))
	up := make([][]int, n)
	lowDepthOfBit := make([]int, 20)
	for i := range 20 {
		lowDepthOfBit[i] = -1
	}

	depth := make([]int, n)
	fp := make([]int, n)
	for i := range fp {
		fp[i] = -1
	}
	pr := make([]int64, n)
	prevNonZero := make([]int, n)
	for i := range prevNonZero {
		prevNonZero[i] = -1
	}

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		up[u] = make([]int, h)
		up[u][0] = p
		for i := 1; i < h; i++ {
			up[u][i] = up[up[u][i-1]][i-1]
		}

		var keep [20]int
		copy(keep[:], lowDepthOfBit)

		if p != u {
			fp[u] = fp[p]
			if a[p] != 0 {
				prevNonZero[u] = p
			} else {
				prevNonZero[u] = prevNonZero[p]
			}
		}

		for i := range 20 {
			if (a[u]>>i)&1 == 1 {
				fp[u] = max(fp[u], lowDepthOfBit[i])
				lowDepthOfBit[i] = depth[u]
			}
		}
		pr[u] = int64(depth[u] - fp[u])
		if p != u {
			pr[u] += pr[p]
		}

		for _, v := range adj[u] {
			if p != v {
				depth[v] = depth[u] + 1
				dfs(u, v)
			}
		}

		copy(lowDepthOfBit, keep[:])
	}

	dfs(0, 0)

	lca := func(u int, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		for i := h - 1; i >= 0; i-- {
			if depth[u]-1<<i >= depth[v] {
				u = up[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := h - 1; i >= 0; i-- {
			if up[u][i] != up[v][i] {
				u = up[u][i]
				v = up[v][i]
			}
		}
		return up[u][0]
	}

	deepestGood := func(x int, w int, strict bool) int {
		dw := depth[w]
		good := func(u int) bool {
			if strict {
				return fp[u] < dw
			}
			return fp[u] <= dw
		}
		if good(x) {
			return x
		}

		// Move to the shallowest bad vertex. Its parent is the deepest good one.
		bad := x
		for i := h - 1; i >= 0; i-- {
			ancestor := up[bad][i]
			if depth[ancestor] > dw && !good(ancestor) {
				bad = ancestor
			}
		}
		return up[bad][0]
	}

	countArm := func(x int, w int) int64 {
		good := deepestGood(x, w, false)
		d := int64(depth[good] - depth[w])
		return d*(d+1)/2 + pr[x] - pr[good]
	}

	type state struct {
		mask int
		ways int64
	}

	buildStates := func(x int, w int) []state {
		good := deepestGood(x, w, true)

		nodes := make([]int, 0, 20)
		u := good
		if a[u] == 0 {
			u = prevNonZero[u]
		}
		for u >= 0 && depth[u] > depth[w] {
			nodes = append(nodes, u)
			u = prevNonZero[u]
		}

		slices.Reverse(nodes)

		states := make([]state, 0, len(nodes)+1)
		mask := 0
		lastDepth := depth[w]
		for _, u := range nodes {
			states = append(states, state{
				mask: mask,
				ways: int64(depth[u] - lastDepth),
			})
			mask |= a[u]
			lastDepth = depth[u]
		}
		states = append(states, state{
			mask: mask,
			ways: int64(depth[good] - lastDepth + 1),
		})
		return states
	}

	play := func(x int, y int) int64 {
		w := lca(x, y)
		left := buildStates(x, w)
		right := buildStates(y, w)

		var cross int64
		for _, p := range left {
			for _, q := range right {
				if p.mask&q.mask == 0 {
					cross += p.ways * q.ways
				}
			}
		}

		return countArm(x, w) + countArm(y, w) + cross
	}

	ans := make([]int64, len(queries))
	for i, c := range queries {
		ans[i] = play(c[0]-1, c[1]-1)
	}
	return ans
}
