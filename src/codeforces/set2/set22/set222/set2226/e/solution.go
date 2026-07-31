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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		for i, v := range res {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	n := len(a)
	ans := make([]int, n)
	mx := n
	for _, x := range a {
		mx = max(mx, x)
	}
	freq := make([]int, mx+1)

	tr := BuildTree(n + 2)
	tr.Set(0, -1)
	var mex int

	for i, x := range a {
		if freq[x] > 0 || x > mex {
			if x > 0 {
				tr.Add(0, min((x-1)/2, n+1), 1)
			}
		} else {
			tr.Set(x, inf)
			tr.Add(0, x, 1)
		}
		freq[x]++

		for tr.Query(0, mex) >= 0 {
			mex++
			tr.Add(0, mex, -1)
			if freq[mex] > 0 {
				tr.Set(mex, inf)
				tr.Add(0, (mex-1)/2, -1)
				tr.Add(0, mex, 1)
			}
		}
		ans[i] = mex
	}

	return ans
}

type Tree struct {
	n    int
	val  []int
	lazy []int
}

const inf = 1 << 60

func BuildTree(n int) *Tree {
	return &Tree{
		n:    n,
		val:  make([]int, 4*n),
		lazy: make([]int, 4*n),
	}
}

func (t *Tree) push(i int) {
	if t.lazy[i] == 0 {
		return
	}
	for _, j := range []int{i*2 + 1, i*2 + 2} {
		t.val[j] += t.lazy[i]
		t.lazy[j] += t.lazy[i]
	}
	t.lazy[i] = 0
}

func (t *Tree) Add(L int, R int, v int) {
	var add func(int, int, int)
	add = func(i, l, r int) {
		if r < L || R < l {
			return
		}
		if L <= l && r <= R {
			t.val[i] += v
			t.lazy[i] += v
			return
		}
		t.push(i)
		mid := (l + r) / 2
		add(i*2+1, l, mid)
		add(i*2+2, mid+1, r)
		t.val[i] = min(t.val[i*2+1], t.val[i*2+2])
	}
	add(0, 0, t.n-1)
}

func (t *Tree) Set(p int, v int) {
	var set func(int, int, int)
	set = func(i, l, r int) {
		if l == r {
			t.val[i] = v
			t.lazy[i] = 0
			return
		}
		t.push(i)
		mid := (l + r) / 2
		if p <= mid {
			set(i*2+1, l, mid)
		} else {
			set(i*2+2, mid+1, r)
		}
		t.val[i] = min(t.val[i*2+1], t.val[i*2+2])
	}
	set(0, 0, t.n-1)
}

func (t *Tree) Query(L int, R int) int {
	var query func(int, int, int) int
	query = func(i, l, r int) int {
		if r < L || R < l {
			return inf
		}
		if L <= l && r <= R {
			return t.val[i]
		}
		t.push(i)
		mid := (l + r) / 2
		return min(query(i*2+1, l, mid), query(i*2+2, mid+1, r))
	}
	return query(0, 0, t.n-1)
}
