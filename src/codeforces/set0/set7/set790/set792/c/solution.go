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
	res, _ := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (string, string) {
	num := readString(reader)
	return solve(num), num
}

const inf = 1 << 60

func solve(num string) string {
	if num == "0" {
		return num
	}
	n := len(num)
	type state struct {
		val int
		del bool
	}
	dp := make([][]state, n+1)
	dp[0] = make([]state, 6)
	for x := range 6 {
		dp[0][x].val = inf
	}
	dp[0][0].val = 0

	for i := range n {
		dp[i+1] = make([]state, 6)
		// 删除这个数
		for x := range 6 {
			dp[i+1][x].val = dp[i][x].val + 1
			dp[i+1][x].del = true
		}

		// 且前面已经有非0的数了不删除这个数
		for x := 3; x < 6; x++ {
			sum := x % 3
			sum = (sum + int(num[i]-'0')) % 3
			y := 3 + sum
			if dp[i][x].val < dp[i+1][y].val {
				dp[i+1][y].val = dp[i][x].val
				dp[i+1][y].del = false
			}
		}
		if num[i] != '0' {
			y := int(num[i]-'0') % 3
			if i < dp[i+1][3+y].val {
				dp[i+1][3+y].val = i
				dp[i+1][3+y].del = false
			}
		}
	}
	if dp[n][3].val >= inf {
		if strings.Count(num, "0") > 0 {
			return "0"
		}
		return "-1"
	}
	// 还要重新构造出来
	// 只有 1 0这个状态
	flag := 3
	var buf []byte
	for i := n; i > 0; i-- {
		if dp[i][flag].del {
			continue
		}
		// keep it
		buf = append(buf, num[i-1])
		for x := range 6 {
			if (x+int(num[i-1]-'0'))%3 == flag%3 && dp[i-1][x].val == dp[i][flag].val {
				flag = x
				break
			}
		}
	}
	slices.Reverse(buf)

	for len(buf) > 1 && buf[0] == '0' {
		buf = buf[1:]
	}

	return string(buf)
}
