package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

func process(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	games := make([][]int, m)
	for i := range m {
		var l, r, x int
		fmt.Fscan(reader, &l, &r, &x)
		games[i] = []int{l, r, x}
	}
	return solve(n, games)
}

func solve(n int, games [][]int) []int {
	next := make([]int, n+2)
	for i := 0; i <= n+1; i++ {
		next[i] = i
	}

	var find func(i int) int

	find = func(i int) int {
		if next[i] != i {
			next[i] = find(next[i])
		}
		return next[i]
	}

	ans := make([]int, n+1)

	for _, cur := range games {
		l, r, x := cur[0], cur[1], cur[2]
		for l <= r {
			j := find(l)
			if j > r {
				break
			}
			if j == x {
				l = j + 1
				continue
			}
			// j != x
			ans[j] = x
			if j < x {
				next[j] = x
			} else {
				// j > x
				next[j] = find(r + 1)
			}
			l = j + 1
		}
	}
	return ans[1:]
}
