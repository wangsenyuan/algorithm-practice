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
		if res {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func drive(reader *bufio.Reader) bool {
	var s, t string
	fmt.Fscan(reader, &s, &t)
	return solve(s, t)
}

func solve(s string, t string) bool {
	if check1(s, t) {
		return true
	}

	n := len(s)
	next := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		next[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			next[i][j] = n + 1
		}
	}

	for i := n - 1; i >= 0; i-- {
		copy(next[i], next[i+1])
		x := int(s[i] - 'a')
		next[i][x] = i
	}

	check := func(t1, t2 string) bool {
		dp := make([][]int, len(t1)+1)
		for i := 0; i <= len(t1); i++ {
			dp[i] = make([]int, len(t2)+1)
			for j := 0; j <= len(t2); j++ {
				dp[i][j] = len(s) + 10
			}
		}
		dp[0][0] = -1

		for i := 0; i <= len(t1); i++ {
			for j := 0; j <= len(t2); j++ {
				if i == 0 && j == 0 {
					continue
				}
				if i > 0 && dp[i-1][j]+1 < n {
					x := int(t1[i-1] - 'a')
					y := next[dp[i-1][j]+1][x]
					dp[i][j] = min(dp[i][j], y)
				}
				if j > 0 && dp[i][j-1]+1 < n {
					x := int(t2[j-1] - 'a')
					y := next[dp[i][j-1]+1][x]
					dp[i][j] = min(dp[i][j], y)
				}
			}
		}

		return dp[len(t1)][len(t2)] < len(s)
	}

	for i := 1; i < len(t); i++ {
		if check(t[:i], t[i:]) {
			return true
		}
	}
	return false
}

func check1(s, t string) bool {
	for i, j := 0, 0; i < len(t); i++ {
		for j < len(s) && s[j] != t[i] {
			j++
		}
		if j == len(s) {
			return false
		}
		j++
	}
	return true
}
