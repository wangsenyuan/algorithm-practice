package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, k, res := drive(reader)
	if k < 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", k))
	for i := 0; i < 2; i++ {
		for j := 0; j < len(res[i]); j++ {
			if j > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteString(fmt.Sprintf("%d", res[i][j]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) (n int, edges [][]int, k int, paths [][]int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	edges = make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	k, paths = solve(n, edges)
	return
}

type state struct {
	dist int
	prev int
}

func solve(n int, edges [][]int) (int, [][]int) {
	adj := make([][]int, n)
	words := (n + 63) / 64
	adjBits := make([][]uint64, n)
	vis := make([][]uint64, n)
	for i := range n {
		adjBits[i] = make([]uint64, words)
		vis[i] = make([]uint64, words)
	}
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		adjBits[u][v/64] |= 1 << (v % 64)
		adjBits[v][u/64] |= 1 << (u % 64)
	}

	dist := make([]state, n*n)
	for i := range dist {
		dist[i] = state{dist: -1, prev: -1}
	}

	id := func(u int, v int) int {
		return u*n + v
	}

	start := id(0, n-1)
	target := id(n-1, 0)
	que := make([]int, 0, n*n)
	que = append(que, start)
	dist[start].dist = 0
	vis[0][(n-1)/64] |= 1 << ((n - 1) % 64)

	for front := 0; front < len(que); front++ {
		cur := que[front]
		if cur == target {
			break
		}
		u1, v1 := cur/n, cur%n
		for _, u2 := range adj[u1] {
			for w := 0; w < words; w++ {
				cand := adjBits[v1][w] &^ vis[u2][w]
				if u2/64 == w {
					cand &^= 1 << (u2 % 64)
				}
				for cand > 0 {
					b := bits.TrailingZeros64(cand)
					v2 := w*64 + b
					if v2 < n {
						next := id(u2, v2)
						dist[next] = state{dist: dist[cur].dist + 1, prev: cur}
						vis[u2][w] |= 1 << b
						que = append(que, next)
					}
					cand &= cand - 1
				}
			}
		}
	}

	if dist[target].dist == -1 {
		return -1, nil
	}
	path := make([][]int, 2)
	for cur := target; cur >= 0; cur = dist[cur].prev {
		u, v := cur/n, cur%n
		path[0] = append(path[0], u+1)
		path[1] = append(path[1], v+1)
	}

	for i := range 2 {
		slices.Reverse(path[i])
	}

	return dist[target].dist, path
}
