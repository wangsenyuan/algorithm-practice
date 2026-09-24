package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	ops := make([][]int, m)
	for i := range m {
		var op string
		fmt.Fscan(reader, &op)
		if op == "Q" {
			ops[i] = make([]int, 4)
			fmt.Fscan(reader, &ops[i][1], &ops[i][2], &ops[i][3])
		} else {
			ops[i] = make([]int, 3)
			ops[i][0] = 1
			fmt.Fscan(reader, &ops[i][1], &ops[i][2])
		}
	}
	return solve(a, ops)
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
	return res
}

func (t fenwick) query(l, r int) int {
	return t.pre(r) - t.pre(l-1)
}

type query struct {
	l int
	r int
	k int
}

func solve(a []int, ops [][]int) []int {
	sorted := slices.Clone(a)
	qs := make([]query, len(a))
	idx := make([]int, len(a))
	// k < 0, 表示update
	for i, v := range a {
		qs[i] = query{i + 1, v, -1}
		idx[i] = i
	}
	for _, cur := range ops {
		if cur[0] == 0 {
			idx = append(idx, len(qs))
			l, r, k := cur[1], cur[2], cur[3]
			qs = append(qs, query{l, r, k})
		} else {
			// update
			l, v := cur[1], cur[2]
			idx = append(idx, len(qs))
			qs = append(qs, query{l, a[l-1], -3})
			idx = append(idx, len(qs))
			qs = append(qs, query{l, v, -1})
			sorted = append(sorted, v)
			a[l-1] = v
		}
	}

	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	tr := make(fenwick, len(a)+2)

	var play func(lo int, hi int, todo []int)

	play = func(lo int, hi int, todo []int) {
		for _, i := range todo {
			if qs[i].k > 0 {
				goto next
			}
		}
		return

	next:
		if lo+1 == hi {
			for _, i := range todo {
				if qs[i].k > 0 {
					qs[i].k = sorted[lo]
				}
			}
			return
		}
		mid := (lo + hi) >> 1
		x := sorted[mid]

		var b, c []int

		for _, i := range todo {
			q := &qs[i]
			if q.k < 0 {
				if q.r < x {
					b = append(b, i)
					tr.update(q.l, q.k+2)
				} else {
					c = append(c, i)
				}
			} else {
				cnt := tr.query(q.l, q.r)
				if cnt >= q.k {
					b = append(b, i)
				} else {
					q.k -= cnt
					c = append(c, i)
				}
			}
		}

		for _, i := range todo {
			if qs[i].k < 0 && qs[i].r < x {
				tr.update(qs[i].l, -(qs[i].k + 2))
			}
		}

		play(lo, mid, b)
		play(mid, hi, c)
	}

	play(0, len(sorted), idx)

	var ans []int
	for _, cur := range qs {
		if cur.k >= 0 {
			ans = append(ans, cur.k)
		}
	}
	return ans
}
