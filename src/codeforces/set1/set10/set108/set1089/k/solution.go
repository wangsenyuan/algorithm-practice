package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

type event struct {
	op byte
	a  int
	b  int
}

func drive(reader *bufio.Reader) []int64 {
	var q int
	fmt.Fscan(reader, &q)
	events := make([]event, q)
	for i := range q {
		var op string
		fmt.Fscan(reader, &op)
		events[i].op = op[0]
		switch op[0] {
		case '+':
			fmt.Fscan(reader, &events[i].a, &events[i].b)
		case '-', '?':
			fmt.Fscan(reader, &events[i].a)
		}
	}
	return solve(events)
}

func solve(events []event) []int64 {
	var mt int
	for _, e := range events {
		if e.op == '+' || e.op == '?' {
			mt = max(mt, e.a)
		}
	}
	tr := NewTree(mt + 1)

	var ans []int64

	for _, cur := range events {
		switch cur.op {
		case '+':
			tr.update(cur.a, cur.b)
		case '-':
			i := cur.a - 1
			evt := events[i]
			t, d := evt.a, evt.b
			tr.update(t, -d)
		default:
			// ?
			t := cur.a
			st := tr.query(0, t)
			tmp := st.fixed - t
			ans = append(ans, max(int64(tmp), 0))
		}
	}

	return ans
}

type state struct {
	sum   int
	fixed int
}

const inf = 1 << 60

func merge(a state, b state) state {
	return state{
		sum:   a.sum + b.sum,
		fixed: max(a.fixed+b.sum, b.fixed),
	}
}

type tree []state

func NewTree(n int) tree {
	arr := make([]state, n*4)
	for i := range arr {
		arr[i] = state{
			sum:   0,
			fixed: -inf,
		}
	}
	return arr
}

func (tr tree) update(t int, d int) {
	n := len(tr) / 4
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			if d < 0 {
				tr[i].sum = 0
				tr[i].fixed = -inf
			} else {
				tr[i].sum = d
				tr[i].fixed = t + d
			}
		} else {
			mid := (l + r) >> 1
			if t <= mid {
				f(i*2+1, l, mid)
			} else {
				f(i*2+2, mid+1, r)
			}
			tr[i] = merge(tr[i*2+1], tr[i*2+2])
		}
	}
	f(0, 0, n-1)
}

func (tr tree) query(L int, R int) state {
	var f func(i int, l int, r int, L int, R int) state
	f = func(i int, l int, r int, L int, R int) state {
		if l == L && r == R {
			return tr[i]
		}
		mid := (l + r) >> 1
		res := state{
			sum:   0,
			fixed: -inf,
		}
		if L <= mid {
			res = f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			res = merge(res, f(i*2+2, mid+1, r, max(mid+1, L), R))
		}
		return res
	}
	n := len(tr) / 4
	return f(0, 0, n-1, L, R)
}
