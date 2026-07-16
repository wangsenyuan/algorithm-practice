package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, b int
	fmt.Fscan(reader, &n, &b)
	c := make([]int, n)
	d := make([]int, n)
	x := make([]int, n)
	fmt.Fscan(reader, &c[0], &d[0])
	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &c[i], &d[i], &x[i])
	}
	return solve(b, c, d, x)
}

func solve(b int, c []int, d []int, x []int) int {
	n := len(c)
	adj := make([][]int, n)
	for i := 1; i < n; i++ {
		adj[x[i]-1] = append(adj[x[i]-1], i)
	}

	sz := make([]int, n)
	g := make([][]int, n)
	best := make([][]int, n)
	big := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		sz[u] = 1
		best[u] = append(best[u], c[u])
		big[u] = -1
		for _, v := range adj[u] {
			dfs(v)
			best[u] = append(best[u], best[v]...)
			sz[u] += sz[v]
			if big[u] < 0 || sz[big[u]] < sz[v] {
				big[u] = v
			}
		}
		slices.Sort(best[u])
		g[u] = make([]int, sz[u]+1)
		var sum int
		for i := range sz[u] {
			g[u][i] = sum
			sum += best[u][i]
		}
		g[u][sz[u]] = sum
	}

	dfs(0)

	var dfs2 func(u int) []int
	dfs2 = func(u int) []int {
		res := make([]int, sz[u]+1)
		res[0] = 0
		res[1] = c[u] - d[u]
		for i := 2; i <= sz[u]; i++ {
			res[i] = inf
		}
		if big[u] < 0 {
			return res
		}

		first := dfs2(big[u])
		cur := 1 + sz[big[u]]
		for i := 1; i <= cur; i++ {
			res[i] = min(res[i], first[i-1]+c[u]-d[u])
		}
		buf := slices.Clone(res)
		for _, v := range adj[u] {
			if big[u] == v {
				continue
			}
			tmp := dfs2(v)
			nxt := cur + sz[v]
			copy(buf[:nxt+1], res[:nxt+1])

			for j := 1; j <= cur; j++ {
				for k := 0; k <= sz[v]; k++ {
					buf[j+k] = min(buf[j+k], res[j]+tmp[k])
				}
			}
			cur = nxt
			copy(res[:cur+1], buf[:cur+1])
		}
		for i := range sz[u] + 1 {
			res[i] = min(res[i], g[u][i])
		}
		return res
	}

	res := dfs2(0)
	for i := n; i > 0; i-- {
		if res[i] <= b {
			return i
		}
	}

	return 0
}

const inf = 1 << 60
