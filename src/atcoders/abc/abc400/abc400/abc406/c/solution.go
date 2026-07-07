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

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(p)
}

func solve(p []int) int {
	n := len(p)
	dp := make([]int, n)
	var l int
	for i := range n {
		if i > 0 && p[i-1] > p[i] {
			l = 0
		}
		l++
		dp[i] = l
	}
	fp := make([]int, n+1)
	var r int
	var inc int
	var res int
	for i := n - 1; i >= 0; i-- {
		if i+1 < n && p[i] > p[i+1] {
			// r >= 1
			j := i + 1 + r
			if j < n {
				res += (dp[i] - 1) * fp[j]
			}
		}
		if i+1 < n && p[i] < p[i+1] {
			r = 0
		}
		r++
		if i+1 < n && p[i] > p[i+1] {
			inc = 0
		}
		inc++
		fp[i] = inc
	}
	return res
}
