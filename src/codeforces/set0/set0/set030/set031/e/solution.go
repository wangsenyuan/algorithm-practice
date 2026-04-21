package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

const inf = 1 << 60

func solve(s string) string {
	n := len(s)
	// 分配 n / 2
	h := n / 2
	pw := make([]int, h+1)
	pw[0] = 1
	for i := 1; i <= h; i++ {
		pw[i] = pw[i-1] * 10
	}

	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, h+1)
		for j := range h + 1 {
			dp[i][j] = -inf
		}
	}

	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		w := int(s[i-1] - '0')
		for j := 0; j <= h && j < i; j++ {
			// 分配s[i-1]给A或者B
			// 且已经分配给A j个字符
			if j+1 <= h {
				dp[i][j+1] = max(dp[i][j+1], dp[i-1][j]+w*pw[h-1-j])
			}
			// 如果分配给B, 那么之前B分配了(i - 1 - j)个字符,
			if i-1-j+1 <= h {
				dp[i][j] = max(dp[i][j], dp[i-1][j]+w*pw[h-1-(i-1-j)])
			}
		}
	}

	var buf []byte

	for i, j := n, h; i > 0; i-- {
		w := int(s[i-1] - '0')
		// 如果i必须分配给A
		if j > 0 && dp[i-1][j-1]+w*pw[h-j] == dp[i][j] {
			buf = append(buf, 'H')
			j--
		} else {
			buf = append(buf, 'M')
		}
	}

	slices.Reverse(buf)

	return string(buf)
}
