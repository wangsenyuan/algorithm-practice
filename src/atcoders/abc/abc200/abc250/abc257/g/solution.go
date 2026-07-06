package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var s, t string
	fmt.Fscan(reader, &s, &t)
	return solve(s, t)
}

func solve(s, t string) int {
	n := len(s)
	z := zFunction(s + "#" + t)
	k := len(z) + 1
	st := make(SegTree, 2*k)
	for i := range 2 * k {
		st[i] = inf
	}

	st.update(len(z), 0)

	for i := len(z) - 1; i >= n+1; i-- {
		if z[i] > 0 {
			w := st.query(i+1, i+z[i]+1)
			st.update(i, w+1)
		}
	}

	ans := st.query(n+1, n+2)

	if ans == inf {
		return -1
	}

	return ans
}

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

type SegTree []int

func (st SegTree) update(i int, v int) {
	n := len(st) / 2
	i += n
	st[i] = v
	for i > 1 {
		st[i>>1] = min(st[i], st[i^1])
		i >>= 1
	}
}

func (st SegTree) query(l int, r int) int {
	n := len(st) / 2
	l += n
	r += n
	res := inf
	for l < r {
		if l&1 == 1 {
			res = min(res, st[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, st[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

const inf = 1 << 60
