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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Fscan(reader, &p[i-2])
	}
	return solve(n, p)
}

func solve(n int, p []int) int {
	adj := make([][]int, n)
	for i := range n - 1 {
		adj[p[i]-1] = append(adj[p[i]-1], i+1)
	}

	var res int
	var dfs func(u int) int
	dfs = func(u int) int {
		first, second := -1, -1
		for _, v := range adj[u] {
			t := dfs(v) + 1
			if t >= first {
				second, first = first, t
			} else if t >= second {
				second = t
			}
		}

		res++
		res += second + 1
		return first
	}

	dfs(0)

	return res
}
