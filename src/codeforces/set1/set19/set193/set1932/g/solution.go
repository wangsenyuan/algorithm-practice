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
	n, m, H := readThreeNums(reader)
	L := readNNums(reader, n)
	S := readNNums(reader, n)
	passages := make([][]int, m)
	for i := 0; i < m; i++ {
		passages[i] = readNNums(reader, 2)
	}
	return solve(H, L, S, passages)
}

func solve(H int, L []int, S []int, passages [][]int) int {
	n := len(L)

	getNextTime := func(t0 int, u int, v int) int {
		t0 %= H
		pos_u := L[u] + S[u]*t0
		pos_v := L[v] + S[v]*t0
		a := ((pos_u-pos_v)%H + H) % H
		b := ((S[v]-S[u])%H + H) % H
		g, x, _ := extgcd(b, H)
		if a%g != 0 {
			return -1
		}

		x *= a / g
		x %= H / g
		if x < 0 {
			x += H / g
		}
		return x
	}

	g := NewGraph(n, len(passages)*2)

	for _, cur := range passages {
		u, v := cur[0], cur[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	items := make([]*Item, n)
	pq := make(PriorityQueue, n)

	for i := range n {
		it := new(Item)
		it.id = i
		it.priority = inf
		it.index = i
		items[i] = it
		pq[i] = it
	}
	items[0].priority = 0

	heap.Init(&pq)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.priority >= inf {
			break
		}
		if it.id == n-1 {
			return it.priority
		}
		u := it.id
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dt := getNextTime(it.priority, u, v)
			if dt < 0 {
				continue
			}
			if it.priority+dt+1 < items[v].priority {
				pq.update(items[v], it.priority+dt+1)
			}
		}
	}

	return -1
}

func abs(num int) int {
	return max(num, -num)
}

const inf = 1 << 60

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

func extgcd(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	k := a / b
	g, x, y := extgcd(b, a%b)
	return g, y, x - k*y
}
