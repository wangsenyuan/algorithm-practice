package main

import (
	"bufio"
	"fmt"
	"os"
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
	n := len(a)
	ans := 1
	for i := 1; i < n; i++ {
		if ans&1 == 1 && a[i-1] > a[i] {
			ans++
		}
		if ans&1 == 0 && a[i-1] < a[i] {
			ans++
		}
	}
	return ans
}

func solve1(a []int) int {
	n := len(a)
	dp := make([][2]int, n)

	var res int

	for i := range n {
		v := a[i]
		dp[i][0] = 1
		// dp[i][1] = 1
		dp[i][1] = -n

		for j := range i {
			if a[j] < v {
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			} else if a[j] > v {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}
		}
		res = max(res, dp[i][0], dp[i][1])
	}

	return res
}
