package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for _, cur := range res {
		fmt.Fprintln(writer, cur)
	}
}

func drive(reader *bufio.Reader) []string {
	var n, m, s, t int
	fmt.Fscan(reader, &n, &m, &s, &t)
	ways := make([][]int, m)
	for i := range m {
		ways[i] = make([]int, 3)
		fmt.Fscan(reader, &ways[i][0], &ways[i][1], &ways[i][2])
	}
	return solve(n, m, ways, s, t)
}

const inf = 1 << 60

func solve(n int, m int, ways [][]int, s int, t int) []string {
	s--
	t--
	g := NewGraph(n, m)
	gr := NewGraph(n, m)
	for i, cur := range ways {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v, i)
		gr.AddEdge(v, u, i)
	}

	bfs := func(s int, g *Graph) []*Item {
		dist := make([]*Item, n)
		pq := make(PriorityQueue, n)
		for i := range n {
			it := new(Item)
			it.id = i
			it.priority = inf
			it.index = i
			dist[i] = it
			pq[i] = it
		}

		dist[s].priority = 0
		heap.Init(&pq)

		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			if it.priority == inf {
				break
			}
			u := it.id
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				j := g.val[i]
				w := ways[j][2]
				if it.priority+w < dist[v].priority {
					pq.update(dist[v], it.priority+w)
				}
			}
		}

		return dist
	}

	d1 := bfs(s, g)
	d2 := bfs(t, gr)

	tr := NewGraph(n, 2*m)

	for i, cur := range ways {
		u, v, w := cur[0]-1, cur[1]-1, cur[2]
		if d1[u].priority < inf && d1[u].priority+w+d2[v].priority == d1[t].priority {
			tr.AddEdge(u, v, i)
			tr.AddEdge(v, u, i)
		}
	}

	bridges := findBridges(tr, s, n, m)

	ans := make([]string, m)
	for _, i := range bridges {
		ans[i] = "YES"
	}

	for i, cur := range ways {
		if ans[i] == "" {
			u, v, w := cur[0]-1, cur[1]-1, cur[2]
			// u is closer to s
			if d1[u].priority == inf || d2[v].priority == inf {
				continue
			}
			e := d1[t].priority - d1[u].priority - d2[v].priority
			x := w - (e - 1)
			if e > 1 && x >= 1 {
				ans[i] = fmt.Sprintf("CAN %d", x)
			}
		}
	}

	for i := range m {
		if ans[i] == "" {
			ans[i] = "NO"
		}
	}

	return ans
}

func findBridges(g *Graph, s int, n int, m int) []int {

	marked := make([]bool, m)

	dfn := make([]int, n)
	low := make([]int, n)
	var timer int
	stack := make([]int, n)
	var top int
	vis := make([]bool, n)

	var res []int

	var dfs func(u int)
	dfs = func(u int) {
		timer++
		dfn[u] = timer
		low[u] = timer
		vis[u] = true
		stack[top] = u
		top++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			j := g.val[i]
			if !marked[j] {
				marked[j] = true
				if dfn[v] == 0 {
					dfs(v)
					if low[v] == dfn[v] {
						// bridge
						res = append(res, j)
					}
					low[u] = min(low[u], low[v])
				} else if vis[v] {
					low[u] = min(low[u], dfn[v])
				}
				marked[j] = false
			}

		}
		if low[u] == dfn[u] {
			for top > 0 {
				v := stack[top-1]
				top--
				vis[v] = false
				if v == u {
					break
				}
			}
		}
	}

	dfs(s)

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

type Item struct {
	id       int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[:n-1]
	it.index = -1
	return it
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
