package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

type pair struct {
	first  int
	second int
}

func solve(n int, lines [][]int) int {
	var edges [][]int
	mp := make(map[pair]int)

	g := NewGraph(n, 2*n)
	deg := make([]int, n)
	add := func(u int, v int) {
		mp[pair{u, v}] = len(edges)
		edges = append(edges, []int{u, v})
		g.AddEdge(u, v)
		deg[u]++
	}

	for _, cur := range lines {
		u, v := cur[0]-1, cur[1]-1
		add(u, v)
		add(v, u)
	}
	m := len(edges)
	dp := make([][]int, m)
	from := make([][]int, n)
	miss := make([]int, n)
	for i := range m {
		dp[i] = make([]int, 2)
		dp[i][0] = -1
		dp[i][1] = -1
	}
	for i := range n {
		miss[i] = -1
		from[i] = make([]int, 2)
	}

	var dfs func(e int)
	// 这段我看不懂
	dfs = func(e int) {
		if dp[e][0] >= 0 || dp[e][1] >= 0 {
			return
		}
		dp[e][0] = 0
		dp[e][1] = 1
		p := edges[e][0]
		u := edges[e][1]
		if miss[u] < 0 {
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p == v {
					continue
				}
				ne := mp[pair{u, v}]
				dfs(ne)
				from[u][0] += max(dp[ne][1], dp[ne][0])
				from[u][1] += dp[ne][0]
			}
			miss[u] = p
		}
		if miss[u] != p && miss[u] != n {
			ne := mp[pair{u, miss[u]}]
			dfs(ne)
			from[u][0] += max(dp[ne][1], dp[ne][0])
			from[u][1] += dp[ne][0]
			miss[u] = n
		}
		if miss[u] == n {
			ne := mp[pair{u, p}]
			dp[e][0] += from[u][0] - max(dp[ne][1], dp[ne][0])
			dp[e][1] += from[u][1] - dp[ne][0]
		} else {
			dp[e][0] += from[u][0]
			dp[e][1] += from[u][1]
		}
	}

	for i := range m {
		dfs(i)
	}
	var ans int
	for u := range n {
		if deg[u] == 1 {
			p := g.to[g.nodes[u]]
			e := mp[pair{u, p}]
			ans = max(ans, 1+max(dp[e][0], dp[e][1]))
		}
	}
	return ans
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
