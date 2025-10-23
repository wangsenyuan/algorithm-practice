package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res[0], res[1])
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func comb3(n int) int {
	return n * (n - 1) * (n - 2) / 6
}

func comb2(n int) int {
	return n * (n - 1) / 2
}

func solve(n int, edges [][]int) []int {
	if len(edges) == 0 {
		return []int{3, comb3(n)}
	}

	g := NewGraph(n, len(edges)*2)
	// 如果无法生成二部图，那么就不用添加边
	// 如果存在长度为3的path, 那么只需要添加一条边，这个可以分层，隔一层的添加
	// 如果只存在长度为2的path， a - b, 答案 = 2, 任意一段和其他的任何一个点，连接
	// 如果只有长度为1的path， 都是独立的点, 那么必须添加3条边, C(n, 3)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	color := make([]int, n)
	for i := range n {
		color[i] = -1
	}
	cnt := make([]int, 2)

	var dfs func(u int, c int) bool

	dfs = func(u int, c int) bool {
		if color[u] >= 0 {
			return color[u] == c
		}
		color[u] = c
		cnt[c]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v, c^1) {
				return false
			}
		}
		return true
	}

	var res1 int

	for i := range n {
		if color[i] >= 0 {
			continue
		}
		clear(cnt)
		if !dfs(i, 0) {
			// 已经存在奇数长度的cycle
			return []int{0, 1}
		}
		res1 += comb2(cnt[0]) + comb2(cnt[1])
	}

	if res1 > 0 {
		return []int{1, res1}
	}
	res2 := len(edges) * (n - 2)

	return []int{2, res2}
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
