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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

var factors = []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71}

func solve(a []int, b []int) int {
	n := len(a)
	c := make([]int, n)
	for i := range n {
		switch i {
		case 0:
			c[i] = gcd(a[i], a[i+1])
		case n - 1:
			c[i] = gcd(a[i-1], a[i])
		default:
			c[i] = lcm(gcd(a[i-1], a[i]), gcd(a[i], a[i+1]))
		}
		// 但是c[i] <= b[i]的时候，还是有可能等于a[i]的
		// 比如 [2, 2, 2], [2, 6, 9]
		if c[i] > b[i] {
			c[i] = a[i]
		}
	}

	// 然后在c[i]的基础上 *1 *2 *3进行处理?
	dp := make([]int, len(factors))
	for w := range len(factors) {
		if c[0]*factors[w] <= b[0] && c[0]*factors[w] != a[0] {
			dp[w] = 1
		}
	}
	ndp := make([]int, len(factors))
	for i := 1; i < n; i++ {
		g1 := gcd(a[i-1], a[i])
		for m1 := range len(factors) {
			x := c[i-1] * factors[m1]
			if x <= b[i-1] || x == a[i-1] {
				for m2 := range len(factors) {
					y := c[i] * factors[m2]
					if (y <= b[i] || y == a[i]) && gcd(x, y) == g1 {
						var add int
						if y != a[i] {
							add++
						}
						ndp[m2] = max(ndp[m2], dp[m1]+add)
					}
				}
			}
		}
		copy(dp, ndp)
		clear(ndp)
	}

	return slices.Max(dp)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	c := gcd(a, b)
	return a / c * b
}
