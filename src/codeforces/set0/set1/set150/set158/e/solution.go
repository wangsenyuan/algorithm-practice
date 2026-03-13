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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	calles := make([][]int, n)
	for i := range n {
		calles[i] = make([]int, 2)
		fmt.Fscan(reader, &calles[i][0], &calles[i][1])
	}
	return solve(calles, k)
}

const dayOfEnd = 86400

func solve(calles [][]int, k int) int {
	slices.SortFunc(calles, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	dp := make([]int, k+1)
	for i := range k + 1 {
		dp[i] = dayOfEnd
	}
	dp[0] = 0

	ndp := make([]int, k+1)

	var best int

	for _, cur := range calles {
		for j := range k + 1 {
			best = max(best, cur[0]-dp[j]-1)
			ndp[j] = dayOfEnd
		}

		for j := range k + 1 {
			ndp[j] = max(dp[j]+1, cur[0]) + cur[1] - 1
			if j > 0 {
				ndp[j] = min(ndp[j], dp[j-1])
			}
		}

		dp, ndp = ndp, dp
	}

	best = max(best, dayOfEnd-dp[k])

	return best
}
