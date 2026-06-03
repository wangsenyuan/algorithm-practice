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

func solve(n int, edges [][]int) []int {
	adj := make([][]int, n)
	deg := make([]int, n)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
		deg[u]++
		deg[v]++
	}

	T := make([]bool, n)
	var st []int
	for i := range n {
		T[i] = true
		if i != 0 && i != n-1 && deg[i] == 1 {
			// 不是0和n-1的叶子节点
			st = append(st, i)
		}
	}

	for len(st) > 0 {
		v := st[len(st)-1]
		st = st[:len(st)-1]
		if v == 0 || v == n-1 {
			continue
		}
		T[v] = false
		for _, eid := range adj[v] {
			u := (edges[eid][0] - 1) ^ (edges[eid][1] - 1) ^ v
			deg[u]--
			if deg[u] == 1 {
				st = append(st, u)
			}
		}
	}

	// 剩余T[u] = true的节点是那些，可以从1到n，经过的点
	S := make([]bool, n)
	for i := range n {
		if i == 0 || i == n-1 || (T[i] == true && deg[i] >= 3) {
			// deg[i] = 2 的内部节点，可以被压缩
			S[i] = true
		}
	}

	H := make([][][]int, n)
	used := make([]bool, len(edges))

	findNext := func(par int, cur int) (u int, e int) {
		for _, eid := range adj[cur] {
			v := (edges[eid][0] - 1) ^ (edges[eid][1] - 1) ^ cur
			if T[v] && v != par {
				return v, eid
			}
		}
		return
	}

	for s := range n {
		if !S[s] {
			continue
		}

		for _, eid := range adj[s] {
			v := (edges[eid][0] - 1) ^ (edges[eid][1] - 1) ^ s
			if !T[v] || used[eid] {
				// 如果v不会经过path 1...n， 或者eid 已经被处理了
				continue
			}
			used[eid] = true
			l := 1
			par, cur := s, v
			for T[cur] && !S[cur] && deg[cur] == 2 {
				nxt, e2 := findNext(par, cur)
				used[e2] = true
				l++
				par, cur = cur, nxt
			}

			H[s] = append(H[s], []int{cur, l})
			H[cur] = append(H[cur], []int{s, l})
		}
	}

	ans := make([]int, n)
	vis := make([]bool, n)

	var dfs func(u int, d int)
	dfs = func(u int, d int) {
		if u == n-1 {
			ans[d]++
			return
		}
		vis[u] = true
		for _, cur := range H[u] {
			v, l := cur[0], cur[1]
			if !vis[v] {
				dfs(v, d+l)
			}
		}
		vis[u] = false
	}

	dfs(0, 0)

	return ans[1:]
}
