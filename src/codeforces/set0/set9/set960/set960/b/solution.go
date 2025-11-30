package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) int {
	var n, k1, k2 int
	fmt.Fscan(reader, &n, &k1, &k2)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(a, b, k1, k2)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, b []int, k1 int, k2 int) int {
	k := k1 + k2
	n := len(a)
	if n == 1 {
		diff := abs(a[0] - b[0])
		if diff >= k {
			diff -= k
			return diff * diff
		}
		// diff < k
		k -= diff
		diff = 0
		if k&1 == 1 {
			diff++
		}
		return diff * diff
	}
	var pq IntHeap
	for i := range n {
		heap.Push(&pq, abs(a[i]-b[i]))
	}

	for k > 0 && pq.Len() > 0 && pq[0] > 0 {
		x := heap.Pop(&pq).(int)
		y := pq[0]
		z := max(x-y, 1)
		if k >= z {
			k -= z
			heap.Push(&pq, x-z)
			continue
		}
		x -= k
		k = 0
		heap.Push(&pq, x)
		break
	}

	if k > 0 {
		k &= 1
		return k * k
	}

	var res int
	for pq.Len() > 0 {
		x := heap.Pop(&pq).(int)
		res += x * x
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
