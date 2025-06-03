package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	x, res := process(reader)
	fmt.Println(x)
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

func process(reader *bufio.Reader) (int, []int) {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) (int, []int) {
	g := NewGraph(n, 2*n)
	for _, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v, 0)
		g.AddEdge(v, u, 1)
	}

	dp := make([]int, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				// 如果是 v -> u的，那么就需要加1
				dp[u] += dp[v] + g.val[i]
			}
		}
	}

	dfs(-1, 0)

	fp := make([]int, n)

	var dfs2 func(p int, u int, x int)

	dfs2 = func(p int, u int, x int) {
		fp[u] = dp[u] + x
		if p >= 0 {
			fp[u] += fp[p]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				fp[u] -= dp[v] + g.val[i]
				// 如果是p -> u 的，那么就需要反向
				dfs2(u, v, 1-g.val[i])
				fp[u] += dp[v] + g.val[i]
			}
		}
	}

	dfs2(-1, 0, 0)

	x := slices.Min(fp)
	var ans []int
	for i := range n {
		if fp[i] == x {
			ans = append(ans, i+1)
		}
	}

	return x, ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
