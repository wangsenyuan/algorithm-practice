package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	r, s := drive(reader)
	fmt.Println(r, s)
}

func drive(reader *bufio.Reader) (r int, s int) {
	var n, m, a int
	fmt.Fscan(reader, &n, &m, &a)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	p := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &p[i])
	}
	return solve(a, b, p)
}

func solve(a int, b []int, p []int) (r int, s int) {
	slices.Sort(b)
	slices.Sort(p)

	m := len(p)
	n := len(b)

	check := func(x int) bool {
		if x > min(n, m) {
			return false
		}
		var sum int
		for i := range x {
			sum += max(p[i]-b[n-x+i], 0)
		}
		return sum <= a
	}

	r = min(n, m) + 1
	var l int
	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	r--
	// 首先要保证，所有人都可以租到车
	pay := make([]int, r)
	for i := range r {
		// 这部分是要从a中支付的
		x := max(p[i]-b[n-r+i], 0)
		a -= x
		// 剩下的部分是个人支付的
		pay[i] = p[i] - x
		s += pay[i]
	}

	for i := range r {
		x := min(a, pay[i])
		s -= x
		a -= x
	}

	return
}
