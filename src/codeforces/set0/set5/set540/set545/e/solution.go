package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	total, res := drive(reader)
	fmt.Println(total)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (total_weight int, res []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		for j := range 3 {
			fmt.Fscan(reader, &edges[i][j])
		}
	}
	var s int
	fmt.Fscan(reader, &s)
	return solve(n, edges, s)
}

const inf = 1 << 60

func solve(n int, edges [][]int, s int) (total_weight int, res []int) {
	g := NewGraph(n, 2*len(edges))
	for i, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}
	s--

	pq := make(PriorityQueue, n)
	items := make([]*Item, n)
	from := make([]int, n)

	for i := range n {
		it := new(Item)
		it.id = i
		it.priority = inf
		it.index = i
		items[i] = it
		pq[i] = it
		from[i] = -1
	}
	items[s].priority = 0
	heap.Init(&pq)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id
		if from[u] != -1 {
			res = append(res, from[u]+1)
			total_weight += edges[from[u]][2]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := edges[g.val[i]][2]
			if items[v].priority > it.priority+w || items[v].priority == it.priority+w && w < edges[from[v]][2] {
				from[v] = g.val[i]
				pq.update(items[v], it.priority+w)
			}
		}
	}
	return
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
