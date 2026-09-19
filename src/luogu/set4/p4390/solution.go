package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var op, w int
	fmt.Fscan(reader, &op, &w)
	var ops [][]int
	for {
		fmt.Fscan(reader, &op)
		if op == 3 {
			break
		}
		if op == 1 {
			var x, y, a int
			fmt.Fscan(reader, &x, &y, &a)
			ops = append(ops, []int{1, x, y, a})
			continue
		}
		var x1, y1, x2, y2 int
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		ops = append(ops, []int{2, x1, y1, x2, y2})
	}
	return solve(w, ops)
}

func solve(mx int, ops [][]int) []int {

	type data struct{ x, y, c, res int }
	a := make([]data, 0, 2e5)
	var q int
	for _, cur := range ops {
		op, x1, y1, x2 := cur[0], cur[1], cur[2], cur[3]
		x1++
		if op == 3 {
			break
		}
		if op == 1 {
			a = append(a, data{x1, y1, x2, -1})
		} else {
			x2++
			y2 := cur[4]
			q++
			a = append(a,
				data{x1 - 1, y1 - 1, q, 0},
				data{x1 - 1, y2, -q, 0},
				data{x2, y1 - 1, -q, 0},
				data{x2, y2, q, 0})
		}
	}

	t := make(fenwick, mx+2)
	var play func(int, int)
	play = func(l, r int) {
		if l+1 == r {
			return
		}
		mid := (l + r) >> 1
		play(l, mid)
		play(mid, r)

		i := l
		for j := mid; j < r; j++ {
			for ; i < mid && a[i].y <= a[j].y; i++ {
				if a[i].res < 0 {
					t.update(a[i].x, a[i].c)
				}
			}
			if a[j].res >= 0 {
				a[j].res += t.pre(a[j].x)
			}
		}
		for i--; i >= l; i-- {
			if a[i].res < 0 {
				t.update(a[i].x, -a[i].c)
			}
		}

		slices.SortFunc(a[l:r], func(a, b data) int { return a.y - b.y })
	}

	play(0, len(a))

	ans := make([]int, q)
	for _, d := range a {
		if d.res > 0 {
			if d.c > 0 {
				ans[d.c-1] += d.res
			} else {
				ans[-d.c-1] -= d.res
			}
		}
	}

	return ans
}

type fenwick []int

func (t fenwick) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}
