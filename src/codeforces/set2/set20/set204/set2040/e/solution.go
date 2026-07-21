package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	var res []int
	for range t {
		var n, q int
		fmt.Fscan(reader, &n, &q)
		edges := make([][]int, n-1)
		for i := range n - 1 {
			edges[i] = make([]int, 2)
			fmt.Fscan(reader, &edges[i][0], &edges[i][1])
		}
		queries := make([][]int, q)
		for i := range q {
			queries[i] = make([]int, 2)
			fmt.Fscan(reader, &queries[i][0], &queries[i][1])
		}
		res = append(res, solve(n, edges, queries)...)
	}
	return res
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int, edges [][]int, queries [][]int) []int {
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var degs []int
	for i := range n {
		if i > 0 {
			degs = append(degs, len(adj[i])-1)
		} else {
			degs = append(degs, len(adj[i]))
		}
	}

	slices.Sort(degs)
	degs = slices.Compact(degs)

	todo := make([][]int, n)
	for i, cur := range queries {
		v := cur[0] - 1
		todo[v] = append(todo[v], i)
	}

	ans := make([]int, len(queries))

	// dp[v] = no coin expectation for v to start
	dp := make([]int, n)
	tr := make([]*Tree, 2)
	for i := range 2 {
		tr[i] = NewTree(n)
	}

	fa := make([]int, n)

	var dfs func(p int, u int, dep int)
	dfs = func(p int, u int, dep int) {
		if u > 0 {
			if p > 0 {
				dp[u] = dp[fa[p]] + 2*(len(adj[p]))
			} else {
				dp[u] = 1
			}
		}
		for _, i := range todo[u] {
			p := queries[i][1]
			d := (dep & 1) ^ 1
			p = min(p, tr[d].cnt[0])
			ans[i] = dp[u]
			if p > 0 {
				ans[i] -= 2 * tr[d].sumKmax(p)
			}
			ans[i] %= mod
			if ans[i] < 0 {
				ans[i] += mod
			}
		}
		if u > 0 {
			x := len(adj[u]) - 1
			i := sort.SearchInts(degs, x)
			tr[dep&1].add(i, x)
		}

		for _, v := range adj[u] {
			if p != v {
				fa[v] = u
				dfs(u, v, dep+1)
			}
		}

		if u > 0 {
			x := len(adj[u]) - 1
			i := sort.SearchInts(degs, x)
			tr[dep&1].remove(i, x)
		}
	}

	dfs(0, 0, 0)

	return ans
}

type Tree struct {
	cnt []int
	sum []int
}

func NewTree(n int) *Tree {
	return &Tree{
		cnt: make([]int, n*4),
		sum: make([]int, n*4),
	}
}

func (t *Tree) add(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if r == l {
			t.cnt[i]++
			t.sum[i] += v
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		t.cnt[i] = t.cnt[i*2+1] + t.cnt[i*2+2]
		t.sum[i] = add(t.sum[i*2+1], t.sum[i*2+2])
	}
	n := len(t.cnt) / 4
	f(0, 0, n-1)
}

func (t *Tree) remove(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if r == l {
			t.cnt[i]--
			t.sum[i] -= v
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		t.cnt[i] = t.cnt[i*2+1] + t.cnt[i*2+2]
		t.sum[i] = add(t.sum[i*2+1], t.sum[i*2+2])
	}
	n := len(t.cnt) / 4
	f(0, 0, n-1)
}

func (t *Tree) sumKmax(k int) int {
	var f func(i int, l int, r int, k int) int
	f = func(i int, l int, r int, k int) int {
		if t.cnt[i] == k {
			return t.sum[i]
		}
		if l == r {
			return t.sum[i] / t.cnt[i] * k
		}
		mid := (l + r) >> 1
		if t.cnt[i*2+2] >= k {
			return f(i*2+2, mid+1, r, k)
		}
		return f(i*2+1, l, mid, k-t.cnt[i*2+2]) + t.sum[i*2+2]
	}
	n := len(t.cnt) / 4
	return f(0, 0, n-1, k)
}
