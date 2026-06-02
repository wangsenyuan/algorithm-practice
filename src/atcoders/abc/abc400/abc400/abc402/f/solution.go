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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, mod int
	fmt.Fscan(reader, &n, &mod)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(mod, a)
}

func solve(mod int, a [][]int) int {
	n := len(a)

	if n == 1 {
		return a[0][0] % mod
	}

	add := func(u int, v int) int {
		return (u + v) % mod
	}

	mul := func(u int, v int) int {
		return (u * v) % mod
	}

	pw10 := make([]int, 2*n-1)
	pw10[0] = 1
	for i := 1; i < 2*n-1; i++ {
		pw10[i] = mul(pw10[i-1], 10)
	}

	for i := range n {
		for j := range n {
			a[i][j] = mul(a[i][j], pw10[2*n-2-(i+j)])
		}
	}

	g1 := make([][]int, n)

	for mask := range 1 << (n - 1) {
		var r, c int
		sum := a[r][c] % mod
		for i := range n - 1 {
			if (mask>>i)&1 == 1 {
				r++
			} else {
				c++
			}
			sum = add(sum, a[r][c])
		}
		g1[r] = append(g1[r], sum)
	}

	g2 := make([][]int, n)
	for mask := range 1 << (n - 1) {
		r, c := n-1, n-1
		var sum int
		for i := range n - 1 {
			sum = add(sum, a[r][c])
			if (mask>>i)&1 == 1 {
				r--
			} else {
				c--
			}
		}
		g2[r] = append(g2[r], sum)
	}

	var best int
	for i := range n {
		slices.Sort(g2[i])
		for _, u := range g1[i] {
			j := sort.SearchInts(g2[i], mod-u) - 1
			if j < 0 {
				j = len(g2[i]) - 1
			}
			best = max(best, add(u, g2[i][j]))
		}
	}

	return best
}
