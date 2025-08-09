package main

import (
	"bufio"
	"fmt"
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
	for i := range edges {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
	}
	var ans int
	for _, d := range deg {
		ans += d*(d-1) - d*(n-d-1) + (n-d-1)*(n-d-2)
	}
	return ans / 6
}
