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
	fmt.Println(drive(reader))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

const inf = 1 << 60

func solve(s string, t string) int {
	p := kmp(t)
	n := len(s)
	m := len(t)

	dp := make([]int, m+1)
	for i := range m + 1 {
		dp[i] = -inf
	}

	dp[0] = 0

	for i := range n {
		for j := m; j >= 0; j-- {
			if j > 0 {
				dp[p[j-1]] = max(dp[p[j-1]], dp[j])
			}
			if j < m && (s[i] == '?' || s[i] == t[j]) {
				nj := j + 1
				if nj < m {
					dp[nj] = max(dp[nj], dp[j])
				} else {
					dp[nj] = max(dp[nj], dp[j]+1)
				}
			}
			if j > 0 {
				dp[j] = -inf
			}
		}
	}

	return slices.Max(dp)
}

func kmp(s string) []int {
	n := len(s)
	p := make([]int, n)
	for i := 1; i < n; i++ {
		j := p[i-1]
		for j > 0 && s[i] != s[j] {
			j = p[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		p[i] = j
	}
	return p
}
