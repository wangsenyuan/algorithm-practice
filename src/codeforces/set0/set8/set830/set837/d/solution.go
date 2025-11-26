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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

const inf = 1 << 60

func solve(k int, a []int) int {
	n := len(a)

	dp := make([][]int, k+1)
	ndp := make([][]int, k+1)

	W := n * 30

	for i := range k + 1 {
		dp[i] = make([]int, W+1)
		ndp[i] = make([]int, W+1)
		for j := range W + 1 {
			dp[i][j] = -inf
			ndp[i][j] = -inf
		}
	}

	dp[0][0] = 0

	for _, v := range a {
		cnt := make([]int, 2)
		for v%10 == 0 {
			cnt[0]++
			cnt[1]++
			v /= 10
		}
		if v%2 == 0 {
			for v%2 == 0 {
				cnt[0]++
				v /= 2
			}
		} else if v%5 == 0 {
			for v%5 == 0 {
				cnt[1]++
				v /= 5
			}
		}

		for j := range k {
			for i := 0; i+cnt[1] <= W; i++ {
				ndp[j+1][i+cnt[1]] = max(ndp[j+1][i+cnt[1]], dp[j][i]+cnt[0])
			}
		}
		for i := range k + 1 {
			for j := range W + 1 {
				dp[i][j] = max(dp[i][j], ndp[i][j])
				ndp[i][j] = -inf
			}
		}
	}

	ans := -inf
	for i := range W + 1 {
		tmp := min(i, dp[k][i])
		ans = max(ans, tmp)
	}
	return ans
}
