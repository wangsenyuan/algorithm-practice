package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) string {
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	children := make([][]int, n)
	f := make([]*big.Int, n)
	g := make([]*big.Int, n)
	h := make([]*big.Int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		f[u] = big.NewInt(1)
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			dfs(u, v)
			children[u] = append(children[u], v)
			f[u] = mulBig(f[u], h[v])
		}

		sort.Slice(children[u], func(i, j int) bool {
			a, b := children[u][i], children[u][j]
			left := mulBig(f[a], h[b])
			right := mulBig(f[b], h[a])
			return left.Cmp(right) > 0
		})

		h[u] = cloneBig(f[u])
		g[u] = mulInt(f[u], 2)

		cur := cloneBig(f[u])
		for i, v := range children[u] {
			cur = divMulBig(cur, h[v], f[v])

			updateMax(&h[u], mulInt(cur, i+2))
			updateMax(&g[u], mulInt(cur, i+3))
		}

		for _, v := range children[u] {
			updateMax(&h[u], divMulBig(f[u], h[v], g[v]))
		}
	}

	dfs(-1, 0)

	return h[0].String()
}

func cloneBig(x *big.Int) *big.Int {
	return new(big.Int).Set(x)
}

func mulBig(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func mulInt(a *big.Int, b int) *big.Int {
	return new(big.Int).Mul(a, big.NewInt(int64(b)))
}

func divMulBig(a, b, c *big.Int) *big.Int {
	tmp := new(big.Int).Quo(a, b)
	return tmp.Mul(tmp, c)
}

func updateMax(dst **big.Int, cand *big.Int) {
	if (*dst).Cmp(cand) < 0 {
		*dst = cand
	}
}
