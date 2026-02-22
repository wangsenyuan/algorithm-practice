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
	// 如果前面有个v,
	arr := slices.Clone(a)
	slices.Sort(arr)
	arr = slices.Compact(arr)

	m := len(arr)
	dp := make([]int, m)

	for _, v := range a {
		i := sort.SearchInts(arr, v)
		if i == 0 || arr[i-1] != v-1 {
			dp[i] = max(1, dp[i])
		} else {
			dp[i] = max(dp[i-1]+1, dp[i])
		}
	}
	return slices.Max(dp)
}
