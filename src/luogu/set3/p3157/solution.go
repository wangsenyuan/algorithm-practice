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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	del := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &del[i])
	}
	return solve(a, del)
}

func solve(input []int, del []int) []int {
	n := len(input)
	type pair struct{ i, t, inv int }
	a := make([]pair, n+1)
	f := make(fenwick, n+1)
	var ans int
	for i, v := range input {
		ans += f.pre(n) - f.pre(v)
		f.update(v, 1)
		a[v].i = i
	}
	m := len(del)
	for t := m; t > 0; t-- {
		a[del[m-t]].t = t
	}

	f = make(fenwick, m+2)
	var play func(int, int)
	play = func(l, r int) {
		if l+1 == r {
			return
		}
		mid := (l + r) >> 1
		play(l, mid)
		play(mid, r)

		v := l
		for w := mid; w < r; w++ {
			for ; v < mid && a[v].i > a[w].i; v++ {
				f.update(a[v].t+1, 1)
			}
			a[w].inv += f.pre(a[w].t)
		}
		for v--; v >= l; v-- {
			f.update(a[v].t+1, -1)
		}

		w := r - 1
		for v := mid - 1; v >= l; v-- {
			for ; w >= mid && a[w].i < a[v].i; w-- {
				f.update(a[w].t+1, 1)
			}
			a[v].inv += f.pre(a[v].t)
		}
		for w++; w < r; w++ {
			f.update(a[w].t+1, -1)
		}

		slices.SortFunc(a[l:r], func(a, b pair) int { return b.i - a.i })
	}

	play(1, n+1)

	inv := make([]int, m)
	for _, p := range a[1:] {
		if p.t > 0 {
			inv[m-p.t] = p.inv
		}
	}
	var res []int
	for _, c := range inv {
		res = append(res, ans)
		ans -= c
	}
	return res
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
