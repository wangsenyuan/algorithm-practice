package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	n, m, k := readThreeNums(reader)
	E := make([][]int, m)

	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}

	res := solve(n, k, E)

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const INF = 1e15

func solve(n int, k int, E [][]int) []int {

	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	dist := make([]int, n)

	for i := 1; i < n; i++ {
		dist[i] = INF
	}
	items := make([]*Item, n)
	for i := range n {
		items[i] = new(Item)
	}
	pq := make(PQ, 0, n)

	dijkstra := func() {
		for i := range n {
			items[i].id = i
			items[i].priority = dist[i]
			items[i].index = i
		}

		pq = append(pq, items...)

		heap.Init(&pq)

		for pq.Len() > 0 {
			u := heap.Pop(&pq).(*Item)
			for i := g.node[u.id]; i > 0; i = g.next[i] {
				v := g.to[i]
				w := g.val[i]
				if dist[v] > dist[u.id]+w {
					dist[v] = dist[u.id] + w
					pq.update(items[v], dist[v])
				}
			}
		}
	}

	dijkstra()

	for range k {
		q := []vec{}
		for i, d := range dist {
			v := vec{i, i*i + d}
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).detCmp(v.sub(q[len(q)-1])) <= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, v)
		}
		for i := range dist {
			p := vec{-2 * i, 1}
			for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
				q = q[1:]
			}
			dist[i] = p.dot(q[0]) + i*i
		}
		dijkstra()
	}

	return dist
}

type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }
func (a vec) detCmp(b vec) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

type Item struct {
	id       int
	priority int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	i := x.(*Item)
	i.index = len(*pq)
	*pq = append(*pq, i)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	rt := old[n-1]
	*pq = old[:n-1]
	return rt
}

func (pq *PQ) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

type Graph struct {
	node []int
	next []int
	to   []int
	val  []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.val = make([]int, e+1)
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
