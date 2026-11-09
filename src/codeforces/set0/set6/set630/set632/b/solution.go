package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	var s string
	fmt.Fscan(reader, &s)
	return solve(p, s)
}

func solve(p []int, s string) int64 {
	n := len(p)
	var best int
	pref := make([]int, n)
	var sum int
	for i := range n {
		sum += p[i]
		if s[i] == 'B' {
			best += p[i]
		}
		pref[i] = best
	}

	best = max(best, sum-best)

	get := func(i int) int {
		if i < 0 {
			return 0
		}
		return pref[i]
	}
	var sum2 int
	var suf int
	for i := n - 1; i >= 0; i-- {
		sum2 += p[i]
		sum -= p[i]
		if s[i] == 'A' {
			suf += p[i]
		}
		best = max(best, suf+get(i-1))
		best = max(best, sum2-suf+sum-get(i-1))
	}

	return int64(best)
}
