package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, arr := range res {
		buf.WriteString(fmt.Sprintf("%d ", len(arr)))
		for _, u := range arr {
			buf.WriteString(fmt.Sprintf("%d ", u))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) [][]int {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) [][]int {

	g := NewGraph(n, 2*len(edges))
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	pq := make(PriorityQueue, n)
	items := make([]*Item, n)
	for i := range n {
		it := new(Item)
		it.id = i
		it.priority = 0
		it.index = i
		items[i] = it
		pq[i] = it
	}

	heap.Init(&pq)

	var res [][]int
	vis := make([]bool, n)

	for pq.Len() > 0 {
		var arr []int

		for pq.Len() > 0 {
			if len(arr) > 0 && pq[0].priority == len(arr) {
				// 当前节点的分离度正好是当前的group的size
				break
			}
			it := heap.Pop(&pq).(*Item)
			u := it.id
			arr = append(arr, u+1)
			vis[u] = true
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if !vis[v] {
					items[v].priority++
					heap.Fix(&pq, items[v].index)
				}
			}
		}

		res = append(res, arr)
		// 要把后面的那些的更新回来
		for j := range arr {
			u := arr[j] - 1
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if !vis[v] {
					items[v].priority--
					heap.Fix(&pq, items[v].index)
				}
			}
		}
	}

	return res
}

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

type Item struct {
	id       int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
