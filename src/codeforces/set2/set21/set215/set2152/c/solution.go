package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		for _, ans := range drive(reader) {
			fmt.Fprintln(writer, ans)
		}
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][2]int, q)
	for i := range queries {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(a, queries)
}

func solve(a []int, queries [][2]int) []int {
	// 010101, 这一种操作cost = 2, 00???0 操作cost = 1
	n := len(a)
	pref := make([]int, n+1)
	// dp[i] = 010101 pattern
	dp := make([]int, n)
	for i, v := range a {
		pref[i+1] = pref[i] + v
		dp[i] = 1
		if i > 0 && a[i] != a[i-1] {
			dp[i] += dp[i-1]
		}
	}

	check := func(l int, r int) int {
		if (r-l+1)%3 != 0 {
			return -1
		}
		if (pref[r+1]-pref[l])%3 != 0 {
			return -1
		}
		res := (r - l + 1) / 3
		if dp[r] >= (r - l + 1) {
			// 如果是 010101
			res++
		}
		return res
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		l, r := cur[0]-1, cur[1]-1
		ans[i] = check(l, r)
	}
	return ans
}
