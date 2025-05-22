package main

import (
	"bufio"
	"bytes"
	"cmp"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {

	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)

	_, _, xp, leave := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", xp))
	s := fmt.Sprintf("%v", leave)
	buf.WriteString(s[1 : len(s)-1])
	buf.WriteByte('\n')
	buf.WriteTo(w)
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

func process(reader *bufio.Reader) (tasks [][]int, T int, xp int, leave []int) {
	n := readNum(reader)
	tasks = make([][]int, n)
	for i := range n {
		tasks[i] = readNNums(reader, 3)
	}
	T = readNum(reader)
	xp, leave = solve(tasks, T)
	return
}

type task struct {
	id       int
	arrive   int
	pages    int
	priority int
}

const inf = 1 << 60

func solve(tasks [][]int, T int) (int, []int) {
	n := len(tasks)
	ts := make([]task, n)
	leave := make([]int, n)

	check := func(p int) bool {
		x := -1
		for i := range n {
			ts[i] = task{
				i, tasks[i][0], tasks[i][1], tasks[i][2],
			}
			if ts[i].priority == -1 {
				ts[i].priority = p
				x = i
			}
		}
		slices.SortFunc(ts, func(a, b task) int {
			return cmp.Or(a.arrive-b.arrive, b.priority-a.priority)
		})

		var pq PriorityQueue
		for i := 0; i < n; i++ {
			var next int = inf
			if i+1 < n {
				next = ts[i+1].arrive
			}
			// 这里的id用下标
			it := &Item{
				id:       i,
				priority: ts[i].priority,
			}
			now := ts[i].arrive
			heap.Push(&pq, it)
			for now < next && pq.Len() > 0 {
				tmp := pq[0]
				j := tmp.id
				if now+ts[j].pages <= next {
					now += ts[j].pages
					leave[ts[j].id] = now
					heap.Pop(&pq)
				} else {
					ts[j].pages -= next - now
					break
				}
			}
		}
		return leave[x] <= T
	}

	var ps []int
	//找到所有可用的priority
	for i := 0; i < n; i++ {
		if tasks[i][2] != -1 {
			ps = append(ps, tasks[i][2])
		}
	}
	// ps是不同的
	sort.Ints(ps)

	var arr []int
	if ps[0] != 1 {
		// 可以选择1
		arr = append(arr, 1)
	}
	for i := 0; i < len(ps); {
		j := i
		for i < len(ps) && ps[i]-i == ps[j]-j {
			i++
		}
		arr = append(arr, ps[i-1]+1)
	}

	ans := sort.Search(len(arr), func(i int) bool {
		return check(arr[i])
	})
	// check will change the leave
	check(arr[ans])
	return arr[ans], leave
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
