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

const inf = 1 << 60

func solve(a []int) int {
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = -inf

	slices.Sort(a)

	n := len(a)
	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		cnt := i - j
		ndp := make([]int, 2)
		// 当前不选择时的最优解
		ndp[0] = max(dp[0], dp[1])
		if j == 0 || a[j] > a[j-1]+1 {
			// 前面的选择对当前选择没有关系
			ndp[1] = max(dp[0], dp[1]) + cnt*a[j]
		} else {
			ndp[1] = dp[0] + cnt*a[j]
		}
		copy(dp, ndp)
	}

	return max(dp[0], dp[1])
}
