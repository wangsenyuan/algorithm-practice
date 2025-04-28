package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

const inf = 1 << 60

func solve(n int, edges [][]int) []int {
	// 叉居然是单向的边
	m := len(edges)
	g := NewGraph(n, m)
	r := NewGraph(n, m)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		r.AddEdge(v, u, i)
	}
	que := make([]int, n)
	bfs := func(x int, g *Graph) []int {
		dist := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = inf
		}
		var head, tail int
		dist[x] = 0
		que[head] = x
		head++
		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] == inf {
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}

		return dist
	}

	d1 := bfs(0, g)

	layers := make([][]int, n)

	for u := 0; u < n; u++ {
		if d1[u] < n {
			layers[d1[u]] = append(layers[d1[u]], u)
		}
	}

	dn := bfs(n-1, r)

	best := dn[0]

	check := func(eid int) int {
		if best == inf {
			// 无论如何都连不起来
			return -1
		}
		u, v := edges[eid][0]-1, edges[eid][1]-1
		// 删除从u到v的边
		if d1[u]+1 != d1[v] || d1[v]+dn[v] != best {
			// 这条边不在最优路线上
			return best
		}

		res := inf

		for _, w := range layers[d1[v]] {
			if w != v && dn[w] >= dn[v] {
				res = min(res, d1[w]+dn[w])
			}
		}

		// 寻找所有和v相连的边
		for i := r.nodes[v]; i > 0; i = r.next[i] {
			w := r.to[i]
			j := r.val[i]
			if j == eid || dn[v] >= dn[w] {
				// 这条边被删除了
				continue
			}
			res = min(res, d1[w]+dn[v]+1)
		}

		if res >= inf {
			return -1
		}

		return res
	}

	ans := make([]int, m)

	for i := 0; i < m; i++ {
		ans[i] = check(i)
	}

	return ans
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


func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}