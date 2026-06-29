package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	const inf = 1 << 60

	dist := make([][]int, n)
	for i := range n {
		dist[i] = make([]int, n)
		for j := range n {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}

	deg := make([]int, n)
	var sum int
	for _, cur := range edges {
		u, v, w := cur[0]-1, cur[1]-1, cur[2]
		sum += w
		if u == v {
			deg[u] += 2
			continue
		}
		deg[u]++
		deg[v]++
		dist[u][v] = min(dist[u][v], w)
		dist[v][u] = min(dist[v][u], w)
	}

	for k := range n {
		for i := range n {
			for j := range n {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	for i := range n {
		if deg[i] > 0 && dist[0][i] >= inf {
			return -1
		}
	}

	var odd []int
	for i := range n {
		if deg[i]%2 == 1 {
			odd = append(odd, i)
		}
	}

	m := len(odd)
	dp := make([]int, 1<<m)
	for i := 1; i < len(dp); i++ {
		dp[i] = inf
	}
	for mask := 1; mask < 1<<m; mask++ {
		i := bits.TrailingZeros(uint(mask))
		for j := i + 1; j < m; j++ {
			if mask>>j&1 == 1 {
				dp[mask] = min(dp[mask], dp[mask^(1<<i)^(1<<j)]+dist[odd[i]][odd[j]])
			}
		}
	}
	return sum + dp[(1<<m)-1]
}
