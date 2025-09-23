package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(cur)
		buf.WriteByte('\n')
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}

func drive(reader *bufio.Reader) (ops []string, res []string) {
	n := readNum(reader)
	ops = make([]string, n)
	for i := range n {
		ops[i] = readString(reader)
	}
	res = solve(ops)
	return
}

func solve(ops []string) []string {
	var pq IntHeap

	var res []string

	for _, cur := range ops {
		if strings.HasPrefix(cur, "insert") {
			var x int
			readInt([]byte(cur), 7, &x)
			heap.Push(&pq, x)
		} else if strings.HasPrefix(cur, "getMin") {
			var x int
			readInt([]byte(cur), 7, &x)
			for pq.Len() > 0 && pq[0] < x {
				// 只要存在比x小的，必须删除掉
				res = append(res, "removeMin")
				heap.Pop(&pq)
			}

			if pq.Len() == 0 || pq[0] != x {
				res = append(res, fmt.Sprintf("insert %d", x))
				heap.Push(&pq, x)
			}
		} else {
			if pq.Len() == 0 {
				res = append(res, "insert 1")
				heap.Push(&pq, 1)
			}
			heap.Pop(&pq)
		}
		res = append(res, cur)
	}

	return res
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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
