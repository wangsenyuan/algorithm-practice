package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var p, q, n int
	fmt.Fscan(reader, &p, &q, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(p, q, a)
}

func solve(p int, q int, a []int) bool {
	// p / q = a1 + 1/(a2 + 1/(a3 + 1/(... + 1/a_n)))
	w := gcd(p, q)
	p /= w
	q /= w
	// 主要的问题是，会不会溢出？
	n := len(a)
	for i := 0; i < n; i++ {
		v := a[i]
		if i == n-1 {
			return p == v && q == 1
		}

		if p/q <= v && p <= q*v {
			return false
		}
		// p / q >= v + ...
		x := p - v*q
		y := q
		z := gcd(x, y)
		x /= z
		y /= z
		p, q = y, x
	}

	return false
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
