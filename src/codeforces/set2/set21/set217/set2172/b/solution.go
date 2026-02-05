package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for _, x := range res {
		fmt.Fprintf(writer, "%.10f\n", x)
	}
}

func drive(reader *bufio.Reader) []float64 {
	var n, m, l, x, y int
	fmt.Fscan(reader, &n, &m, &l, &x, &y)
	buses := make([][]int, n)
	for i := range n {
		buses[i] = make([]int, 2)
		fmt.Fscan(reader, &buses[i][0], &buses[i][1])
	}
	p := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &p[i])
	}
	return solve(l, x, y, buses, p)
}

type person struct {
	id  int
	pos int
}

type bus struct {
	id    int
	start int
	end   int
}

func solve(l int, x int, y int, buses [][]int, p []int) []float64 {
	m := len(p)
	people := make([]person, m)

	for i, v := range p {
		people[i] = person{id: i, pos: v}
	}

	n := len(buses)
	arr := make([]bus, n)
	for i, cur := range buses {
		arr[i] = bus{id: i, start: cur[0], end: cur[1]}
	}

	slices.SortFunc(people, func(a person, b person) int {
		return a.pos - b.pos
	})

	slices.SortFunc(arr, func(a bus, b bus) int {
		return a.start - b.start
	})
	var pq PriorityQueue

	z := x - y

	ans := make([]float64, m)
	for _, cur := range people {
		for len(arr) > 0 && arr[0].start <= cur.pos {
			// s + (e - s) * z / x 最大
			it := new(Item)
			it.id = arr[0].id
			it.priority = x*arr[0].start + z*(arr[0].end-arr[0].start)
			heap.Push(&pq, it)
			arr = arr[1:]
		}

		tmp := cur.pos

		// 确实不是e最大的最好，
		for pq.Len() > 0 {
			id := pq[0].id
			s, e := buses[id][0], buses[id][1]
			if float64(s)+float64(e-s)*float64(z)/float64(x) > float64(cur.pos) {
				break
			}
			heap.Pop(&pq)
		}

		if pq.Len() > 0 {
			id := pq[0].id
			s, e := buses[id][0], buses[id][1]

			ans[cur.id] = float64(e-s) / float64(x)
			tmp = e
		}
		ans[cur.id] += float64(l-tmp) / float64(y)
	}

	return ans
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
	index    int
}

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
