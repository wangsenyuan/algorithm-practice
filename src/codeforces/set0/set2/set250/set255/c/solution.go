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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	// dp[i][j] 表示(i, j)为前两个时的最优解
	// dp[i][j] = dp[j][k] + 1, 如果 a[j] = a[i] - q
	// a[k] = a[j] + q
	// => a[k] = a[i] ?
	arr := slices.Clone(a)
	slices.Sort(arr)
	arr = slices.Compact(arr)
	m := len(arr)
	n := len(a)
	// dp[i][j] = dp[j][k]
	dp := make([][]int, n)

	pos := make([]int, n)

	for j, v := range a {
		dp[j] = make([]int, m)
		pos[j] = sort.SearchInts(arr, v)

		for i := range m {
			dp[j][i] = 1
		}

		for l := range j {
			i := pos[l]
			dp[j][i] = max(dp[j][i], dp[l][pos[j]]+1)
		}
	}

	var res int
	for j := range n {
		for i := range m {
			res = max(res, dp[j][i])
		}
	}
	return res
}
