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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]int, t)
	for i := range t {
		var a, b string
		fmt.Fscan(reader, &a, &b)
		res[i] = solve(a, b)
	}
	return res
}

const inf = 1 << 60

func solve(a, b string) int {
	n := len(a)
	m := len(b)

	pa := play(a)
	pb := play(b)

	if pa[n] != pb[m] {
		return -1
	}

	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, m+1)
	}

	dp[0][0] = 0
	for i := range n {
		for j := range m {
			dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			if pa[i+1] == pb[j+1] {
				dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j]+1)
			}
		}
	}

	return dp[n][m]
}

func play(s string) []int {
	n := len(s)
	res := make([]int, n+1)
	for i := range n {
		res[i+1] = res[i] + int(s[i]-'0')
		if res[i+1] >= 10 {
			res[i+1] -= 10
		}
	}
	return res
}
