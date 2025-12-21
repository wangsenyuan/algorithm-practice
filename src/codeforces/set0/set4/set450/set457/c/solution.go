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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	voters := make([][]int, n)
	for i := range n {
		voters[i] = make([]int, 2)
		fmt.Fscan(reader, &voters[i][0], &voters[i][1])
	}
	return solve(voters)
}

func prepare(people [][]int) [][]int {
	slices.SortFunc(people, func(a, b []int) int {
		return b[1] - a[1]
	})

	var ids []int
	ids = append(ids, 0)

	for _, cur := range people {
		ids = append(ids, cur[0])
	}

	slices.Sort(ids)
	ids = slices.Compact(ids)
	m := len(ids)
	arr := make([][]int, m)

	for _, cur := range people {
		j := sort.SearchInts(ids, cur[0])
		arr[j] = append(arr[j], cur[1])
	}
	return arr
}

func solve(people [][]int) int {
	arr := prepare(people)

	cnt0 := len(arr[0])
	n := len(people)
	k := n - cnt0
	todo := make([][]int, k)
	m := len(arr)

	var cost int
	var cnt int

	var nums []int

	for i := 1; i < m; i++ {
		for j := 0; j < len(arr[i]); j++ {
			todo[j] = append(todo[j], arr[i][j])
			cost += arr[i][j]
			cnt++
			nums = append(nums, arr[i][j])
		}
	}
	slices.Sort(nums)
	nums = slices.Compact(nums)
	w := len(nums)
	tr := NewTree(w)

	// buy every one
	best := cost

	for d := range k {
		// 这些人做为候选
		for _, v := range todo[d] {
			i := sort.SearchInts(nums, v)
			tr.Update(i, v)
			cost -= v
			cnt--
		}

		if cnt0+cnt > d+1 {
			// 其他人最多d+1张票
			best = min(best, cost)
		} else {
			// cnt0 + cnt <= d + 1
			need := d + 2 - (cnt0 + cnt)
			// 还需要这么多张票
			if tr.cnt[0] >= need {
				tmp := tr.GetFirstKSum(need)
				best = min(best, cost+tmp)
			}
		}
	}

	return best
}

func solve1(people [][]int) int {

	arr := prepare(people)
	m := len(arr)

	check := func(k int) int {
		cnt := len(arr[0])
		var pq IntHeap
		var sum int
		for i := 1; i < m; i++ {
			for j := 0; j < min(k, len(arr[i])); j++ {
				heap.Push(&pq, arr[i][j])
			}
			for j := k; j < len(arr[i]); j++ {
				sum += arr[i][j]
				cnt++
			}
		}
		for cnt <= k && pq.Len() > 0 {
			sum += heap.Pop(&pq).(int)
			cnt++
		}
		return sum
	}

	l, r := 0, len(people)-len(arr[0])

	for l < r {
		mid := (l + r) >> 1
		x := check(mid)
		y := check(mid + 1)
		if x < y {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return check(l)
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

type Tree struct {
	val []int
	cnt []int
	sz  int
}

func NewTree(n int) *Tree {
	return &Tree{
		val: make([]int, 4*n),
		cnt: make([]int, 4*n),
		sz:  n,
	}
}

func (tr *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.val[i] += v
			tr.cnt[i]++
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.val[i] = tr.val[i*2+1] + tr.val[i*2+2]
		tr.cnt[i] = tr.cnt[i*2+1] + tr.cnt[i*2+2]
	}
	f(0, 0, tr.sz-1)
}

func (tr *Tree) GetFirstKSum(k int) int {
	var f func(i int, l int, r int, k int) int
	f = func(i int, l int, r int, k int) int {
		if l == r {
			return tr.val[i] / tr.cnt[i] * k
		}
		mid := (l + r) / 2
		if tr.cnt[i*2+1] >= k {
			return f(i*2+1, l, mid, k)
		}
		return tr.val[i*2+1] + f(i*2+2, mid+1, r, k-tr.cnt[i*2+1])
	}

	return f(0, 0, tr.sz-1, k)
}
