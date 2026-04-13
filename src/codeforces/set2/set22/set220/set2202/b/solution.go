package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	tc := readNums(reader)[0]
	for range tc {
		readString(reader)
		s := readString(reader)
		res := solve(s)
		if res {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	res := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		res[i], _ = strconv.Atoi(ss[i])
	}
	return res
}

func solve(s string) bool {
	n := len(s)
	if n&1 == 1 {
		s = "b" + s
	}
	// n 是偶数
	n = len(s)

	for i := 0; i < n; i += 2 {
		if s[i] != '?' && s[i+1] != '?' && s[i] == s[i+1] {
			return false
		}
	}

	return true
}
func solve1(s string) bool {
	n := len(s)
	dp := make([]map[int]bool, n)

	var f func(i int, j int) bool
	f = func(i int, j int) bool {
		if i == n {
			return true
		}
		if dp[i] != nil && dp[i][j] {
			return false
		}
		if dp[i] == nil {
			dp[i] = make(map[int]bool)
		}
		dp[i][j] = true

		if j&1 == 0 && (s[i] == 'a' || s[i] == '?') || (j&1 == 1 && (s[i] == 'b' || s[i] == '?')) {
			if f(i+1, j+1) {
				return true
			}
		}

		r := n - 1 - (i - j)

		if r&1 == 0 && (s[i] == 'a' || s[i] == '?') || (r&1 == 1 && (s[i] == 'b' || s[i] == '?')) {
			if f(i+1, j) {
				return true
			}
		}

		return false
	}

	return f(0, 0)
}
