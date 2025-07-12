package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	s := readString(reader)
	return solve(n, s)
}

const inf = 1 << 60

func solve(n int, s string) int {
	m := len(s)

	ln := len(strconv.Itoa(n))

	pn := make([]int, m+1)
	p10 := make([]int, m+1)
	pn[0] = 1
	p10[0] = 1

	for i := 1; i <= m; i++ {
		pn[i] = inf

		if pn[i-1] < inf/n {
			pn[i] = pn[i-1] * n
		}
		p10[i] = inf
		if p10[i-1] < inf/10 {
			p10[i] = p10[i-1] * 10
		}
	}

	dp := make([][]int, m+1)
	for i := range m + 1 {
		dp[i] = make([]int, m+1)
		for j := range m + 1 {
			dp[i][j] = inf
		}
	}
	dp[m][0] = 0

	// k, k-1, k - 2, ... 1, 0

	for i := m - 1; i >= 0; i-- {
		var num int
		for j := 1; i-j+1 >= 0 && j <= ln; j++ {
			num += int(s[i-j+1]-'0') * p10[j-1]
			if num >= n {
				break
			}
			if s[i-j+1] == '0' && j > 1 {
				// 0开始的不可以
				continue
			}
			for k, v := range dp[i+1] {
				if v == inf || pn[k] == inf || num > inf/pn[k] {
					continue
				}

				dp[i-j+1][k+1] = min(dp[i-j+1][k+1], v+num*pn[k])
			}
		}
	}
	return slices.Min(dp[0])
}
