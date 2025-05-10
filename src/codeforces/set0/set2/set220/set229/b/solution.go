package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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
	n, m := readTwoNums(reader)
	gates := make([][]int, m)
	for i := range m {
		gates[i] = readNNums(reader, 3)
	}
	travellers := make([][]int, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')
		var k int
		pos := readInt(s, 0, &k) + 1
		travellers[i] = make([]int, k)
		for j := range k {
			pos = readInt(s, pos, &travellers[i][j]) + 1
		}
	}
	return solve(n, gates, travellers)
}

const inf = 1 << 60

func solve(n int, gates [][]int, travellers [][]int) int {
	// 始终计算到到一个星球的最早时间
	// 然后计算什么时候可以离开它（使用gate）
	g := NewGraph(n, 2*len(gates))
	for _, gate := range gates {
		u, v, w := gate[0]-1, gate[1]-1, gate[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	arr := make([][]pair, n)
	for i, cur := range travellers {
		if len(cur) == 0 {
			continue
		}
		for j := 0; j < len(cur); {
			k := j
			for j < len(cur) && cur[j]-cur[k] == j-k {
				j++
			}
			arr[i] = append(arr[i], pair{cur[k], cur[j-1]})
		}
	}

	dist := make([]*Item, n)
	pq := make(PriorityQueue, n)
	for i := range n {
		it := new(Item)
		it.id = i
		it.priority = inf
		it.index = i
		it.priority = inf
		dist[i] = it
		pq[i] = it
	}
	dist[0].priority = 0
	if len(arr[0]) > 0 && arr[0][0].first == 0 {
		dist[0].priority = arr[0][0].second + 1
	}

	heap.Init(&pq)

	res := inf

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.priority == inf {
			break
		}
		u := it.id

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v].index < 0 {
				continue
			}
			w := g.val[i]
			x := it.priority + w
			if v == n-1 {
				res = min(res, x)
			}
			j := sort.Search(len(arr[v]), func(j int) bool {
				return arr[v][j].first > x
			})
			j--
			// arr[v]j].first <= x
			if j >= 0 && arr[v][j].second >= x {
				x = arr[v][j].second + 1
			}
			if dist[v].priority > x {
				// x是可以从v离开的最早时间
				pq.update(dist[v], x)
			}
		}
	}

	if res == inf {
		return -1
	}

	return res
}

type pair struct {
	first  int
	second int
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
