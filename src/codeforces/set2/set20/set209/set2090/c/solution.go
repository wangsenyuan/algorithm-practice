package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
		}
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) [][]int {
	n := readNum(reader)
	t := readNNums(reader, n)
	return solve(t)
}

func solve(t []int) [][]int {
	// t[i] = 0, 肯定是在线 x + y = c的线上
	theLine := 2

	var pq PQ

	n := len(t)
	ans := make([][]int, n)

	type pair struct {
		first  int
		second int
	}

	cnt := make(map[int]int)

	items := make(map[pair]*Item)

	add := func(x int, y int) {
		it := new(Item)
		it.x = x
		it.y = y
		it.priority = x + y
		heap.Push(&pq, it)
		items[pair{x, y}] = it
	}

	for i, v := range t {
		for cnt[theLine] == (theLine+1)/3 {
			theLine += 3
		}

		if _, ok := cnt[theLine]; !ok {
			for x := 1; x < theLine; x += 3 {
				y := theLine - x
				add(x, y)
			}

			cnt[theLine] = 0
		}

		if v == 1 {
			// 只需要最近可用的
			it := heap.Pop(&pq).(*Item)
			x, y := it.x, it.y
			if it.cnt == 0 {
				cnt[x+y]++
			}
			ans[i] = []int{x, y}
			it.cnt++
			if it.cnt == 1 {
				it.y++
				it.priority++
			} else if it.cnt == 2 {
				it.x++
				it.y--
			} else if it.cnt == 3 {
				it.y++
				it.priority += 3
			}

			if it.cnt < 4 {
				heap.Push(&pq, it)
			}

		} else {
			x := 3*cnt[theLine] + 1
			y := theLine - x
			ans[i] = []int{x, y}
			cnt[x+y]++

			it := items[pair{x, y}]
			it.cnt++
			it.y++
			it.priority = x + y + 1
			heap.Fix(&pq, it.index)
		}
	}

	return ans
}

type Item struct {
	x        int
	y        int
	priority int
	cnt      int
	index    int
}

func (a *Item) Less(b *Item) bool {
	if a.priority != b.priority {
		return a.priority < b.priority
	}
	if a.x != b.x {
		return a.x < b.x
	}
	return a.y < b.y
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].Less(pq[j])
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(*pq)
	item := old[n-1]
	item.index = -1
	*pq = old[:n-1]
	return item
}
