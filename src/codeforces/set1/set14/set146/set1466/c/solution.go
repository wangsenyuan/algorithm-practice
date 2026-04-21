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

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	tc := readNums(reader)[0]

	for range tc {
		s := readString(reader)
		res := solve(s)
		fmt.Fprintln(writer, res)
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
	for i, s := range ss {
		res[i], _ = strconv.Atoi(s)
	}
	return res
}

const inf = 1 << 60

func solve(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}
	dp := make([]int, 4)
	if s[0] == s[1] {
		dp[0] = inf
	} else {
		dp[0] = 0
	}
	dp[1] = 1
	dp[2] = 1
	dp[3] = 2

	for i := 2; i < n; i++ {
		ndp := []int{inf, inf, inf, inf}

		for mask := range 4 {
			newMask := (mask << 1) & 3
			// 如果不修改当前字符，那么必须保证s[i] != s[i-1] 和 s[i] != s[i-2]
			if (mask&1 == 1 || s[i] != s[i-1]) && (mask&2 == 2 || s[i] != s[i-2]) {
				ndp[newMask] = min(ndp[newMask], dp[mask])
			}
			// 这个始终是成立的
			ndp[newMask|1] = min(ndp[newMask|1], dp[mask]+1)
		}

		dp = ndp
	}

	return slices.Min(dp)
}

func solve1(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}
	if n == 2 {
		if s[0] == s[1] {
			return 1
		}
		return 0
	}
	change := make([]bool, n)
	var res int
	// 如果s[i] == s[i-1]，
	for i := 1; i < n; i++ {
		if i-1 >= 0 && s[i] == s[i-1] && !change[i-1] {
			res++
			change[i] = true
			continue
		}
		if i-2 >= 0 && s[i] == s[i-2] && !change[i-2] {
			res++
			change[i] = true
		}
	}

	return res
}
