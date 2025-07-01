package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	assign, _, _, _ := process(reader)
	if len(assign) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, x := range assign {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (assign []int, a []int, b []int, c []int) {
	n, m, s := readThreeNums(reader)
	a = readNNums(reader, m)
	b = readNNums(reader, n)
	c = readNNums(reader, n)
	assign = solve(a, b, c, s)
	return
}

func solve(a []int, b []int, c []int, s int) []int {
	type bug struct {
		id         int
		complexity int
	}
	type student struct {
		id      int
		ability int
		cost    int
	}
	m := len(a)
	bugs := make([]bug, m)

	for i, x := range a {
		bugs[i] = bug{id: i, complexity: x}
	}

	slices.SortFunc(bugs, func(a, b bug) int {
		return b.complexity - a.complexity
	})

	n := len(b)

	students := make([]student, n)
	for i, x := range b {
		students[i] = student{id: i, ability: x, cost: c[i]}
	}

	slices.SortFunc(students, func(a, b student) int {
		return b.ability - a.ability
	})

	assign := make([]int, m)

	items := make([]*Item, n)
	for i := range n {
		it := new(Item)
		it.id = students[i].id
		it.priority = students[i].cost
		items[i] = it
	}

	check := func(days int) bool {
		// 当一个学生，开始接受任务，只要队列没有满，就可以一直接受，不增加cost
		// 所以这种情况下，似乎应该选择cost最低的学生
		if days == 0 {
			return false
		}

		var sum int
		var j int
		var pq PriorityQueue
		for i := 0; i < m; {
			for j < n && students[j].ability >= bugs[i].complexity {
				heap.Push(&pq, items[j])
				j++
			}
			if len(pq) == 0 {
				return false
			}
			it := heap.Pop(&pq).(*Item)
			u := it.id + 1
			sum += it.priority
			for v := 0; v < days && i+v < m; v++ {
				assign[bugs[i+v].id] = u
			}

			i += days
		}

		return sum <= s
	}

	if !check(m) {
		// 每天一个任务，那么就不会有队列满的情况，如果还无法完成，就无解了
		return nil
	}

	days := sort.Search(m, check)

	check(days)

	return assign
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
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[:n-1]
	it.index = -1
	return it
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
