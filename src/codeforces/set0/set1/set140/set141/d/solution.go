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
	_, _, best, res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, best)
	fmt.Fprintln(writer, len(res))
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) (l int, ramps [][]int, best int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &l)
	ramps = make([][]int, n)
	for i := range n {
		ramps[i] = make([]int, 4)
		fmt.Fscan(reader, &ramps[i][0], &ramps[i][1], &ramps[i][2], &ramps[i][3])
	}
	best, res = solve(l, ramps)
	return
}

type Ramp struct {
	id int
	x  int
	d  int
	t  int
	p  int
}

type Pair struct {
	first  int
	second int
}

func minPair(a, b Pair) Pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

func solve(l int, ramps [][]int) (res int, use []int) {

	n := len(ramps)
	arr := make([]Ramp, n)

	var ys []int
	ys = append(ys, 0)

	for i, cur := range ramps {
		x, d, t, p := cur[0], cur[1], cur[2], cur[3]
		arr[i] = Ramp{i, x, d, t, p}
		ys = append(ys, x+d)
		if x-p >= 0 {
			ys = append(ys, x-p)
		}
	}
	slices.Sort(ys)
	ys = slices.Compact(ys)
	m := len(ys)

	slices.SortFunc(arr, func(a Ramp, b Ramp) int {
		return a.x - b.x
	})

	t1 := NewSegTree(m)
	t2 := NewSegTree(m)

	res = l
	last := -1

	fa := make([]int, n)

	for i := range n {
		fa[i] = -1
	}

	for _, cur := range arr {
		x, d, t, p := cur.x, cur.d, cur.t, cur.p
		if x-p < 0 {
			// 这个跳台无法被使用到
			continue
		}
		// 不借助前面ramp时，最好的成绩就是x
		best := Pair{x, -1}
		i := sort.SearchInts(ys, x-p)
		// 那么在这之前的最优解
		tmp1 := t1.Query(0, i+1)

		if tmp1.first+x < best.first {
			best = Pair{tmp1.first + x, tmp1.second}
		}

		tmp2 := t2.Query(i, m)

		if tmp2.first+2*p-x < best.first {
			best = Pair{tmp2.first + 2*p - x, tmp2.second}
		}

		best.first += t

		fa[cur.id] = best.second

		// else 无法起跳
		// 落掉在y
		y := x + d
		if res > best.first+l-y {
			res = best.first + l - y
			last = cur.id
		}

		i = sort.SearchInts(ys, y)
		if t1.Query(i, i+1).first > best.first-y {
			t1.Update(i, Pair{best.first - y, cur.id})
		}

		if t2.Query(i, i+1).first > best.first+y {
			t2.Update(i, Pair{best.first + y, cur.id})
		}
	}

	for last != -1 {
		use = append(use, last+1)
		last = fa[last]
	}

	slices.Reverse(use)

	return
}

const inf = 1 << 60

type SegTree []Pair

func NewSegTree(n int) SegTree {
	arr := make([]Pair, 2*n)
	for i := range arr {
		arr[i] = Pair{inf, -1}
	}
	return SegTree(arr)
}

func (t SegTree) Update(p int, v Pair) {
	p += len(t) / 2
	t[p] = v
	for p > 1 {
		t[p>>1] = minPair(t[p], t[p^1])
		p >>= 1
	}
}

func (t SegTree) Query(l int, r int) Pair {
	l += len(t) / 2
	r += len(t) / 2
	res := Pair{inf, -1}
	for l < r {
		if l&1 == 1 {
			res = minPair(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = minPair(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
