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
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

const N = 100010

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

var P [N]int

func init() {
	P[0] = 1
	for i := 1; i < N; i++ {
		P[i] = mul(P[i-1], 2)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	q := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &q[i])
	}
	return solve(p, q)
}

type pair struct {
	first  int
	second int
}

func (p pair) swap() pair {
	return pair{p.second, p.first}
}

func solve(p []int, q []int) []int {
	pos := []int{0, 0}
	n := len(p)
	res := make([]int, n)

	res[0] = add(P[p[0]], P[q[0]])

	cmpAndGet := func(i int, j int) pair {
		return pair{p[j], q[i-j]}
	}

	for i := 1; i < n; i++ {
		if p[i] > p[pos[0]] {
			pos[0] = i
		}
		if q[i] > q[pos[1]] {
			pos[1] = i
		}

		f := cmpAndGet(i, pos[0])
		s := cmpAndGet(i, i-pos[1]).swap()
		if f.first > s.first || f.first == s.first && f.second > s.second {
			res[i] = add(P[f.first], P[f.second])
		} else {
			res[i] = add(P[s.first], P[s.second])
		}

	}

	return res
}
