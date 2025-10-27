package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	x := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	return solve(n, x, edges)
}

func solve(n int, x []int, edges [][]int) float64 {
	var ans float64
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		cur := float64(x[u]+x[v]) / float64(w)
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
