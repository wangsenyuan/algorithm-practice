package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
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
	// 只需要考虑递增序列就可以了
	dp := make([][]int, k+1)
	for i := range k + 1 {
		dp[i] = make([]int, k+1)
		for j := range k + 1 {
			dp[i][j] = -inf
		}
	}

	n := len(a)
	dp[0][0] = 0
	pos := make([]int, k+1)
	for x := 1; x <= k; x++ {
		pos[x] = -1
		for i := range n {
			if a[i] >= x {
				pos[x] = i
				break
			}
		}
	}
	for i := 1; i <= k; i++ {
		for j := 1; j <= i; j++ {
			if pos[j] >= 0 {
				for l := 0; l <= i-j; l++ {
					dp[i][j] = max(dp[i][j], dp[i-j][l]+(j-l)*(n-pos[j]))
				}
			}
		}
	}
	var ans int
	for i := range k + 1 {
		for j := range k + 1 {
			ans = max(ans, dp[i][j])
		}
	}

	return ans
}
