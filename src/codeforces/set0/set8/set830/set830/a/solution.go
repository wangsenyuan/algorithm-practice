package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, p int
	fmt.Fscan(reader, &n, &m, &p)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b, p)
}

const inf = 1 << 60

func solve(a []int, b []int, p int) int {
	sort.Ints(a)
	sort.Ints(b)
	n := len(a)
	m := len(b)
	// n <= m
	best := inf
	for l := 0; l+n <= m; l++ {
		var cur int
		for i := range n {
			cur = max(cur, abs(b[l+i]-a[i])+abs(b[l+i]-p))
		}
		best = min(best, cur)
	}
	return best
}

func abs(num int) int {
	return max(num, -num)
}
