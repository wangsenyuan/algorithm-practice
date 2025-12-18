package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	return solve(k, a, b)
}

func solve(k int, a []int, b []int) bool {
	// 不需要整数，和k是无关的
	slices.Sort(a)
	slices.Reverse(a)
	slices.Sort(b)
	slices.Reverse(b)

	n := len(a)
	m := len(b)

	if n > m {
		// 让所有的相等
		return true
	}

	var i int
	// a[i] <= b[i]
	for i < n && i < m {
		if a[i] <= b[i] || i > 0 && b[i-1] == b[i] {
			// w[b[i-1]] = w[a[i-1]] = w[b[i]] => w[a[i]] = w[b[i]]
			i++
			continue
		}

		return true
	}

	return false
}
