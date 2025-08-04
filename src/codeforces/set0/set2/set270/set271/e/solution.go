package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
}

func process(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(m, a)
}

func solve(m int, a []int) int {
	var g int
	for _, x := range a {
		g = gcd(g, x-1)
	}

	var divs []int
	for i := 1; i <= g/i; i++ {
		if g%i == 0 {
			divs = append(divs, i)
			if i != g/i {
				divs = append(divs, g/i)
			}
		}
	}

	for i, d := range divs {
		for d%2 == 0 {
			d /= 2
		}
		divs[i] = d
	}

	sort.Ints(divs)
	divs = slices.Compact(divs)

	n := len(divs)

	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		d := divs[i]
		for j := 0; (1<<j)*d <= m; j++ {
			dp[i] += m - (1<<j)*d
		}
		for j := i + 1; j < n; j++ {
			if divs[j]%divs[i] == 0 {
				divs[i] -= divs[j]
			}
		}
	}

	var res int
	for _, v := range dp {
		res += v
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
