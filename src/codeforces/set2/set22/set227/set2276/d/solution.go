package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) string {
	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
	}
	if slices.Contains(deg, 0) {
		return "Brain Melt"
	}
	var cnt1 int
	var cnt2 int
	for _, v := range deg {
		if v == 1 {
			cnt1++
		} else if v == n-1 {
			cnt2++
		}
	}
	if cnt1 == n-1 && cnt2 == 1 {
		return "2"
	}
	return "Many"
}
