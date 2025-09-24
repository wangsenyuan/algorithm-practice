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
	_, _, _, _, ok, res := drive(reader)
	if !ok {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) (n int, roads [][]int, p int, q int, ok bool, res [][]int) {
	var m int
	fmt.Fscan(reader, &n, &m, &p, &q)
	roads = make([][]int, m)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		roads[i] = []int{u, v, w}
	}
	ok, res = solve(n, roads, p, q)
	return
}

func solve(n int, roads [][]int, p int, q int) (bool, [][]int) {
	set := NewDSU(n)
	for _, cur := range roads {
		u, v := cur[0]-1, cur[1]-1
		set.Union(u, v)
	}

	if set.sz < q {
		// 区域已经比q少了，或者能增加的边太少了
		return false, nil
	}

	var add [][]int
	sum := make([]int, n)
	for _, cur := range roads {
		u, w := cur[0]-1, cur[2]
		u = set.Find(u)
		sum[u] += w
	}

	var pq PriorityQueue

	for u := range n {
		if set.Find(u) == u {
			it := new(Item)
			it.id = u
			it.priority = sum[u]
			heap.Push(&pq, it)
		}
	}

	for pq.Len() > q {
		a := heap.Pop(&pq).(*Item)
		b := heap.Pop(&pq).(*Item)
		w := min(1e9, sum[a.id]+sum[b.id]+1)
		tmp := sum[a.id] + sum[b.id] + w
		add = append(add, []int{a.id + 1, b.id + 1})
		set.Union(a.id, b.id)
		u := set.Find(a.id)
		sum[u] = tmp
		it := new(Item)
		it.id = u
		it.priority = sum[u]
		heap.Push(&pq, it)
	}

	p -= len(add)
	if p < 0 {
		return false, nil
	}
	if p > 0 {
		x := -1
		for u := range n {
			if set.Find(u) == u && set.cnt[u] > 1 {
				x = u
				break
			}
		}
		if x < 0 {
			return false, nil
		}

		var y int
		for u := range n {
			if u != x && set.Find(u) == x {
				y = u
				break
			}
		}

		for p > 0 {
			p--
			add = append(add, []int{x + 1, y + 1})
		}
	}

	return true, add
}

type DSU struct {
	arr []int
	cnt []int
	sz  int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt, n}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	this.sz--
	return true
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
