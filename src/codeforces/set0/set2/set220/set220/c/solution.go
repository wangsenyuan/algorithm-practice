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
	res := process(reader)
	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(a, b)
}

func solve(a []int, b []int) (ans []int) {
	a, b = b, a
	n := len(a)
	p1 := make([]int, n)
	p2 := make([]int, n)
	for i := range n {
		a[i]--
		b[i]--
		p1[a[i]] = i
		p2[b[i]] = i
	}

	pqs := make([]PQ, 2)
	items := make([]*Item, n)
	for i := range n {
		x := b[i]
		j := p1[x]
		it := new(Item)
		it.id = x
		if j <= i {
			it.priority = i - j
			heap.Push(&pqs[0], it)
		} else {
			it.priority = j - i
			heap.Push(&pqs[1], it)
		}
		items[x] = it
	}

	for k := range n {
		for pqs[1].Len() > 0 && pqs[1][0].priority == k {
			it := heap.Pop(&pqs[1]).(*Item)
			// 要把它移动到pqs[0]中
			it.priority = -k
			heap.Push(&pqs[0], it)
		}

		tmp := n
		if pqs[0].Len() > 0 {
			tmp = pqs[0][0].priority + k
		}
		if pqs[1].Len() > 0 {
			tmp = min(tmp, pqs[1][0].priority-k)
		}
		ans = append(ans, tmp)
		// 它肯定在负堆的， 要把它移动到正堆里面去
		x := a[k]
		pqs[0].update(items[x], -inf)
		heap.Pop(&pqs[0])
		j := p1[x] + n
		i := p2[x]
		// j - i > 0
		items[x].priority = j - i
		heap.Push(&pqs[1], items[x])
	}

	return ans
}

const inf = 1 << 60

type Item struct {
	id       int
	priority int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	res.index = -1
	return res
}

func (pq *PQ) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}

func bruteForce(a []int, b []int) []int {
	a, b = b, a
	n := len(a)

	pos := make([]int, n)

	for i := range n {
		a[i]--
		b[i]--
		pos[b[i]] = i
	}

	arr := make([]int, 2*n)
	copy(arr, a)
	copy(arr[n:], a)

	var ans []int

	for i := range n {
		tmp := n
		for j := range n {
			x := arr[i+j]
			k := pos[x]
			tmp = min(tmp, abs(j-k))
		}
		ans = append(ans, tmp)
	}
	return ans
}

func abs(num int) int {
	return max(num, -num)
}
