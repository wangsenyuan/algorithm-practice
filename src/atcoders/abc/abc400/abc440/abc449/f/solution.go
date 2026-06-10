package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int64 {
	var H, W, h, w int64
	var n int
	fmt.Fscan(reader, &H, &W, &h, &w, &n)
	blacks := make([][]int64, n)
	for i := range blacks {
		blacks[i] = make([]int64, 2)
		fmt.Fscan(reader, &blacks[i][0], &blacks[i][1])
	}
	return solve(H, W, h, w, blacks)
}

type event struct {
	x  int64
	y1 int64
	y2 int64
	v  int
}

func solve(H int64, W int64, h int64, w int64, blacks [][]int64) int64 {
	rowLimit := H - h + 1
	colLimit := W - w + 1
	total := rowLimit * colLimit

	var events []event
	var ys []int64

	for _, cur := range blacks {
		r, c := cur[0], cur[1]
		x1, x2 := max(int64(1), r-h+1), min(rowLimit, r)
		y1, y2 := max(int64(1), c-w+1), min(colLimit, c)
		if x1 > x2 || y1 > y2 {
			continue
		}
		events = append(events, event{x1, y1, y2 + 1, 1})
		events = append(events, event{x2 + 1, y1, y2 + 1, -1})
		ys = append(ys, y1, y2+1)
	}

	if len(events) == 0 {
		return total
	}

	sort.Slice(ys, func(i, j int) bool {
		return ys[i] < ys[j]
	})
	ys = unique(ys)

	sort.Slice(events, func(i, j int) bool {
		return events[i].x < events[j].x
	})

	tr := NewSegTree(ys)
	var covered int64
	prev := events[0].x
	for i := 0; i < len(events); {
		x := events[i].x
		covered += tr.Covered() * (x - prev)
		for i < len(events) && events[i].x == x {
			l := sort.Search(len(ys), func(j int) bool {
				return ys[j] >= events[i].y1
			})
			r := sort.Search(len(ys), func(j int) bool {
				return ys[j] >= events[i].y2
			})
			tr.Add(l, r, events[i].v)
			i++
		}
		prev = x
	}

	return total - covered
}

func unique(arr []int64) []int64 {
	res := arr[:0]
	for _, x := range arr {
		if len(res) == 0 || res[len(res)-1] != x {
			res = append(res, x)
		}
	}
	return res
}

type SegTree struct {
	ys  []int64
	cnt []int
	sum []int64
}

func NewSegTree(ys []int64) *SegTree {
	n := len(ys)
	return &SegTree{
		ys:  ys,
		cnt: make([]int, 4*n),
		sum: make([]int64, 4*n),
	}
}

func (tr *SegTree) Covered() int64 {
	return tr.sum[0]
}

func (tr *SegTree) Add(l int, r int, v int) {
	if l >= r {
		return
	}
	tr.add(0, 0, len(tr.ys)-1, l, r, v)
}

func (tr *SegTree) add(i int, l int, r int, ql int, qr int, v int) {
	if qr <= l || r <= ql {
		return
	}
	if ql <= l && r <= qr {
		tr.cnt[i] += v
		tr.pull(i, l, r)
		return
	}
	mid := (l + r) / 2
	tr.add(i*2+1, l, mid, ql, qr, v)
	tr.add(i*2+2, mid, r, ql, qr, v)
	tr.pull(i, l, r)
}

func (tr *SegTree) pull(i int, l int, r int) {
	if tr.cnt[i] > 0 {
		tr.sum[i] = tr.ys[r] - tr.ys[l]
		return
	}
	if l+1 == r {
		tr.sum[i] = 0
		return
	}
	tr.sum[i] = tr.sum[i*2+1] + tr.sum[i*2+2]
}
