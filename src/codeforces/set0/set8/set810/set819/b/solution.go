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
	_, res := drive(reader)
	fmt.Println(res[0], res[1])
}

func drive(reader *bufio.Reader) (p []int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	p = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	res = solve(slices.Clone(p))
	return
}

func solve(a []int) []int {
	var res int
	var gt, lt int
	n := len(a)
	todo := make([]int, n*2+1)
	for i := range n {
		a[i]--
		res += abs(i - a[i])
		if i >= a[i] {
			gt++
		} else {
			lt++
			todo[a[i]-i]++
		}
	}

	best := res
	var k int
	for i := 1; i < n; i++ {
		gt--
		res -= (n - 1 - a[n-i])
		res += a[n-i]
		res += gt - lt
		lt++
		todo[a[n-i]+i]++
		lt -= todo[i]
		gt += todo[i]
		if res < best {
			best = res
			k = i
		}
	}
	return []int{best, k}
}

func solve1(p []int) []int {
	n := len(p)
	sum := make([]int, n+2)
	diff := make([]int, n+2)

	add := func(l int, r int, k int, b int) {
		if l > r {
			return
		}
		sum[l] += b
		diff[l] += k

		sum[r+1] -= b + k*(r-l)
		diff[r] -= k
	}

	for i := range n {
		c1 := i + 1
		p1 := 0
		c2 := n
		p2 := p1 + c2 - c1
		c3 := i
		p3 := p2 + c3
		if p[i] <= c3 {
			add(p1, p2, 1, c1-p[i])
			add(p2+1, p2+p[i], -1, p[i]-1)
			add(p2+p[i]+1, p3, 1, 1)
		} else {
			add(p1, p1+p[i]-c1, -1, p[i]-c1)
			add(p1+p[i]-c1+1, p2, 1, 1)
			add(p2+1, p3, -1, p[i]-1)
		}
	}

	var curDiff int

	for i := range n {
		sum[i] += curDiff
		curDiff += diff[i]
	}

	res := make([]int, n)

	var curSum int
	for i := range n {
		curSum += sum[i]
		res[i] += curSum
	}

	var k int
	for i := range n {
		if res[i] < res[k] {
			k = i
		}
	}

	return []int{res[k], k}
}

func abs(num int) int {
	return max(num, -num)
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	// item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
