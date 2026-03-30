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
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	marked := make([]bool, n)

	que := make([]int, n)
	var head, tail int
	que[head] = 0
	head++
	marked[0] = true
	res := 1
	for tail < head {
		// 这一层又cnt个
		cnt := head - tail
		for range cnt {
			u := que[tail]
			tail++
			var childrenCnt int
			for _, v := range adj[u] {
				if !marked[v] {
					marked[v] = true
					que[head] = v
					head++
					childrenCnt++
				}
			}

			res = max(res, childrenCnt+1+max(0, cnt-1-childrenCnt))
		}

	}

	return res
}
