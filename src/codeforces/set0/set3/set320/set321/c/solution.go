package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _ := process(reader)
	fmt.Println(res)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (string, [][]int) {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges), edges
}

func solve(n int, edges [][]int) string {

	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	marked := make([]bool, n)

	sz := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u] = 1

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && !marked[v] {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(0, 0)

	var findCentroid func(p int, u int, exp int) int

	findCentroid = func(p int, u int, exp int) int {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && !marked[v] && sz[v]*2 >= exp {
				return findCentroid(u, v, exp)
			}
		}
		return u
	}

	root := findCentroid(0, 0, n)

	ans := make([]byte, n)

	var dfs2 func(u int, c byte)

	dfs2 = func(u int, c byte) {
		ans[u] = c
		marked[u] = true
		sz[u] = 0

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !marked[v] {
				dfs(u, v)
				root := findCentroid(-1, v, sz[v])
				dfs2(root, c+1)
			}
		}
	}

	dfs2(root, 'A')

	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(ans[i])
		buf.WriteByte(' ')
	}
	return buf.String()
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
