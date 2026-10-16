package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	var buf bytes.Buffer
	for _, x := range ans {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func process(reader *bufio.Reader) []int {
	tc := readNum(reader)
	ans := make([]int, tc)
	for i := 0; i < tc; i++ {
		n := readNum(reader)
		jobs := make([][]int, n)
		for i := 0; i < n; i++ {
			jobs[i] = readNNums(reader, 2)
		}
		ans[i] = solve(jobs)
	}

	return ans
}

func solve(jobs [][]int) int {
	n := len(jobs)
	slices.SortFunc(jobs, func(a, b []int) int {
		return (a[0] + a[1]) - (b[0] + b[1])
	})
	var sum int
	var res int

	pq := make(IntHeap, 0, n)
	for i := 0; i < n; i++ {
		heap.Push(&pq, jobs[i][1])
		sum += jobs[i][1]
		res++
		if sum-jobs[i][1] > jobs[i][0] {
			res--
			sum -= heap.Pop(&pq).(int)
		}
	}

	return res
}

// An IntHeap is a min-heap of ints.
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
