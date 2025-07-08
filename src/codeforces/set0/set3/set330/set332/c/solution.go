package main

import (
	"bufio"
	"cmp"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := process(reader)
	s := fmt.Sprintf("%v", res)
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

func process(reader *bufio.Reader) ([]int, [][]int, int) {
	n, p, k := readThreeNums(reader)
	orders := make([][]int, n)
	for i := range n {
		orders[i] = readNNums(reader, 2)
	}
	return solve(n, p, k, orders), orders, k
}

func solve(n int, p int, k int, orders [][]int) []int {
	type order struct {
		id int
		a  int
		b  int
	}

	arr := make([]order, n)
	for i := range n {
		arr[i] = order{
			i + 1, orders[i][0], orders[i][1],
		}
	}

	slices.SortFunc(arr, func(u order, v order) int {
		return cmp.Or(u.b-v.b, v.a-u.a)
	})

	sum := make([]int, n+1)
	for i, x := range arr {
		sum[i+1] = sum[i] + x.b
	}

	var pq PriorityQueue

	var data struct {
		pos   int
		sum_b int
		sum_a int
	}

	var sum2 int
	for i := n - 1; i >= p-k; i-- {
		sum2 += arr[i].a
		heap.Push(&pq, Item{arr[i].id, arr[i].a, arr[i].b})
		if pq.Len() > k {
			it := heap.Pop(&pq).(Item)
			sum2 -= it.a
		}
		if pq.Len() == k {
			if data.sum_a < sum2 || data.sum_a == sum2 && data.sum_b < sum[i]-sum[i-(p-k)] {
				data.pos = i
				data.sum_a = sum2
				data.sum_b = sum[i] - sum[i-(p-k)]
			}
		}
	}
	var ans []int

	for pq.Len() > 0 {
		ans = append(ans, heap.Pop(&pq).(Item).id)
	}

	for j := data.pos - (p - k); j < data.pos; j++ {
		ans = append(ans, arr[j].id)
	}

	return ans
}

type Item struct {
	id int
	a  int
	b  int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].a != pq[j].a {
		return pq[i].a < pq[j].a
	}
	return pq[i].b < pq[j].b
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
