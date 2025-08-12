package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
}

func solve(s string) []string {
	n := len(s)

	// 增加一个0的前缀
	dp := make([][2]int, n+1)
	dp[0][0] = 0
	dp[0][1] = 1
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			//要么在前面不借位的情况下 +1, 要么在前面产生借位的情况下, -1,   100 - 010 = 010
			// 这样子借位都没有了
			dp[i+1][0] = min(dp[i][0]+1, dp[i][1]+1)
			// 继续保持对后的借位, 比如 100 - 001 = 011 (最后仍然会产生1，但是目前不用操作)
			dp[i+1][1] = dp[i][1]
		} else {
			// s[i] == '0‘
			// 要么之前没有借位的情况下，这里啥也不用干
			// 要么在由借位的情况下，那么这里如果进行一次-1的话，借位还是会继续存在的
			// 如果原来没有借位的情况下， 当前位+1, 那么给后续产生借位
			dp[i+1][0] = dp[i][0]
			dp[i+1][1] = min(dp[i][0]+1, dp[i][1]+1)
		}
	}
	// 不能产生借位
	best := dp[n][0]
	borrow := 0

	var ans []string

	for i := n; i > 0; i-- {
		if s[i-1] == '1' {
			if borrow == 0 {
				// 必须操作
				if dp[i-1][0]+1 == best {
					ans = append(ans, fmt.Sprintf("+2^%d", n-i))
				} else {
					ans = append(ans, fmt.Sprintf("-2^%d", n-i))
					borrow = 1
				}
				best--
			}
			// else borrow = 1, keep
		} else {
			if borrow == 1 {
				if dp[i-1][1]+1 == best {
					ans = append(ans, fmt.Sprintf("-2^%d", n-i))
				} else {
					ans = append(ans, fmt.Sprintf("+2^%d", n-i))
					borrow = 0
				}
				best--
			}
			// else keep
		}
	}

	if borrow == 1 {
		ans = append(ans, fmt.Sprintf("+2^%d", n))
	}
	// asset borrow = 0, and best = 0
	slices.Reverse(ans)

	return ans
}
