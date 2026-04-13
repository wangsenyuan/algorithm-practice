package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum, p := solve(n)
	fmt.Println(sum)
	s := fmt.Sprintf("%v", p)
	fmt.Println(s[1 : len(s)-1])
}

func solve(n int) (sum int, p []int) {
	p = make([]int, n+1)

	sum = n * (n + 1)
	for n > 0 {
		b := bits.Len(uint(n)) - 1
		m := n
		for m > 0 && (m>>b)&1 == 1 {
			m--
		}
		m++
		p[m] = m - 1
		p[m-1] = m
		j := m - 2
		for i := m + 1; i <= n; i++ {
			p[i] = j
			p[j] = i
			j--
		}

		n = j
	}

	return
}
