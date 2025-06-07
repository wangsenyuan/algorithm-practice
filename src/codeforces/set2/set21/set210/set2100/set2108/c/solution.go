package main

import (
	"bufio"
	"bytes"
	"cmp"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for range tc {
		buf.WriteString(fmt.Sprintf("%d\n", process(reader)))
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a []int) int {
	n := len(a)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(b.first-a.first, a.second-b.second)
	})

	var active PriorityQueue

	create := func(p pair) *Item {
		it := new(Item)
		it.id = p.second
		it.priority = p.first
		return it
	}

	var res int
	marked := make([]bool, n)
	for k := 0; k < n; {
		p := arr[k]
		if marked[p.second] {
			k++
			continue
		}
		if active.Len() == 0 || active[0].priority < p.first {
			heap.Push(&active, create(p))
			marked[p.second] = true
			res++
			k++
		}

		it := heap.Pop(&active).(*Item)
		i := it.id
		if i > 0 && !marked[i-1] {
			heap.Push(&active, create(pair{a[i-1], i - 1}))
			marked[i-1] = true
		}
		if i < n-1 && !marked[i+1] {
			heap.Push(&active, create(pair{a[i+1], i + 1}))
			marked[i+1] = true
		}
	}

	return res
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}
