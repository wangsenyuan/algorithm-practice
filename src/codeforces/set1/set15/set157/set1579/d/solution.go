package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type pair struct {
	x int
	y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for i := 0; i < tc; i++ {
		ans := drive(reader)
		fmt.Fprintf(writer, "%d\n", len(ans))
		for _, p := range ans {
			fmt.Fprintf(writer, "%d %d\n", p.x, p.y)
		}
	}
}

func drive(reader *bufio.Reader) []pair {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)

}

func readNum(reader *bufio.Reader) int {
	var x int
	_, _ = fmt.Fscan(reader, &x)
	return x
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = readNum(reader)
	}
	return res
}

func solve(a []int) []pair {
	var res []pair
	pq := make(pq, len(a))
	for i, v := range a {
		pq[i] = person{i + 1, v}
	}

	heap.Init(&pq)

	for len(pq) > 1 {
		p0 := heap.Pop(&pq).(person)
		p1 := heap.Pop(&pq).(person)
		if p1.val == 0 {
			break
		}
		w := p1.val
		if len(pq) > 0 && pq[0].val > 0 {
			w = p1.val - pq[0].val + 1
		}
		for range w {
			res = append(res, pair{p0.id, p1.id})
		}
		p0.val -= w
		p1.val -= w
		if p0.val > 0 {
			heap.Push(&pq, p0)
		}
		if p1.val > 0 {
			heap.Push(&pq, p1)
		}
	}

	return res
}

type person struct {
	id  int
	val int
}

type pq []person

func (pq pq) Len() int {
	return len(pq)
}

func (pq pq) Less(i, j int) bool {
	return pq[i].val > pq[j].val
}

func (pq pq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pq) Push(x any) {
	*pq = append(*pq, x.(person))
}

func (pq *pq) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
