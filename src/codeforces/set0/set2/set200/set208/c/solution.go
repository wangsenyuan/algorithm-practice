package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.12f\n", res)
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

func process(reader *bufio.Reader) float64 {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}
func solve(n int, edges [][]int) float64 {
	g := NewGraph(n, 2*len(edges))
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	que := make([]int, n)

	bfs := func(s int) (dist []int, dp []int) {
		dist = make([]int, n)
		dp = make([]int, n)
		for i := range n {
			dist[i] = -1
			dp[i] = 0
		}
		dist[s] = 0
		dp[s] = 1
		var head, tail int
		que[head] = s
		head++

		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					dp[v] = dp[u]
					que[head] = v
					head++
				} else if dist[v] == dist[u]+1 {
					dp[v] = min(inf, dp[v]+dp[u])
				}
			}
		}
		return
	}

	d1, dp1 := bfs(0)
	d2, dp2 := bfs(n - 1)

	if d1[n-1] == 1 {
		return 1.0
	}

	var res float64
	for u := range n {
		if d1[u]+d2[u] == d1[n-1] {
			// u在最短路径上
			cnt1 := dp1[u]
			cnt2 := dp2[u]
			if u == 0 {
				res = max(res, float64(cnt2)/float64(dp1[n-1]))
			} else if u == n-1 {
				res = max(res, float64(cnt1)/float64(dp1[n-1]))
			} else {
				res = max(res, float64(cnt1)*2/float64(dp1[n-1])*float64(cnt2))
			}
		}
	}

	return res
}

const inf = 1 << 60

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
