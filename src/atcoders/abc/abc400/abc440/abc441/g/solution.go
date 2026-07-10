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

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int64 {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	queries := make([][]int, q)
	for i := range q {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			queries[i] = make([]int, 4)
			queries[i][0] = t
			fmt.Fscan(reader, &queries[i][1], &queries[i][2], &queries[i][3])
		} else {
			queries[i] = make([]int, 3)
			queries[i][0] = t
			fmt.Fscan(reader, &queries[i][1], &queries[i][2])
		}
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []int64 {
	tr := make(seg, n*4)
	var ans []int64
	for _, cur := range queries {
		switch cur[0] {
		case 1:
			l, r, x := cur[1], cur[2], cur[3]
			tr.update(1, 1, n, l, r, tag{0, x})
		case 2:
			l, r := cur[1], cur[2]
			tr.update(1, 1, n, l, r, tag{1, 0})
		default:
			l, r := cur[1], cur[2]
			ans = append(ans, int64(tr.query(1, 1, n, l, r)))
		}
	}

	return ans
}

// https://github.com/EndlessCheng
type data struct{ state, mx int }
type tag struct{ flipCnt, add int }
type seg []struct {
	data
	tag
}

func mergeData(l, r data) data {
	state := 1
	if l.state == r.state {
		state = l.state
	}
	return data{state, max(l.mx, r.mx)}
}

func mergeTag(f, old tag) tag {
	if f.flipCnt == 0 {
		old.add += f.add
		return old
	}
	f.flipCnt += old.flipCnt
	return f
}

func (t seg) apply(o int, f tag) {
	cur := &t[o]
	if f.flipCnt > 0 {
		if f.flipCnt%2 > 0 {
			cur.state = 2 - cur.state
		}
		cur.mx = 0
	}
	if cur.state != 2 {
		cur.mx += f.add
	}

	cur.tag = mergeTag(f, cur.tag)
}

func (t seg) spread(o int) {
	f := t[o].tag
	if f == (tag{}) {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].tag = tag{}
}

func (t seg) update(o, l, r, ql, qr int, f tag) {
	if ql <= l && r <= qr {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (l + r) >> 1
	if ql <= m {
		t.update(o<<1, l, m, ql, qr, f)
	}
	if m < qr {
		t.update(o<<1|1, m+1, r, ql, qr, f)
	}
	t[o].data = mergeData(t[o<<1].data, t[o<<1|1].data)
}

func (t seg) query(o, l, r, ql, qr int) int {
	if ql <= l && r <= qr {
		return t[o].mx
	}
	t.spread(o)
	m := (l + r) >> 1
	if qr <= m {
		return t.query(o<<1, l, m, ql, qr)
	}
	if ql > m {
		return t.query(o<<1|1, m+1, r, ql, qr)
	}
	return max(t.query(o<<1, l, m, ql, qr), t.query(o<<1|1, m+1, r, ql, qr))
}
