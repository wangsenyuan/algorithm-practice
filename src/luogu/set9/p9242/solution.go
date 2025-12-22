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

func solve(a []int) int {
	dp := make([]int, 10)

	for _, v := range a {
		x := v % 10
		var y int
		for v > 0 {
			y = v % 10
			v /= 10
		}

		dp[x] = max(dp[x], dp[y]+1)
	}

	return len(a) - slices.Max(dp)
}
