package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tc int
	fmt.Fscan(in, &tc)
	for range tc {
		_, res := drive(in)
		fmt.Fprintln(out, len(res))
		for _, cur := range res {
			fmt.Fprint(out, len(cur))
			for _, x := range cur {
				fmt.Fprint(out, " ", x)
			}
			fmt.Fprintln(out)
		}
	}
}

func drive(reader *bufio.Reader) (edges [][]int, res [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	edges = make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) [][]int {
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	fa := make([]int, n)
	for i := range n {
		fa[i] = -1
	}
	assign := make([]int, n)

	que := make([]int, n)
	var head, tail int

	fa[0] = 0
	que[head] = 0
	head++

	marked := make([]bool, n)

	for tail < head {
		mark := head
		var bad []int
		var rest []int

		for i := tail; i < mark; i++ {
			u := que[i]
			if u > 0 {
				if !marked[assign[fa[u]]] {
					assign[u] = assign[fa[u]]
					marked[assign[u]] = true
					bad = append(bad, u)
				} else {
					rest = append(rest, u)
				}
			}
			for _, v := range adj[u] {
				if fa[v] == -1 {
					fa[v] = u
					que[head] = v
					head++
				}
			}
		}
		var cur int
		for _, v := range rest {
			for marked[cur] {
				cur++
			}
			assign[v] = cur
			marked[cur] = true
		}

		if len(bad) > 1 {
			// 循环替换
			first := assign[bad[0]]
			for i := 0; i+1 < len(bad); i++ {
				assign[bad[i]] = assign[bad[i+1]]
			}
			assign[bad[len(bad)-1]] = first
		} else if len(bad) == 1 {
			// 分配一个新的颜色
			old := assign[bad[0]]
			marked[old] = false
			var cur int
			for marked[cur] || cur == old {
				cur++
			}
			assign[bad[0]] = cur
			marked[cur] = true
		}

		for i := tail; i < mark; i++ {
			u := que[i]
			marked[assign[u]] = false
		}

		tail = mark
	}

	k := slices.Max(assign)
	res := make([][]int, k+1)
	for i := range n {
		res[assign[i]] = append(res[assign[i]], i+1)
	}

	return res
}
