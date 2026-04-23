package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, _, _, d, res := drive(reader)
		fmt.Fprintln(writer, d)
		if d >= 0 {
			s := fmt.Sprintf("%v", res)
			fmt.Fprintln(writer, s[1:len(s)-1])
		}
	}
}

func drive(reader *bufio.Reader) (k int, a []int, b []int, d int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &k)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	d, res = solve(k, a, b)
	return
}

func solve(k int, a []int, b []int) (int, []int) {
	n := len(a)
	pref := make([][2]int, n+1)
	score := make([][2]int, n+1)
	for i := range n {
		pref[i+1][0] = pref[i][0] + a[i]
		pref[i+1][1] = pref[i][1] + b[i]
		if pref[i+1][0] >= pref[i+1][1] {
			score[i+1][0] = pref[i+1][0] - pref[i+1][1]
			score[i+1][1] = 0
		} else {
			score[i+1][0] = 0
			score[i+1][1] = pref[i+1][1] - pref[i+1][0]
		}
	}

	win := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		j1 := sort.Search(n+1, func(j int) bool {
			return pref[j][0]-pref[i][0]+score[i][0] >= k
		})
		j2 := sort.Search(n+1, func(j int) bool {
			return pref[j][1]-pref[i][1]+score[i][1] >= k
		})
		win[i] = j2 < j1
	}

	if win[0] {
		return 0, nil
	}

	var pq PriorityQueue

	// priority = -(-pref_a[i] + score_a)
	heap.Push(&pq, &data{id: 0, priority: 0})

	dp := make(PriorityQueue, n+1)
	items := make([]*data, n+1)
	for i := range n + 1 {
		it := &data{id: i, priority: 1 << 60}
		it.index = i
		items[i] = it
		dp[i] = it
	}
	items[0].priority = 0
	heap.Init(&dp)

	fp := make([]int, n+1)
	best := n
	best_id := -1
	for r := 1; r <= n; r++ {
		// 即使在这些位置reset，也无法到达位置r
		for pq.Len() > 0 && -pq[0].priority+pref[r][0] >= k {
			it := heap.Pop(&pq).(*data)
			// dp失效
			heap.Remove(&dp, items[it.id].index)
		}
		if len(pq) == 0 {
			// 无法通过r的时候，也无法到达后面的位置
			break
		}
		w := dp[0].priority
		if win[r] && w+1 < best {
			best = w + 1
			best_id = r
		}
		fp[r] = dp[0].id
		heap.Push(&pq, &data{id: r, priority: -(-pref[r][0] + score[r][0])})
		items[r].priority = w + 1
		heap.Fix(&dp, items[r].index)
	}
	if best_id < 0 {
		return -1, nil
	}
	var res []int
	for r := best_id; r > 0; r = fp[r] {
		res = append(res, r)
	}

	slices.Reverse(res)
	return best, res
}

type data struct {
	id       int
	priority int
	index    int
}

type PriorityQueue []*data

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*data)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
