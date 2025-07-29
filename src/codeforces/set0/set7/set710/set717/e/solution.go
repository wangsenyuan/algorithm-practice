package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}
func process(reader *bufio.Reader) (color []int, edges [][]int, res []int) {
	n := readNum(reader)
	color = make([]int, n)
	for i := range n {
		color[i] = readNum(reader)
	}
	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	res = solve(slices.Clone(color), edges)
	return
}

func solve(color []int, edges [][]int) []int {
	n := len(color)
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dp := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		dp[u] = color[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				if dp[v] < 0 {
					dp[u] = -1
				}
			}
		}
	}
	dfs(-1, 0)
	if dp[0] > 0 {
		// 全部是black的
		return []int{1}
	}

	var res []int
	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		res = append(res, u+1)
		if u > 0 {
			color[u] *= -1
		}
		var a_child int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				a_child = v
			}
			if p == v || dp[v] > 0 {
				// 黑色的子树不处理
				continue
			}
			dfs2(u, v)
			res = append(res, u+1)
			color[u] *= -1
		}
		if color[u] < 0 {
			// u还是pink的
			if u > 0 {
				res = append(res, p+1)
				color[p] *= -1
				res = append(res, u+1)
				color[u] *= -1
			} else {
				res = append(res, a_child+1)
				color[a_child] *= -1
				res = append(res, u+1)
				color[u] *= -1
				res = append(res, a_child+1)
				color[a_child] *= -1
			}
		}
	}

	dfs2(-1, 0)

	return res
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
