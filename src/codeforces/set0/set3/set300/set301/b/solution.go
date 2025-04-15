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
	n, d := readTwoNums(reader)
	a := readNNums(reader, n-2)
	stations := make([][]int, n)
	for i := range n {
		stations[i] = readNNums(reader, 2)
	}
	return solve(n, d, a, stations)
}

const inf = 1e9

func solve(n int, d int, a []int, stations [][]int) int {
	// binary search
	items := make([]*Item, n)
	for i := range n {
		it := new(Item)
		it.id = i
		it.priority = -inf
		items[i] = it
	}
	check := func(x int) bool {
		pq := make(PriorityQueue, n)
		for i := range n {
			items[i].priority = -inf
			items[i].index = i
			pq[i] = items[i]
		}
		items[0].priority = x
		heap.Init(&pq)
		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			if it.priority < 0 || it.id == n-1 {
				break
			}
			u := it.id
			for i := range n {
				if i == u {
					continue
				}
				tmp := distance(stations[u], stations[i])
				var add int
				if i < n-1 && i > 0 {
					add = a[i-1]
				}
				if it.priority-tmp*d >= 0 && it.priority-tmp*d+add > items[i].priority && items[i].index >= 0 {
					pq.update(items[i], it.priority-tmp*d+add)
				}
			}
		}
		return items[n-1].priority >= 0
	}

	return sort.Search(inf, check)
}

func distance(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(a int) int {
	return max(a, -a)
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
