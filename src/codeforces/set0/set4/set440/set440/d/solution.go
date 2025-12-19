package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, res := drive(reader)
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (n int, k int, edges [][]int, res []int) {
	fmt.Fscan(reader, &n, &k)
	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	res = solve(n, k, edges)
	return
}

type pair struct {
	first  int
	second int
}

func solve(n int, k int, edges [][]int) []int {
	adj := make([][]pair, n)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], pair{v, i})
		adj[v] = append(adj[v], pair{u, i})
	}

	dp := make([][]int, n)
	fp := make([][][]int, n)
	for u := range n {
		dp[u] = make([]int, k+1)
		fp[u] = make([][]int, k+1)
	}

	fa := make([]int, n)
	sz := make([]int, n)
	var dfs func(p int, u int)

	ndp := make([]int, k+1)

	dfs = func(p int, u int) {
		fa[u] = p
		sz[u] = 1
		for j := range k + 1 {
			dp[u][j] = n
		}
		dp[u][1] = 0
		fp[u][1] = []int{}

		for _, e := range adj[u] {
			v := e.first
			eid := e.second
			if p != v {
				dfs(u, v)

				nfp := make([][]int, k+1)
				for j := range k + 1 {
					ndp[j] = n
				}

				for w := 1; w <= k; w++ {
					if dp[u][w]+1 < ndp[w] {
						ndp[w] = dp[u][w] + 1
						nfp[w] = slices.Clone(fp[u][w])
						nfp[w] = append(nfp[w], eid)
					}
					for x := 1; x < w && x <= sz[v]; x++ {
						if dp[u][w-x]+dp[v][x] < ndp[w] {
							ndp[w] = dp[u][w-x] + dp[v][x]
							nfp[w] = slices.Clone(fp[u][w-x])
							nfp[w] = append(nfp[w], fp[v][x]...)
						}
					}
				}

				for w := 1; w <= k; w++ {
					fp[u][w] = nfp[w]
					dp[u][w] = ndp[w]
					ndp[w] = n
				}

				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0)

	best := dp[0][k]
	var root int
	for u := 1; u < n; u++ {
		if dp[u][k]+1 < best {
			best = dp[u][k] + 1
			root = u
		}
	}

	var res []int
	for _, eid := range fp[root][k] {
		res = append(res, eid+1)
	}

	if root > 0 {
		for _, e := range adj[root] {
			if e.first == fa[root] {
				res = append(res, e.second+1)
				break
			}
		}
	}

	return res
}
