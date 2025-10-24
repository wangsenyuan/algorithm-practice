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
	res := drive(reader)
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, d int
	fmt.Fscan(reader, &n, &d)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, d)
}

func solve(a []int, d int) []int {
	mx := slices.Max(a)
	t := NewTree(mx + 1)
	seen := make([]bool, mx+1)

	ans := make([]int, len(a))
	for i, v := range a {
		if !seen[v] {
			seen[v] = true
			cnt := t.Query(v, min(v+d, mx)).cnt
			t.Set(v, cnt)
			if v > 1 {
				t.Update(max(1, v-d), v-1, 1)
			}
		} else {
			seen[v] = false
			t.Set(v, -1)
			if v > 1 {
				t.Update(max(1, v-d), v-1, -1)
			}
		}
		ans[i] = (t.val[0].s2 - t.val[0].s1) / 2
	}

	return ans
}

type data struct {
	cnt int
	s1  int
	s2  int
}

func (this data) merge(that data) data {
	return data{this.cnt + that.cnt, this.s1 + that.s1, this.s2 + that.s2}
}

type Tree struct {
	val  []data
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	val := make([]data, 4*n)
	lazy := make([]int, 4*n)
	return &Tree{val, lazy, n}
}

func (t *Tree) apply(p int, v int) {
	// update s2
	cur := &t.val[p]
	cur.s2 += cur.s1*v*2 + cur.cnt*v*v
	cur.s1 += cur.cnt * v
	t.lazy[p] += v
}

func (t *Tree) push(p int) {
	if t.lazy[p] != 0 {
		t.apply(2*p+1, t.lazy[p])
		t.apply(2*p+2, t.lazy[p])
		t.lazy[p] = 0
	}
}

func (t *Tree) pull(p int) {
	t.val[p] = t.val[2*p+1].merge(t.val[2*p+2])
}

func (t *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.apply(i, v)
			return
		}
		t.push(i)
		mid := (l + r) / 2
		if L <= mid {
			f(2*i+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(2*i+2, mid+1, r, max(mid+1, L), R)
		}
		t.pull(i)
	}
	f(0, 0, t.sz-1, L, R)
}

func (t *Tree) Set(p int, cnt int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			if cnt < 0 {
				t.val[i] = data{}
			} else {
				t.val[i] = data{1, cnt, cnt * cnt}
			}
			return
		}
		t.push(i)
		mid := (l + r) / 2
		if p <= mid {
			f(2*i+1, l, mid)
		} else {
			f(2*i+2, mid+1, r)
		}
		t.pull(i)
	}
	f(0, 0, t.sz-1)
}

func (t *Tree) Query(L int, R int) data {
	var f func(i int, l int, r int, L int, R int) data
	f = func(i int, l int, r int, L int, R int) data {
		if l == L && r == R {
			return t.val[i]
		}
		t.push(i)
		mid := (l + r) / 2
		if R <= mid {
			return f(2*i+1, l, mid, L, min(mid, R))
		}
		if mid < L {
			return f(2*i+2, mid+1, r, max(mid+1, L), R)
		}
		return f(2*i+1, l, mid, L, min(mid, R)).merge(f(2*i+2, mid+1, r, max(mid+1, L), R))
	}
	return f(0, 0, t.sz-1, L, R)
}
