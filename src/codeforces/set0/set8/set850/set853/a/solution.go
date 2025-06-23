package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res, d := process(reader)
	fmt.Println(res)
	s := fmt.Sprintf("%v", d)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) (k int, c []int, res int, d []int) {
	n, k := readTwoNums(reader)
	c = readNNums(reader, n)
	res, d = solve(n, k, c)
	return
}

func solve(n int, k int, c []int) (res int, d []int) {
	pq := make(PriorityQueue, 0, k+1)
	for i := range k {
		it := &Item{
			id:       i,
			priority: c[i],
		}
		heap.Push(&pq, it)
	}

	d = make([]int, n)
	for i := k; i < n; i++ {
		it := &Item{
			id:       i,
			priority: c[i],
		}
		heap.Push(&pq, it)
		best := heap.Pop(&pq).(*Item)
		j := best.id
		d[j] = i + 1
		res += (i - j) * c[j]
	}
	i := n
	for pq.Len() > 0 {
		best := heap.Pop(&pq).(*Item)
		j := best.id
		d[j] = i + 1
		res += (i - j) * c[j]
		i++
	}

	return
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
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	item := (*pq)[n-1]
	(*pq)[n-1] = nil
	*pq = (*pq)[:n-1]
	item.index = -1
	return item
}
