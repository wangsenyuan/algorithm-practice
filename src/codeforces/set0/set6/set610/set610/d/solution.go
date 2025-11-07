package main

import (
	"bufio"
	"cmp"
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
	segments := make([][]int, n)
	for i := range n {
		segments[i] = make([]int, 4)
		fmt.Fscan(reader, &segments[i][0], &segments[i][1], &segments[i][2], &segments[i][3])
	}
	return solve(n, segments)
}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func solve(n int, segments [][]int) int {
	var vertical []Line
	var horizontal []Line
	var xs []int
	var ys []int
	for _, cur := range segments {
		x1, y1, x2, y2 := cur[0], cur[1], cur[2], cur[3]
		x1, x2 = min(x1, x2), max(x1, x2)
		y1, y2 = min(y1, y2), max(y1, y2)
		if x1 == x2 {
			vertical = append(vertical, Line{x1, y1, x2, y2})
		} else {
			horizontal = append(horizontal, Line{x1, y1, x2, y2})
		}
		xs = append(xs, x1, x2)
		ys = append(ys, y1, y2)
	}

	slices.SortFunc(vertical, func(a, b Line) int {
		return cmp.Or(a.x1-b.x1, a.y1-b.y1)
	})

	var p int
	for i := 0; i < len(vertical); {
		x := vertical[i].x1
		lo := vertical[i].y1
		hi := vertical[i].y2
		for i < len(vertical) && vertical[i].x1 == x {
			if vertical[i].y1 <= hi {
				hi = max(hi, vertical[i].y2)
			} else {
				vertical[p] = Line{x, lo, x, hi}
				p++
				lo = vertical[i].y1
				hi = vertical[i].y2
			}
			i++
		}
		vertical[p] = Line{x, lo, x, hi}
		p++
	}
	// 现在vertical没有重叠的线段了
	vertical = vertical[:p]

	slices.SortFunc(horizontal, func(a, b Line) int {
		return cmp.Or(a.y1-b.y1, a.x1-b.x1)
	})

	p = 0
	for i := 0; i < len(horizontal); {
		y := horizontal[i].y1
		lf, rg := horizontal[i].x1, horizontal[i].x2
		for i < len(horizontal) && horizontal[i].y1 == y {
			if horizontal[i].x1 <= rg {
				rg = max(rg, horizontal[i].x2)
			} else {
				horizontal[p] = Line{lf, y, rg, y}
				p++
				lf = horizontal[i].x1
				rg = horizontal[i].x2
			}
			i++
		}
		horizontal[p] = Line{lf, y, rg, y}
		p++
	}
	horizontal = horizontal[:p]

	slices.Sort(xs)
	xs = slices.Compact(xs)
	slices.Sort(ys)
	ys = slices.Compact(ys)

	// sum = 所有的区域的和
	// - 有交叉的地方
	var res int
	for _, cur := range vertical {
		res += cur.y2 - cur.y1 + 1
	}
	for _, cur := range horizontal {
		res += cur.x2 - cur.x1 + 1
	}

	begin := make([][]int, len(ys))
	end := make([][]int, len(ys))
	at := make([][]int, len(ys))

	for _, cur := range vertical {
		u := sort.SearchInts(xs, cur.x1)
		j := sort.SearchInts(ys, cur.y1)
		begin[j] = append(begin[j], u)
		j = sort.SearchInts(ys, cur.y2)
		end[j] = append(end[j], u)
	}

	for i, cur := range horizontal {
		u := sort.SearchInts(ys, cur.y1)
		at[u] = append(at[u], i)
	}

	sweep := make(BIT, len(xs)+3)

	for i := range len(ys) {
		for _, x := range begin[i] {
			sweep.update(x, 1)
		}

		for _, j := range at[i] {
			l, r := horizontal[j].x1, horizontal[j].x2
			l = sort.SearchInts(xs, l)
			r = sort.SearchInts(xs, r)
			res -= sweep.query(l, r)
		}

		for _, x := range end[i] {
			sweep.update(x, -1)
		}
	}

	return res
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) get(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) query(l int, r int) int {
	return bit.get(r) - bit.get(l-1)
}
