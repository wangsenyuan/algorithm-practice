package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
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
	adj := make([][]int, n)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}

	que := make([]int, n)
	marked := make([]bool, n)

	// 只关心高位情况
	check := func(exp int, d int) bool {
		var head, tail int
		clear(marked)
		que[head] = 0
		head++
		marked[0] = true

		for tail < head {
			u := que[tail]
			tail++
			for _, i := range adj[u] {
				e := edges[i]
				v := (e[0] - 1) ^ (e[1] - 1) ^ u
				w := e[2]
				if (w>>d)&(exp>>d) == (w>>d) && !marked[v] {
					marked[v] = true
					que[head] = v
					head++
				}
			}
		}

		return marked[n-1]
	}

	var res int
	for d := 30; d >= 0; d-- {
		if !check(res, d) {
			res |= 1 << d
		}
	}

	return res
}
