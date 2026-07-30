package main

import (
	"bufio"
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
	var n, d, w int
	fmt.Fscan(reader, &n, &d, &w)
	apples := make([][2]int, n)
	for i := range n {
		fmt.Fscan(reader, &apples[i][0], &apples[i][1])
	}
	return solve(n, d, w, apples)
}

func solve(n, d, w int, apples [][2]int) int {
	var times []int
	var pos []int
	for _, cur := range apples {
		t, x := cur[0], cur[1]
		times = append(times, t, t+d)
		pos = append(pos, x, x-w+1)
	}

	slices.Sort(times)
	times = slices.Compact(times)
	slices.Sort(pos)
	pos = slices.Compact(pos)

	active := make([][]int, len(times))
	expire := make([][]int, len(times))
	for i, cur := range apples {
		t := cur[0]
		j := sort.SearchInts(times, t)
		active[j] = append(active[j], i)

		j = sort.SearchInts(times, t+d)
		expire[j] = append(expire[j], i)
	}

	tr := NewTree(len(pos))

	var best int

	for i := range times {
		for _, j := range active[i] {
			x := apples[j][1]
			// 它在区间 (x - w, x] 中有效
			// 假设w等于1, x = 2, 当L = 1的时候, L + w - 0.5 = 1.5, 不包含x
			// 如果要包含x, 必须是在区间 x - w + 1
			l := sort.SearchInts(pos, x-w+1)
			r := sort.SearchInts(pos, x)
			tr.update(l, r, 1)
		}

		for _, j := range expire[i] {
			x := apples[j][1]
			l := sort.SearchInts(pos, x-w+1)
			r := sort.SearchInts(pos, x)
			tr.update(l, r, -1)
		}
		best = max(best, tr.val[0])
	}

	return best
}

type tree struct {
	val  []int
	lazy []int
}

func NewTree(n int) *tree {
	return &tree{
		val:  make([]int, 4*n),
		lazy: make([]int, 4*n),
	}
}

func (t *tree) apply(i int, v int) {
	t.val[i] += v
	t.lazy[i] += v
}

func (t *tree) push(i int) {
	if t.lazy[i] != 0 {
		t.apply(i*2+1, t.lazy[i])
		t.apply(i*2+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *tree) update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.apply(i, v)
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		t.val[i] = max(t.val[i*2+1], t.val[i*2+2])
	}
	n := len(t.val) / 4
	f(0, 0, n-1, L, R)
}
