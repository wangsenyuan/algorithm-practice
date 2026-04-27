package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	return solve(a, c)
}

func solve(a []int, c []int) int {
	n := len(a)
	// 修改a[i]处的值，需要花费c[i]
	// 那么要找到最小值，使的a[i]非递减
	// 那么反过来，就是找到最大的，不变的部分，且这些部分的c[i]的sum最大
	dp := make([]int, n)
	var sum int
	for i := range n {
		dp[i] = c[i]
		sum += c[i]
		for j := range i {
			if a[j] <= a[i] {
				dp[i] = max(dp[i], dp[j]+c[i])
			}
		}
	}
	return sum - slices.Max(dp)
}
