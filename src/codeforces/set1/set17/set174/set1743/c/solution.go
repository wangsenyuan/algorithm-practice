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

func drive(reader *bufio.Reader) int64 {
	var n int
	var s string
	fmt.Fscan(reader, &n, &s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(s, a)
}

const inf = 1 << 60

func solve(s string, a []int) int64 {
	n := len(s)
	var res int
	for i := 0; i < n; {
		if s[i] == '0' {
			i++
			continue
		}
		j := i
		minVal := a[i]
		for i < n && s[i] == '1' {
			minVal = min(minVal, a[i])
			res += a[i]
			i++
		}
		if j > 0 {
			res += a[j-1]
			minVal = min(minVal, a[j-1])
			res -= minVal
		}
	}
	return int64(res)
}

func solve1(s string, a []int) int64 {
	n := len(s)
	dp := make([][2]int, n+1)
	dp[n][1] = -inf

	// dp[i][0] 表示没有盖子多出来
	// dp[i][1] 表示还有个盖子可以用
	for i := n - 1; i >= 0; i-- {
		dp[i][0] = max(dp[i+1][0], dp[i+1][1])
		dp[i][1] = -inf
		for d := range 2 {
			if dp[i+1][d] >= 0 {
				x := int(s[i] - '0')
				if x == 1 || d == 1 {
					dp[i][0] = max(dp[i][0], dp[i+1][d]+a[i])
				}
				if x == 1 && d == 1 {
					dp[i][1] = max(dp[i][1], dp[i+1][d]+a[i])
				}
				if x == 1 {
					// 把盖子留给下一位
					dp[i][1] = max(dp[i][1], dp[i+1][d])
				}
			}
		}
	}

	return int64(dp[0][0])
}
