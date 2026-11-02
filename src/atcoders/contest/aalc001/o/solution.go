package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int64 {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		queries[i] = []int{l, r}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int64 {
	n := len(a)
	left, right := minBounds(a)
	buckets := make([][]update, n+1)
	addRect := func(l1, l2, r1, r2, id int, v int64) {
		if l1 > l2 || r1 > r2 || v == 0 {
			return
		}
		buckets[l1] = append(buckets[l1], update{r1, r2, id, wint(v)})
		buckets[l2+1] = append(buckets[l2+1], update{r1, r2, id, wneg(wint(v))})
	}

	// i is the leftmost minimum on [s, t] iff left[i] < s <= i <= t < right[i].
	// A query cuts that rectangle in four ways. Each way is a polynomial in (ql, qr)
	// whose coefficients are constant on a rectangle of queries.
	for i := range n {
		if a[i] == 0 {
			continue
		}
		p := left[i] + 1
		q := right[i] - 1
		ai := int64(a[i])
		fullL := int64(i - p + 1)
		fullR := int64(q - i + 1)
		addRect(0, p, q, n-1, idA, ai*fullL*fullR)
		if q > i {
			c := ai * fullL
			addRect(0, p, i, q-1, idB1, c)
			addRect(0, p, i, q-1, idB2, c*int64(i))
		}
		if p < i {
			d := ai * fullR
			addRect(p+1, i, q, n-1, idC1, d*int64(i+1))
			addRect(p+1, i, q, n-1, idC2, d)
		}
		if p < i && q > i {
			addRect(p+1, i, i, q-1, idD0, ai)
			addRect(p+1, i, i, q-1, idD1, ai*int64(i))
			addRect(p+1, i, i, q-1, idD2, ai*int64(i+1))
			addRect(p+1, i, i, q-1, idD3, ai*int64(i)*int64(i+1))
		}
	}

	byL := make([][]query, n)
	for id, cur := range queries {
		byL[cur[0]-1] = append(byL[cur[0]-1], query{cur[1] - 1, id})
	}

	fw := make([]fenwick, idN)
	for i := range fw {
		fw[i] = newFenwick(n)
	}
	ans := make([]int64, len(queries))
	for ql := range n {
		for _, u := range buckets[ql] {
			fw[u.id].rangeAdd(u.r1, u.r2, u.val)
		}
		for _, qry := range byL[ql] {
			ans[qry.id] = combine(fw, ql, qry.qr).int64()
		}
	}
	return ans
}

const (
	idA = iota
	idB1
	idB2
	idC1
	idC2
	idD0
	idD1
	idD2
	idD3
	idN
)

func combine(fw []fenwick, ql, qr int) wide {
	a := fw[idA].prefix(qr)
	b1 := fw[idB1].prefix(qr)
	b2 := fw[idB2].prefix(qr)
	c1 := fw[idC1].prefix(qr)
	c2 := fw[idC2].prefix(qr)
	d0 := fw[idD0].prefix(qr)
	d1 := fw[idD1].prefix(qr)
	d2 := fw[idD2].prefix(qr)
	d3 := fw[idD3].prefix(qr)

	res := a
	res = wadd(res, wmul(b1, int64(qr+1)))
	res = wsub(res, b2)
	res = wadd(res, c1)
	res = wsub(res, wmul(c2, int64(ql)))
	res = wadd(res, wmul(d2, int64(qr+1)))
	res = wsub(res, d3)
	res = wsub(res, wmul(d0, int64(ql)*int64(qr+1)))
	res = wadd(res, wmul(d1, int64(ql)))
	return res
}

func minBounds(a []int) ([]int, []int) {
	n := len(a)
	left := make([]int, n)
	right := make([]int, n)
	stack := make([]int, 0, n)
	for i := range n {
		for len(stack) > 0 && a[stack[len(stack)-1]] > a[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && a[stack[len(stack)-1]] >= a[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return left, right
}

type update struct {
	r1, r2 int
	id     int
	val    wide
}

type query struct {
	qr int
	id int
}

type fenwick []wide

func newFenwick(n int) fenwick {
	return make(fenwick, n+3)
}

func (f fenwick) add(i int, v wide) {
	for i++; i < len(f); i += i & -i {
		f[i] = wadd(f[i], v)
	}
}

func (f fenwick) rangeAdd(l, r int, v wide) {
	f.add(l, v)
	f.add(r+1, wneg(v))
}

func (f fenwick) prefix(i int) wide {
	var s wide
	for i++; i > 0; i -= i & -i {
		s = wadd(s, f[i])
	}
	return s
}

// wide is a signed 128-bit integer. Products of the prefix sums exceed 2^63
// before they cancel down to the answer.
type wide struct {
	hi int64
	lo uint64
}

func wint(v int64) wide {
	if v >= 0 {
		return wide{0, uint64(v)}
	}
	return wneg(wide{0, uint64(-v)})
}

func (w wide) int64() int64 {
	return int64(w.lo)
}

func wadd(a, b wide) wide {
	lo, c := bits.Add64(a.lo, b.lo, 0)
	return wide{a.hi + b.hi + int64(c), lo}
}

func wsub(a, b wide) wide {
	lo, br := bits.Sub64(a.lo, b.lo, 0)
	return wide{a.hi - b.hi - int64(br), lo}
}

func wneg(a wide) wide {
	lo := ^a.lo
	hi := ^a.hi
	lo, c := bits.Add64(lo, 1, 0)
	return wide{hi + int64(c), lo}
}

func wmul(a wide, k int64) wide {
	if k == 0 || (a.hi == 0 && a.lo == 0) {
		return wide{}
	}
	neg := k < 0
	if neg {
		k = -k
	}
	if a.hi < 0 {
		neg = !neg
		a = wneg(a)
	}
	hi, lo := bits.Mul64(a.lo, uint64(k))
	hi += uint64(a.hi) * uint64(k)
	res := wide{int64(hi), lo}
	if neg {
		return wneg(res)
	}
	return res
}
