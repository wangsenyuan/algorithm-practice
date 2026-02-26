package main

import (
	"bufio"
	"cmp"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	fmt.Println(res[0])
	if res[0] > 0 {
		fmt.Println(res[1], res[2])
	}
}

func drive(reader *bufio.Reader) (shows [][]int, channels [][]int, res []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	shows = make([][]int, n)
	for i := range n {
		shows[i] = make([]int, 2)
		fmt.Fscan(reader, &shows[i][0], &shows[i][1])
	}
	channels = make([][]int, m)
	for i := range m {
		channels[i] = make([]int, 3)
		fmt.Fscan(reader, &channels[i][0], &channels[i][1], &channels[i][2])
	}
	res = solve(shows, channels)
	return
}

type data struct {
	pos int
	id  int
}

type Event struct {
	pos       int
	id        int
	tp        int // 0 for show, 1 for channel
	eventType int // 0 for start, 1 for end
}

type result struct {
	val    int
	showId int
}

func solve(shows [][]int, channels [][]int) []int {
	n := len(shows)
	m := len(channels)

	arr := make([]data, n)
	for i, cur := range shows {
		arr[i] = data{pos: cur[0], id: i}
	}

	slices.SortFunc(arr, func(a, b data) int {
		return cmp.Or(a.pos-b.pos, a.id-b.id)
	})

	pos := make([]int, n)

	for i := range n {
		pos[arr[i].id] = i
	}

	var events []Event

	its1 := make([]*Item, n)
	its2 := make([]*Item, n)
	for i := range n {
		its1[i] = new(Item)
		its1[i].id = i
		its1[i].priority = shows[i][0]
		its2[i] = new(Item)
		its2[i].id = i
		its2[i].priority = -shows[i][1]
		events = append(events, Event{pos: shows[i][0], id: i, tp: 0, eventType: 0})
		events = append(events, Event{pos: shows[i][1], id: i, tp: 0, eventType: 3})
	}

	for i := range m {
		events = append(events, Event{pos: channels[i][0], id: i, tp: 1, eventType: 1})
		events = append(events, Event{pos: channels[i][1], id: i, tp: 1, eventType: 2})
	}

	slices.SortFunc(events, func(a, b Event) int {
		return cmp.Or(a.pos-b.pos, a.eventType-b.eventType, a.tp-b.tp)
	})

	var pq1 MinHeap
	var pq2 MinHeap
	tr := NewSegTree(n)

	ans := make([]result, m)

	update := func(i int, val int, sid int) {
		if val > ans[i].val {
			ans[i] = result{val: val, showId: sid}
		}
	}

	for _, evt := range events {
		i := evt.id
		if evt.tp == 0 {
			if evt.eventType == 0 {
				// a show start
				heap.Push(&pq1, its1[i])
				heap.Push(&pq2, its2[i])
			} else {
				// a show end
				tr.Update(pos[i], shows[i][1]-shows[i][0])
				heap.Remove(&pq1, its1[i].index)
				heap.Remove(&pq2, its2[i].index)
			}
		} else {
			if evt.eventType == 1 {
				if pq2.Len() > 0 {
					y := min(channels[i][1], -pq2[0].priority)
					update(i, (y-channels[i][0])*channels[i][2], pq2[0].id)
				}
			} else {
				if pq1.Len() > 0 {
					x := max(channels[i][0], pq1[0].priority)
					update(i, (channels[i][1]-x)*channels[i][2], pq1[0].id)
				}
				j := sort.Search(n, func(j int) bool {
					return arr[j].pos >= channels[i][0]
				})
				if j < n {
					// w.second 是位置
					w := tr.Get(j, n)
					if w.first > 0 {
						update(i, w.first*channels[i][2], arr[w.second].id)
					}
				}
			}
		}
	}

	best := []int{0, 1, 1}
	for i := range m {
		if ans[i].val > best[0] {
			best[0] = ans[i].val
			best[1] = ans[i].showId + 1
			best[2] = i + 1
		}
	}
	return best
}

type Item struct {
	id       int
	priority int
	index    int
}

type MinHeap []*Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].priority < h[j].priority
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *MinHeap) Push(x any) {
	item := x.(*Item)
	item.index = len(*h)
	*h = append(*h, item)
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	item.index = -1
	return item
}

type pair struct {
	first  int
	second int
}

func maxPair(a, b pair) pair {
	if a.first > b.first {
		return a
	}
	return b
}

type SegTree []pair

func NewSegTree(n int) SegTree {
	arr := make([]pair, 2*n)
	for i := n; i < len(arr); i++ {
		arr[i] = pair{0, i - n}
	}
	for i := n - 1; i > 0; i-- {
		arr[i] = maxPair(arr[i*2], arr[i*2+1])
	}
	return SegTree(arr)
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p].first = v
	for p > 1 {
		tr[p>>1] = maxPair(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) pair {
	n := len(tr) / 2
	l += n
	r += n
	res := pair{0, -1}
	for l < r {
		if l&1 == 1 {
			res = maxPair(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = maxPair(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
