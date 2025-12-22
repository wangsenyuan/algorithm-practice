package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		line := readNNums(reader, 4)
		n, m, k, p := line[0], line[1], line[2], line[3]
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			edges[i] = readNNums(reader, 3)
		}
		res := solve(n, k, p, edges)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

const inf = 0x3f3f3f3f

func solve(n int, k int, p int, roads [][]int) int {
	add := func(a, b int) int {
		return (a + b) % p
	}

	g := NewGraph(n, len(roads))
	r := NewGraph(n, len(roads))
	deg := make([]int, n)
	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		g.AddEdge(u-1, v-1, w)
		r.AddEdge(v-1, u-1, w)
		if w == 0 {
			deg[v-1]++
		}
	}

	dist1 := getDist(n, g, 0)
	distn := getDist(n, r, n-1)

	que := make([]int, n)
	var head, tail int
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			que[head] = i
			head++
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if g.val[i] == 0 {
				deg[v]--
				if deg[v] == 0 {
					que[head] = v
					head++
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if deg[i] > 0 && dist1[i]+distn[i] <= dist1[n-1]+k {
			return -1
		}
	}

	type item struct {
		id   int
		dist int
		ord  int
	}

	pos := make([]int, n)
	for i := 0; i < head; i++ {
		pos[que[i]] = i
	}

	items := make([]item, n)
	for i := 0; i < n; i++ {
		items[i] = item{i, dist1[i], pos[i]}
	}

	slices.SortFunc(items, func(a, b item) int {
		if a.dist != b.dist {
			return a.dist - b.dist
		}
		return a.ord - b.ord
	})

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
	}

	dp[0][0] = 1

	for j := 0; j <= k; j++ {
		for _, it := range items {
			u := it.id
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				w := g.val[i]
				nj := dist1[u] + w + j - dist1[v]
				if nj <= k {
					dp[v][nj] = add(dp[v][nj], dp[u][j])
				}
			}
		}
	}

	var res int
	for j := 0; j <= k; j++ {
		res = add(res, dp[n-1][j])
	}
	return res
}

func getDist(n int, g *Graph, src int) []int {
	items := make([]*Node, n)
	pq := make(PQ, n)

	for i := 0; i < n; i++ {
		items[i] = NewNode(i, inf)
		items[i].index = i
		pq[i] = items[i]
	}

	heap.Init(&pq)

	pq.update(items[src], 0)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Node)
		u := it.value
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if items[v].priority > items[u].priority+g.val[i] {
				pq.update(items[v], items[u].priority+g.val[i])
			}
		}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = items[i].priority
	}
	return res
}

func solve1(n int, k int, p int, roads [][]int) int {
	add := func(a, b int) int {
		return (a + b) % p
	}

	g := NewGraph(n, len(roads))
	r := NewGraph(n, len(roads))
	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		g.AddEdge(u-1, v-1, w)
		r.AddEdge(v-1, u-1, w)
	}

	dist := getDist(n, g, 0)

	vis := make([][]bool, n)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, k+1)
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(u, x int) bool
	dfs = func(u, x int) bool {
		if vis[u][x] {
			return false
		}
		if dp[u][x] >= 0 {
			return true
		}
		var ans int
		vis[u][x] = true
		for i := r.nodes[u]; i > 0; i = r.next[i] {
			v := r.to[i]
			w := r.val[i]
			nx := dist[u] + x - (dist[v] + w)
			// v is closer to 0
			// dist[v] + w + nx = dist[u] + x
			if nx >= 0 && nx <= k {
				if !dfs(v, nx) {
					return false
				}
				ans = add(ans, dp[v][nx])
			}
		}
		dp[u][x] = ans
		vis[u][x] = false
		return true
	}

	// 0在0环中
	if !dfs(0, 0) {
		return -1
	}

	var res int

	dp[0][0] = 1
	for x := k; x >= 0; x-- {
		if !dfs(n-1, x) {
			return -1
		}
		res = add(res, dp[n-1][x])
	}

	return res
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
	next := make([]int, e+10)
	to := make([]int, e+10)
	val := make([]int, e+10)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

type Node struct {
	value    int
	priority int
	index    int
}

func NewNode(v int, p int) *Node {
	n := new(Node)
	n.value = v
	n.priority = p
	n.index = -1
	return n
}

type PQ []*Node

func (this PQ) Len() int {
	return len(this)
}

func (this PQ) Less(i, j int) bool {
	return this[i].priority < this[j].priority
}

func (this PQ) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *PQ) Push(x any) {
	i := x.(*Node)
	i.index = len(*this)
	*this = append(*this, i)
}

func (this *PQ) Pop() any {
	old := *this
	n := len(old)
	ret := old[n-1]
	ret.index = -1
	*this = old[:n-1]
	return ret
}

func (this *PQ) update(n *Node, v int) {
	n.priority = v
	heap.Fix(this, n.index)
}
