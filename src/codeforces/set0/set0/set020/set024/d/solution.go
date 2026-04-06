package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	var x, y int
	fmt.Fscan(in, &n, &m)
	fmt.Fscan(in, &x, &y)
	fmt.Printf("%.10f\n", solve(n, m, x, y))
}

func solve(n, m, x, y int) float64 {
	if x == n {
		return 0
	}

	next := make([]float64, m)

	if m == 1 {
		for row := n - 1; row >= x; row-- {
			next[0] += 2.0
		}
		return next[0]
	}

	cur := make([]float64, m)
	a := make([]float64, m)
	b := make([]float64, m)

	for row := n - 1; row >= x; row-- {
		// cur[j] = a[j] * cur[j+1] + b[j]
		a[0] = 0.5
		b[0] = (3.0 + next[0]) / 2.0
		for j := 1; j < m-1; j++ {
			den := 3.0 - a[j-1]
			a[j] = 1.0 / den
			b[j] = (4.0 + next[j] + b[j-1]) / den
		}

		cur[m-1] = (3.0 + next[m-1] + b[m-2]) / (2.0 - a[m-2])
		for j := m - 2; j >= 0; j-- {
			cur[j] = a[j]*cur[j+1] + b[j]
		}

		copy(next, cur)
	}

	return next[y-1]
}
