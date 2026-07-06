package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

const inf = 1 << 60

func solve(n int, edges [][]int) []int {
	// TODO: solve by hand first.
	adj := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	// 0 就是特殊点w

	que := make([]int, n)

	bfs := func(start int) []int {
		var head, tail int
		que[head] = start
		head++
		dist := make([]int, n+1)
		for i := range dist {
			dist[i] = inf
		}
		dist[start] = 0

		for tail < head {
			u := que[tail]
			tail++
			for _, v := range adj[u] {
				if v == 0 || dist[v] < inf {
					continue
				}
				dist[v] = dist[u] + 1
				que[head] = v
				head++
			}
		}
		return dist
	}

	dp1 := bfs(1)
	dp2 := bfs(n)

	// 现在需要知道到0的最近的距离
	var bestFrom1 [][]int
	var bestFromN [][]int
	for _, u := range adj[0] {
		if dp1[u] >= 0 {
			tmp := []int{u, dp1[u]}
			for i, cur := range bestFrom1 {
				if tmp[1] <= cur[1] {
					bestFrom1[i], tmp = tmp, cur
				}
			}
			if len(bestFrom1) < 3 {
				bestFrom1 = append(bestFrom1, tmp)
			}
		}
		if dp2[u] >= 0 {
			tmp := []int{u, dp2[u]}
			for i, cur := range bestFromN {
				if tmp[1] <= cur[1] {
					bestFromN[i], tmp = tmp, cur
				}
			}
			if len(bestFromN) < 3 {
				bestFromN = append(bestFromN, tmp)
			}
		}
	}

	getBestAvoiding := func(best [][]int, u int) []int {
		for _, cur := range best {
			if cur[0] != u {
				return cur
			}
		}
		return nil
	}

	play := func(u int) int {
		res := dp1[n]

		tmp1 := getBestAvoiding(bestFrom1, u)
		tmp2 := getBestAvoiding(bestFromN, u)

		if tmp1 != nil {
			// 连接到u, u还是连接到n
			res = min(res, tmp1[1]+1+dp2[u])
		}

		if tmp2 != nil {
			res = min(res, tmp2[1]+1+dp1[u])
		}

		if tmp1 != nil && tmp2 != nil {
			res = min(res, tmp1[1]+tmp2[1]+2)
		}

		return res
	}

	ans := make([]int, n+1)

	for u := 1; u <= n; u++ {
		ans[u] = play(u)
		if ans[u] == inf {
			ans[u] = -1
		}
	}

	return ans[1:]
}
