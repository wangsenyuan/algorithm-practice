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
	var n int
	fmt.Fscan(reader, &n)
	nodes := make([][]int, n)
	for i := range n {
		nodes[i] = make([]int, 3)
		fmt.Fscan(reader, &nodes[i][0], &nodes[i][1], &nodes[i][2])
	}
	return solve(n, nodes)
}

func solve(n int, nodes [][]int) int {

	fa := make([]int, n)
	for i := range n {
		fa[i] = -1
	}

	for i, cur := range nodes {
		l, r := cur[1], cur[2]
		if l > 0 {
			fa[l-1] = i
		}
		if r > 0 {
			fa[r-1] = i
		}
	}

	var root int
	for i := range n {
		if fa[i] == -1 {
			root = i
			break
		}
	}

	marked := make(map[int]bool)

	var dfs func(u int, lo int, hi int)
	dfs = func(u int, lo int, hi int) {
		w := nodes[u][0]
		if w >= lo && w <= hi {
			marked[w] = true
		}
		l, r := nodes[u][1], nodes[u][2]
		if l > 0 {
			dfs(l-1, lo, min(w, hi))
		}
		if r > 0 {
			dfs(r-1, max(w, lo), hi)
		}
	}

	dfs(root, -inf, inf)

	var ans int
	for i := range n {
		if !marked[nodes[i][0]] {
			ans++
		}
	}

	return ans
}

const inf = 1 << 60
