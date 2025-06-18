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
	n, k := readTwoNums(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, k, edges)
}

func solve(n int, k int, edges [][]int) int {
	if n == 2 {
		// k == 1
		return 1
	}
	if k == 1 {
		return n - 1
	}
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	items := make([]*Item, n)

	for i := range n {
		items[i] = new(Item)
		items[i].id = i
		items[i].priority = 0
		items[i].index = -1
	}
	// 先把叶子节点加进去
	marked := make([]bool, n)
	for u := range n {
		if len(g[u]) == 1 {
			// u是一个叶子节点
			marked[u] = true
			v := g[u][0]
			items[v].priority++
		}
	}

	var pq PriorityQueue

	for u := range n {
		if items[u].priority >= k {
			heap.Push(&pq, items[u])
		}
	}

	sz := make([]int, n)

	var res int
	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.priority < k {
			break
		}
		x := it.priority / k
		u := it.id
		sz[u] += x * k
		res += x
		it.priority %= k

		if sz[u]+1 == len(g[u]) {
			// 找到它的parent
			// u成了一个叶子节点
			marked[u] = true
			for _, v := range g[u] {
				if !marked[v] {
					items[v].priority++
					if items[v].priority >= k {
						if items[v].index < 0 {
							heap.Push(&pq, items[v])
						} else {
							heap.Fix(&pq, items[v].index)
						}
					}
					break
				}
			}
		}
	}

	return res
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
	return pq[i].priority > pq[j].priority
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
