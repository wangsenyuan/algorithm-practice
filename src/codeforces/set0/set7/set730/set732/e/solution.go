package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res, a, b := drive(reader)
	fmt.Println(res[0], res[1])
	s := fmt.Sprintf("%v", a)
	fmt.Println(s[1 : len(s)-1])
	s = fmt.Sprintf("%v", b)
	fmt.Println(s[1 : len(s)-1])
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) (p []int, s []int, res []int, a []int, assign []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	p = readNNums(reader, n)
	s = readNNums(reader, m)
	res, a, assign = solve(p, s)
	return
}

type computer struct {
	id int
	pw int
}

type socket struct {
	id int
	pw int
}

func solve(p []int, s []int) ([]int, []int, []int) {
	// 有些s最终会变成相同的值
	arr := make([]computer, len(p))
	for i, v := range p {
		arr[i] = computer{i, v}
	}

	slices.SortFunc(arr, func(a, b computer) int {
		return a.pw - b.pw
	})

	arr2 := make([]socket, len(s))

	for i, v := range s {
		arr2[i] = socket{i, v}
	}

	slices.SortFunc(arr2, func(a, b socket) int {
		return a.pw - b.pw
	})

	res := make([]int, 2)
	a := make([]int, len(s))
	b := make([]int, len(p))

	marked := make([]bool, len(arr2))

	for d := range 31 {
		var l int
		for j, cur := range arr2 {
			for l < len(arr) && (b[arr[l].id] > 0 || arr[l].pw < cur.pw) {
				l++
			}
			if l < len(arr) && arr[l].pw == cur.pw && !marked[j] {
				b[arr[l].id] = cur.id + 1
				res[0]++
				res[1] += d
				a[cur.id] = d

				marked[j] = true
			}
		}
		for j := range len(arr2) {
			arr2[j].pw = (arr2[j].pw + 1) / 2
		}
	}

	return res, a, b
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority || pq[i].priority == pq[j].priority && pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
