package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
}

func process(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	parent := make([]int, n)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &parent[i])
		parent[i]--
		children[parent[i]] = append(children[parent[i]], i)
	}

	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
		edges[i][0]--
		edges[i][1]--
	}

	return solve(parent, children, edges)
}

func solve(parent []int, children [][]int, edges [][]int) int {
	n := len(parent)
	depth := make([]int, n)
	size := make([]int, n)
	heavy := make([]int, n)
	for i := range n {
		size[i] = 1
		heavy[i] = -1
	}

	for u := 1; u < n; u++ {
		depth[u] = depth[parent[u]] + 1
	}
	for u := n - 1; u > 0; u-- {
		p := parent[u]
		size[p] += size[u]
		if heavy[p] < 0 || size[u] > size[heavy[p]] {
			heavy[p] = u
		}
	}

	top := make([]int, n)
	dfn := make([]int, n)
	var timer int
	stack := []int{0}
	for len(stack) > 0 {
		head := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for u := head; u >= 0; u = heavy[u] {
			top[u] = head
			dfn[u] = timer
			timer++
			for _, v := range children[u] {
				if v != heavy[u] {
					stack = append(stack, v)
				}
			}
		}
	}

	lca := func(u, v int) int {
		for top[u] != top[v] {
			if depth[top[u]] < depth[top[v]] {
				v = parent[top[v]]
			} else {
				u = parent[top[u]]
			}
		}
		if depth[u] < depth[v] {
			return u
		}
		return v
	}

	type path struct {
		u int
		v int
		w int
	}
	byLca := make([][]path, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		x := lca(u, v)
		byLca[x] = append(byLca[x], path{u, v, w})
	}

	tr := NewSegTree(n)
	queryPath := func(u, v int) int {
		var res int
		for top[u] != top[v] {
			if depth[top[u]] < depth[top[v]] {
				u, v = v, u
			}
			res += tr.Query(dfn[top[u]], dfn[u])
			u = parent[top[u]]
		}
		if depth[u] > depth[v] {
			u, v = v, u
		}
		res += tr.Query(dfn[u], dfn[v])
		return res
	}

	dp := make([]int, n)
	for u := n - 1; u >= 0; u-- {
		var sum int
		for _, v := range children[u] {
			sum += dp[v]
		}
		dp[u] = sum
		for _, p := range byLca[u] {
			dp[u] = max(dp[u], sum+p.w+queryPath(p.u, p.v))
		}
		tr.Add(dfn[u], sum-dp[u])
	}

	return dp[0]
}

type SegTree struct {
	arr []int
}

func NewSegTree(n int) *SegTree {
	return &SegTree{make([]int, 2*n)}
}

func (tr *SegTree) Add(pos int, v int) {
	pos += len(tr.arr) / 2
	for pos > 0 {
		tr.arr[pos] += v
		pos >>= 1
	}
}

func (tr *SegTree) Query(l int, r int) int {
	l += len(tr.arr) / 2
	r += len(tr.arr)/2 + 1
	var res int
	for l < r {
		if l&1 == 1 {
			res += tr.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += tr.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
