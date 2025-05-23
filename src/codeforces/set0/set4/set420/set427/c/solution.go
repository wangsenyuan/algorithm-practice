package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%d %d\n", res[0], res[1])
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	cost := readNNums(reader, n)
	m := readNum(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	return solve(cost, edges)
}

const mod = 1e9 + 7

func mul(a, b int) int {
	return a * b % mod
}

func solve(cost []int, edges [][]int) []int {
	n := len(cost)
	g := NewGraph(n, len(edges))
	r := NewGraph(n, len(edges))
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		r.AddEdge(v, u)
	}

	vis := make([]bool, n)
	var order []int
	var dfs func(u int)
	dfs = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
		order = append(order, u)
	}

	for u := range n {
		if !vis[u] {
			dfs(u)
		}
	}

	var cur int
	var sz int

	var dfs2 func(u int)
	dfs2 = func(u int) {
		if cost[u] < cur {
			sz = 1
			cur = cost[u]
		} else if cost[u] == cur {
			sz++
		}
		vis[u] = true
		for i := r.nodes[u]; i > 0; i = r.next[i] {
			v := r.to[i]
			if !vis[v] {
				dfs2(v)
			}
		}
	}

	ways := 1
	var sum int
	clear(vis)
	for i := n - 1; i >= 0; i-- {
		if !vis[order[i]] {
			cur = 1 << 30
			sz = 0
			dfs2(order[i])
			sum += cur
			ways = mul(ways, sz)
		}
	}

	return []int{sum, ways}
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
