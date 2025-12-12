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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const H = 20

const X = 5010

const inf = 1 << 30

func solve(a []int) int {
	n := len(a)

	mx := slices.Max(a)
	first := make([]int, mx+1)
	last := make([]int, mx+1)
	for i := range mx + 1 {
		first[i] = -1
		last[i] = -1
	}

	for i := range n {
		if first[a[i]] == -1 {
			first[a[i]] = i
		}
		last[a[i]] = i
	}

	dp := make([]int, n+1)

	dp[0] = 0

	for i := range n {
		dp[i+1] = dp[i]
		if i == last[a[i]] {
			l := first[a[i]]
			var sum int
			// 在区间[j..i]中间的，都必须被选中，不能只选一部分
			for j := i; j >= 0; j-- {
				if last[a[j]] > i {
					break
				}
				l = min(l, first[a[j]])
				if j == first[a[j]] {
					sum ^= a[j]
				}
				if j == l {
					dp[i+1] = max(dp[i+1], dp[l]+sum)
				}
			}
		}
	}

	return dp[n]
}
