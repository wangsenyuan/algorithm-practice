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
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	rabbits := make([][]int, n)
	for i := range n {
		rabbits[i] = make([]int, 2)
		fmt.Fscan(reader, &rabbits[i][0], &rabbits[i][1])
	}
	return solve(rabbits)
}

func solve(rabbits [][]int) int {
	// n := len(rabbits)
	var xs []int
	for _, cur := range rabbits {
		// 先按照左边处理
		xs = append(xs, cur[0]-cur[1], cur[0]+cur[1])
	}
	slices.Sort(xs)
	xs = slices.Compact(xs)
	m := len(xs)
	adj := make([][]int, m)
	n := len(rabbits)

	edges := make([][]int, n)
	for i, cur := range rabbits {
		u := sort.SearchInts(xs, cur[0]-cur[1])
		v := sort.SearchInts(xs, cur[0]+cur[1])
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
		// avoid re-search
		edges[i] = []int{u, v}
	}

	que := make([]int, m)
	vis := make([]bool, m)

	marked := make([]bool, n)

	bfs := func(s int) int {
		var head, tail int
		vis[s] = true
		que[head] = s
		head++

		var cnt int

		for tail < head {
			u := que[tail]
			tail++
			for _, id := range adj[u] {
				if !marked[id] {
					cnt++
				}
				marked[id] = true
				v := edges[id][0] ^ edges[id][1] ^ u
				if !vis[v] {
					vis[v] = true
					que[head] = v
					head++
				}
			}
		}

		return min(cnt, head)
	}

	var ans int
	for i := range m {
		if !vis[i] {
			ans += bfs(i)
		}
	}

	return ans
}
