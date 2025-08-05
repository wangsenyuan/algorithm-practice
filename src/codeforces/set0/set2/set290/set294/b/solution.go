package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	t := make([]int, n)
	w := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i], &w[i])
	}
	return solve(t, w)
}

func solve(t []int, w []int) int {

	n := len(t)
	var sum int
	for i := range n {
		sum += t[i]
	}

	// dp[x] = 表示在下面宽度是x时，上边的最小宽度
	dp := make([]int, sum+1)
	ndp := make([]int, sum+1)

	for i := range sum + 1 {
		dp[i] = sum
		ndp[i] = sum
	}
	dp[0] = 0

	for i := range n {
		for s := range sum + 1 {
			ndp[s] = sum + 1
		}
		for j := range sum + 1 {
			ndp[j] = min(ndp[j], dp[j]+w[i])
			if j+t[i] <= sum {
				ndp[j+t[i]] = min(ndp[j+t[i]], dp[j])
			}
		}
		copy(dp, ndp)
	}

	for i := range sum + 1 {
		if dp[i] <= i {
			return i
		}
	}

	return sum
}
