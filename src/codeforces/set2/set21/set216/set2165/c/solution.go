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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		for _, x := range res {
			fmt.Fprintln(writer, x)
		}
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	c := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &c[i])
	}
	return solve(a, c)
}

const H = 30

type pair struct {
	first  int
	second int
}

func maxPair(a, b pair) pair {
	if a.first > b.first || a.first == b.first && a.second > b.second {
		return a
	}
	return b
}

func solve(a []int, c []int) []int {
	// n := len(a)
	slices.Sort(a)
	if len(a) > H {
		// 只需要最大的H个数
		a = a[len(a)-H:]
	}

	find := func(c int) int {
		var pq IntHeap
		for _, v := range a {
			heap.Push(&pq, v)
		}
		var todo []int
		var res int
		for i := H - 1; i >= 0; i-- {
			if (c>>i)&1 == 0 {
				continue
			}
			cur := heap.Pop(&pq).(int)
			todo = append(todo, cur)
			if cur >= c {
				break
			}

			if len(pq) == 0 {
				res += c - cur
				break
			}

			if cur&(1<<i) == 0 {
				res += 1<<i - cur
			} else {
				// 把去掉最高位的，重新返回去
				cur ^= 1 << i
				heap.Push(&pq, cur)
			}
			c ^= 1 << i

		}

		return res
	}

	ans := make([]int, len(c))

	for i, x := range c {
		ans[i] = find(x)
	}

	return ans
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
