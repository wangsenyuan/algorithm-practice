package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := process(reader)
	if res < 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Println(res)
	}
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

func process(reader *bufio.Reader) (int, [][]int, []int) {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	color := readNNums(reader, n)
	return solve(n, edges, color), edges, color
}

func solve(n int, edges [][]int, color []int) int {

	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	// dp[u] 表示，是否u的子树，都和u是同样的颜色
	dp := make([]bool, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		dp[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				if !dp[v] || color[v] != color[u] {
					dp[u] = false
				}
			}
		}
	}

	dfs(-1, 0)

	if dp[0] {
		// 所有的都是同一个颜色
		return 1
	}

	var dfs2 func(p int, u int) int
	dfs2 = func(p int, u int) int {
		// dp[u] = false
		// 不管dp[p]
		var bad []int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if !dp[v] {
					bad = append(bad, v)
				}
			}
		}
		if len(bad) > 1 {
			return -1
		}
		if len(bad) == 0 {
			return u
		}
		return dfs2(u, bad[0])
	}

	r := dfs2(-1, 0)
	if r == -1 {
		return -1
	}
	// 检查r是否ok

	dfs(-1, r)

	for i := g.nodes[r]; i > 0; i = g.next[i] {
		v := g.to[i]
		if !dp[v] {
			return -1
		}
	}

	return r + 1
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
