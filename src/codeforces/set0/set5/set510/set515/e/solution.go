package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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
	d := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &d[i])
	}
	h := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &h[i])
	}
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		for j := range 2 {
			fmt.Fscan(reader, &queries[i][j])
		}
	}
	return solve(d, h, queries)
}

type data struct {
	val  int
	pref [2]int // 2 * h[i] - pref[i], 2 * h[i] + pref[i]
}

func merge(a data, b data) (res data) {
	res.val = max(a.val, b.val)
	res.val = max(res.val, b.pref[1]+a.pref[0])
	for j := range 2 {
		res.pref[j] = max(a.pref[j], b.pref[j])
	}

	return res
}

type seg []data

const inf = 1 << 60

func NewSegTree(h []int, pref []int) seg {
	n := len(h)
	res := make(seg, 4*n)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			res[i].val = -inf
			res[i].pref[0] = 2*h[l] - pref[l]
			// pref[l+1] 包括d[l]
			res[i].pref[1] = 2*h[l] + pref[l]
			return
		}
		mid := (l + r) / 2
		f(2*i+1, l, mid)
		f(2*i+2, mid+1, r)
		res[i] = merge(res[i*2+1], res[i*2+2])
	}
	f(0, 0, n-1)
	return res
}

func (tr seg) Query(L int, R int) int {
	var f func(i int, l int, r int) data
	f = func(i int, l int, r int) data {
		if L <= l && r <= R {
			return tr[i]
		}
		mid := (l + r) >> 1
		if R <= mid {
			return f(2*i+1, l, mid)
		}
		if mid < L {
			return f(2*i+2, mid+1, r)
		}
		a := f(i*2+1, l, mid)
		b := f(i*2+2, mid+1, r)
		return merge(a, b)
	}
	res := f(0, 0, len(tr)/4-1)
	return res.val
}

func solve(d []int, h []int, queries [][]int) []int {
	n := len(d)

	h2 := make([]int, 2*n)
	copy(h2, h)
	copy(h2[n:], h)

	pref := make([]int, 2*n)
	for i := 1; i < 2*n; i++ {
		pref[i] = pref[i-1] + d[(i-1)%n]
	}

	tr := NewSegTree(h2, pref)

	ans := make([]int, len(queries))

	for i, cur := range queries {
		a, b := cur[0]-1, cur[1]-1
		if a <= b {
			// 这个地方有点问题， 这里没有包括，跨过0的那些
			ans[i] = tr.Query(b+1, a+n-1)
		} else {
			// b < a
			ans[i] = tr.Query(b+1, a-1)
		}
	}

	return ans
}
