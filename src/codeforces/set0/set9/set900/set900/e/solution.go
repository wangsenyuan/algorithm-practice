package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func process(reader *bufio.Reader) int {
	readNum(reader)
	s := readString(reader)
	m := readNum(reader)
	return solve(s, m)
}

type data struct {
	val          int
	modification int
}

func solve(s string, m int) int {
	n := len(s)
	dp := make([]data, n+1)
	fp := make([]int, n+1)

	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i]
		if s[i] == '?' {
			sum[i+1]++
		}
	}

	for i := n - 1; i >= 0; i-- {
		if s[i] == '?' || s[i] == 'a' {
			// 至少是1
			fp[i] = 1
			if i+1 < n && (s[i+1] == '?' || s[i+1] == 'b') {
				fp[i] = 2 + fp[i+2]
			}
		}
		if fp[i] >= m {
			if i+m == n {
				dp[i] = data{1, sum[i+m] - sum[i]}
			} else {
				dp[i] = data{1 + dp[i+m].val, sum[i+m] - sum[i] + dp[i+m].modification}
			}
		}
		if dp[i].val < dp[i+1].val || dp[i].val == dp[i+1].val && dp[i].modification > dp[i+1].modification {
			dp[i] = dp[i+1]
		}
	}

	return dp[0].modification
}
