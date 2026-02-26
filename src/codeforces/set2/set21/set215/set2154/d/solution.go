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
		_, _, res := drive(reader)
		fmt.Fprintln(writer, len(res))
		for _, row := range res {
			for _, v := range row {
				fmt.Fprint(writer, v, " ")
			}
			fmt.Fprintln(writer)
		}
	}
}

func drive(reader *bufio.Reader) (n int, edges [][]int, res [][]int) {
	fmt.Fscan(reader, &n)
	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) [][]int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	color := make([]int, n)
	for i := range n {
		color[i] = -1
	}
	dist := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		for _, v := range adj[u] {
			if color[v] == -1 {
				color[v] = color[u] ^ 1
				dist[v] = dist[u] + 1
				dfs(v)
			}
		}
	}
	color[n-1] = 0
	dfs(n - 1)

	arr := make([][]pair, 2)
	for i := range n {
		arr[color[i]] = append(arr[color[i]], pair{dist[i], i})
	}

	for i := range 2 {
		slices.SortFunc(arr[i], func(x pair, y pair) int {
			return y.first - x.first
		})
	}

	cat := color[0]

	var res [][]int
	cnt := n
	for cnt > 1 {
		// 看看能不能删除另外一个颜色的节点
		if len(arr[cat^1]) > 0 && arr[cat^1][0].first > arr[cat][0].first {
			u := arr[cat^1][0].second
			res = append(res, []int{2, u + 1})
			arr[cat^1] = arr[cat^1][1:]
			cnt--
		}
		// 但是可能存在相等的情况吗？不可能（因为颜色不一样，奇偶性不一样）
		res = append(res, []int{1})
		cat ^= 1
	}

	return res
}

type pair struct {
	first  int
	second int
}
