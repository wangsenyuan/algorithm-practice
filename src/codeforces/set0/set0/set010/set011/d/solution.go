package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	dp := make([][]int, 1<<n)

	for i := range 1 << n {
		dp[i] = make([]int, n)
	}
	for i := range n {
		dp[1<<i][i] = 1
	}

	var sum int

	for mask := range 1 << n {
		st := bits.TrailingZeros(uint(mask))
		for v, fv := range dp[mask] {
			if fv > 0 {
				for _, w := range adj[v] {
					if w == st {
						// 产生了环
						sum += fv
					} else if w > st && mask&(1<<w) == 0 {
						dp[mask|(1<<w)][w] += fv
					}
				}
			}
		}
	}

	return (sum - len(edges)) / 2
}

func solve1(n int, edges [][]int) int {
	g := make([]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u] |= 1 << v
		g[v] |= 1 << u
	}
	N := 1 << n

	dp := make([][]int, N)
	for i := range N {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[1<<i][i] = 1
	}

	for mask := 1; mask < N; mask++ {
		f := lowestSetBit(mask)
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 && i != f {
				state := mask ^ (1 << i)
				for j := 0; j < n; j++ {
					if (g[i]>>j)&1 == 1 {
						dp[mask][i] += dp[state][j]
					}
				}
			}
		}
	}

	var ans int

	for state := 1; state < N; state++ {
		cnt := bits.OnesCount(uint(state))
		if cnt < 3 {
			continue
		}
		f := lowestSetBit(state)
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 && (g[f]>>i)&1 == 1 {
				ans += dp[state][i]
			}
		}
	}
	return ans / 2
}

func lowestSetBit(num int) int {
	for i := 0; i < 30; i++ {
		if (num>>i)&1 == 1 {
			return i
		}
	}
	return -1
}
