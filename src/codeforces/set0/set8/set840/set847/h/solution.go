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
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a []int) int {
	n := len(a)
	dp := make([]pair, n)
	dp[0] = pair{0, a[0]}
	for i := 1; i < n; i++ {
		dp[i].second = max(a[i], dp[i-1].second+1)
		dp[i].first = dp[i-1].first + dp[i].second - a[i]
	}

	res := dp[n-1].first
	val := a[n-1]
	var suf int
	for i := n - 2; i >= 0; i-- {
		// 如果递增到i, 从i+1开始递减
		tmp := dp[i].first + suf
		if dp[i].second <= val {
			tmp += val + 1 - dp[i].second
		}
		res = min(res, tmp)
		val = max(val+1, a[i])
		suf += val - a[i]
	}

	return res
}
