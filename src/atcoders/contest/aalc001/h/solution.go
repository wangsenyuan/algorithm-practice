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

	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var pos, x int
			fmt.Fscan(reader, &pos, &x)
			queries[i] = []int{t, pos, x}
		} else {
			var l, r, x int
			fmt.Fscan(reader, &l, &r, &x)
			queries[i] = []int{t, l, r, x}
		}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	s := make(fenwick, n+2)
	z := make(fenwick, n+2)

	for i := range n {
		s.update(i+1, a[i])
		if a[i] == 0 {
			z.update(i+1, 1)
		}
	}

	var ans []int
	const mod = 998244353
	for _, cur := range queries {
		if cur[0] == 1 {
			i, x := cur[1], cur[2]
			if a[i-1] == 0 {
				z.update(i, -1)
			}
			if x == 0 {
				z.update(i, 1)
			}
			s.update(i, x-a[i-1])
			a[i-1] = x
		} else {
			l, r := cur[1], cur[2]
			u, v := z.pre(l-1), z.pre(r)
			x := cur[3]
			if u == v {
				sum := s.query(l, r) + x
				ans = append(ans, sum)
			} else {
				j := z.kth(v)
				sum := (x+s.query(l, j))%mod + s.query(j, r)
				ans = append(ans, sum)
			}
		}
	}

	return ans
}

type fenwick []int

func (f fenwick) update(i int, v int) {
	for i < len(f) {
		f[i] += v
		i += i & -i
	}
}

func (f fenwick) pre(i int) int {
	var res int
	for i > 0 {
		res += f[i]
		i -= i & -i
	}
	return res
}

func (f fenwick) query(l int, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func (f fenwick) kth(k int) int {
	var i int
	for bit := 1 << 17; bit > 0; bit >>= 1 {
		nxt := i | bit
		if nxt < len(f) && f[nxt] < k {
			k -= f[nxt]
			i = nxt
		}
	}
	return i + 1
}
