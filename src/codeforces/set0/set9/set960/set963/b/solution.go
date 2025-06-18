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
	_, _, res := process(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (n int, p []int, res []int) {
	n = readNum(reader)
	p = readNNums(reader, n)
	res = solve(n, p)
	return
}

func solve(n int, p []int) []int {
	if n%2 == 0 {
		return nil
	}
	// n is odd
	g := NewGraph(n, n)
	var root int
	for i := 0; i < n; i++ {
		if p[i] > 0 {
			g.AddEdge(p[i]-1, i)
		} else {
			root = i
		}
	}

	sz := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			sz[u] += sz[v]
		}
		sz[u]++
	}
	dfs(root)

	var res []int

	var dfs2 func(u int)
	dfs2 = func(u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if sz[v]%2 == 0 {
				dfs2(v)
			}
		}
		res = append(res, u+1)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if sz[v]%2 == 1 {
				dfs2(v)
			}
		}
	}

	dfs2(root)

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
