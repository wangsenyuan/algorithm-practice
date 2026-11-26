package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][2]int, m)
	for i := range m {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][2]int) int64 {
	fa := make([]int, n)
	for i := range n {
		fa[i] = i
	}
	find := func(u int) int {
		p := fa[u]
		for p != fa[p] {
			p = fa[p]
		}
		for p != fa[u] {
			fa[u], u = p, fa[u]
		}
		return p
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		u = find(u)
		v = find(v)
		if u != v {
			fa[u] = v
		}
	}
	var cnt int
	for u := range n {
		if find(u) == u {
			cnt++
		}
	}
	return 1 << (n - cnt)
}
