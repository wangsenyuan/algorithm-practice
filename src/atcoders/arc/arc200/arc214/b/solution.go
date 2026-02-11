package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		for j := range 3 {
			fmt.Fscan(reader, &edges[i][j])
		}
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	if n&1 == 1 {
		return -1
	}

	adj := make([][]int, n)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}

	marked := make([]bool, n)
	var que []int
	que = append(que, 0)
	marked[0] = true
	a := make([]int, n)

	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		for _, i := range adj[u] {
			v := (edges[i][0] - 1) ^ (edges[i][1] - 1) ^ u
			w := edges[i][2]
			if !marked[v] {
				a[v] = a[u] ^ w
				marked[v] = true
				que = append(que, v)
			}
		}
	}

	var res int
	for i := range n + 1 {
		res ^= i
	}
	for i := range n {
		res ^= a[i]
	}

	return res
}
