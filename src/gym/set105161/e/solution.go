package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
	"os"
)

const maxValue = 100000

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	data, _ := io.ReadAll(reader)
	pos := 0
	readInt := func() int {
		for pos < len(data) && (data[pos] < '0' || data[pos] > '9') {
			pos++
		}
		res := 0
		for pos < len(data) && data[pos] >= '0' && data[pos] <= '9' {
			res = res*10 + int(data[pos]-'0')
			pos++
		}
		return res
	}

	n, q := readInt(), readInt()
	a := make([]int, n)
	for i := range a {
		a[i] = readInt()
	}
	queries := make([][3]int, q)
	for i := range queries {
		queries[i] = [3]int{readInt(), readInt(), readInt()}
	}
	return solve(a, queries)
}

type persistentTree struct {
	left, right, count []int
	roots              []int
}

func newPersistentTree(a []int) *persistentTree {
	t := &persistentTree{
		left:  make([]int, 1, len(a)*18+1),
		right: make([]int, 1, len(a)*18+1),
		count: make([]int, 1, len(a)*18+1),
		roots: make([]int, len(a)+1),
	}
	for i, v := range a {
		t.roots[i+1] = t.update(t.roots[i], 0, maxValue, v)
	}
	return t
}

func (t *persistentTree) update(prev, lo, hi, value int) int {
	cur := len(t.count)
	t.left = append(t.left, t.left[prev])
	t.right = append(t.right, t.right[prev])
	t.count = append(t.count, t.count[prev]+1)
	if lo == hi {
		return cur
	}
	mid := (lo + hi) >> 1
	if value <= mid {
		t.left[cur] = t.update(t.left[prev], lo, mid, value)
	} else {
		t.right[cur] = t.update(t.right[prev], mid+1, hi, value)
	}
	return cur
}

func (t *persistentTree) countLess(rightRoot, leftRoot, value int) int {
	if value <= 0 {
		return 0
	}
	if value > maxValue {
		return t.count[rightRoot] - t.count[leftRoot]
	}
	lo, hi := 0, maxValue
	res := 0
	for lo < hi {
		mid := (lo + hi) >> 1
		if value <= mid {
			rightRoot = t.left[rightRoot]
			leftRoot = t.left[leftRoot]
			hi = mid
		} else {
			res += t.count[t.left[rightRoot]] - t.count[t.left[leftRoot]]
			rightRoot = t.right[rightRoot]
			leftRoot = t.right[leftRoot]
			lo = mid + 1
		}
	}
	if lo < value {
		res += t.count[rightRoot] - t.count[leftRoot]
	}
	return res
}

func (t *persistentTree) operationsNeeded(left, right, target, limit int) int {
	length := right - left
	used := 0
	for threshold := target + 1; threshold <= maxValue; threshold <<= 1 {
		used += length - t.countLess(t.roots[right], t.roots[left], threshold)
		if used > limit {
			return used
		}
	}
	return used
}

func solve1(a []int, queries [][3]int) []int {
	tree := newPersistentTree(a)
	ans := make([]int, len(queries))
	for i, query := range queries {
		left, right, k := query[0]-1, query[1], query[2]
		lo, hi := 0, maxValue
		for lo < hi {
			mid := (lo + hi) >> 1
			if tree.operationsNeeded(left, right, mid, k) <= k {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		ans[i] = lo
	}
	return ans
}

// https://github.com/EndlessCheng
type fenwick []int

func (t fenwick) reset(i int) {
	for ; i < len(t); i += i & -i {
		t[i] = 0
	}
}

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

func solve(a []int, queries [][3]int) []int {
	n := len(a)
	idx := make([]int, 0, n)
	for i := range a {
		if a[i] > 0 {
			idx = append(idx, i)
		}
	}

	type query struct{ l, r, k int }
	m := len(queries)
	qs := make([]query, m)
	qIdx := make([]int, m)
	for i := range qs {
		qs[i] = query{l: queries[i][0], r: queries[i][1], k: queries[i][2]}
		qIdx[i] = i
	}

	t := make(fenwick, n+1)

	var play func([]int, []int, int, int)
	play = func(idx, qIdx []int, low, high int) {
		if len(qIdx) == 0 {
			return
		}

		if low+1 == high {
			for _, i := range qIdx {
				qs[i].l = low
			}
			return
		}

		mid := (low + high) >> 1
		var b, c []int
		for _, i := range idx {
			v := a[i]
			shift := bits.Len32(uint32(v/(low+1))) - 1
			if v>>shift < mid {
				b = append(b, i)
			}
			if v >= mid {
				if v > mid {
					c = append(c, i)
				}
				t.update(i+1, bits.Len32(uint32(v/mid))-bits.Len32(uint32(v/high)))
			}
		}

		var d, e []int
		for _, i := range qIdx {
			q := &qs[i]
			cnt := t.query(q.l, q.r)
			if cnt <= q.k {
				q.k -= cnt
				d = append(d, i)
			} else {
				e = append(e, i)
			}
		}

		for _, i := range idx {
			if a[i] >= mid {
				t.reset(i + 1)
			}
		}

		play(b, d, low, mid)
		play(c, e, mid, high)
	}
	play(idx, qIdx, 0, 1e5+1)
	var ans = make([]int, m)
	for i, q := range qs {
		ans[i] = q.l
	}
	return ans
}
