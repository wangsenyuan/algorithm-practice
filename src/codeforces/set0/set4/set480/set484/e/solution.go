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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	h := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &h[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	queries := make([][]int, m)
	for i := range m {
		var l, r, w int
		fmt.Fscan(reader, &l, &r, &w)
		queries[i] = []int{l, r, w}
	}
	return solve(h, queries)
}

type fence struct {
	pos int
	h   int
}

func solve(h []int, queries [][]int) []int {
	n := len(h)
	arr := make([]fence, n)
	for i, v := range h {
		arr[i] = fence{pos: i, h: v}
	}
	slices.SortFunc(arr, func(a, b fence) int {
		return cmp.Or(b.h-a.h, a.pos-b.pos)
	})

	// Persistent segment tree with node indices instead of pointers
	type segNode struct {
		left  int
		right int
		cnt   int
		pref  int
		suf   int
	}

	type segVal struct {
		cnt  int
		pref int
		suf  int
		len  int
	}

	seg := make([]segNode, 0, n*25)

	var merge func(a, b segVal) segVal
	merge = func(a, b segVal) segVal {
		if a.len == 0 {
			return b
		}
		if b.len == 0 {
			return a
		}
		res := segVal{}
		res.len = a.len + b.len
		res.cnt = max(a.cnt, b.cnt, a.suf+b.pref)
		res.pref = a.pref
		if a.cnt == a.len {
			res.pref += b.pref
		}
		res.suf = b.suf
		if b.cnt == b.len {
			res.suf += a.suf
		}
		return res
	}

	var build func(l, r int) int
	build = func(l, r int) int {
		id := len(seg)
		seg = append(seg, segNode{})
		if l == r {
			return id
		}
		mid := (l + r) >> 1
		seg[id].left = build(l, mid)
		seg[id].right = build(mid+1, r)
		return id
	}

	var pull func(id, l, r int)
	pull = func(id, l, r int) {
		if l == r {
			return
		}
		mid := (l + r) >> 1
		ln := seg[seg[id].left]
		rn := seg[seg[id].right]
		leftVal := segVal{
			cnt:  ln.cnt,
			pref: ln.pref,
			suf:  ln.suf,
			len:  mid - l + 1,
		}
		rightVal := segVal{
			cnt:  rn.cnt,
			pref: rn.pref,
			suf:  rn.suf,
			len:  r - mid,
		}
		res := merge(leftVal, rightVal)
		seg[id].cnt = res.cnt
		seg[id].pref = res.pref
		seg[id].suf = res.suf
	}

	var update func(prev, l, r, pos int) int
	update = func(prev, l, r, pos int) int {
		id := len(seg)
		seg = append(seg, seg[prev]) // copy previous node
		if l == r {
			seg[id].cnt = 1
			seg[id].pref = 1
			seg[id].suf = 1
			return id
		}
		mid := (l + r) >> 1
		if pos <= mid {
			seg[id].left = update(seg[prev].left, l, mid, pos)
		} else {
			seg[id].right = update(seg[prev].right, mid+1, r, pos)
		}
		pull(id, l, r)
		return id
	}

	var query func(id, l, r, ql, qr int) segVal
	query = func(id, l, r, ql, qr int) segVal {
		if qr < l || r < ql {
			return segVal{}
		}
		if ql <= l && r <= qr {
			nn := seg[id]
			return segVal{
				cnt:  nn.cnt,
				pref: nn.pref,
				suf:  nn.suf,
				len:  r - l + 1,
			}
		}
		mid := (l + r) >> 1
		if qr <= mid {
			return query(seg[id].left, l, mid, ql, qr)
		}
		if ql > mid {
			return query(seg[id].right, mid+1, r, ql, qr)
		}
		leftVal := query(seg[id].left, l, mid, ql, qr)
		rightVal := query(seg[id].right, mid+1, r, ql, qr)
		return merge(leftVal, rightVal)
	}

	baseRoot := build(0, n-1)

	roots := make([]int, 0, n+1)
	hs := make([]int, 0, n+1)
	roots = append(roots, baseRoot)
	hs = append(hs, 0)

	for i := 0; i < n; {
		j := i
		root := roots[len(roots)-1]
		for i < n && arr[i].h == arr[j].h {
			root = update(root, 0, n-1, arr[i].pos)
			i++
		}
		roots = append(roots, root)
		hs = append(hs, arr[j].h)
	}

	ans := make([]int, len(queries))

	find := func(l int, r int, w int) int {
		i := sort.Search(len(roots), func(i int) bool {
			if seg[roots[i]].cnt < w {
				return false
			}
			return query(roots[i], 0, n-1, l, r).cnt >= w
		})
		return hs[i]
	}

	for i, cur := range queries {
		l, r, w := cur[0]-1, cur[1]-1, cur[2]
		ans[i] = find(l, r, w)
	}

	return ans
}
