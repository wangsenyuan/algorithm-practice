package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	s := readString(reader)
	tc, _ := strconv.Atoi(s)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	readString(reader)
	a := make([]string, 2)
	for i := range 2 {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) int {
	n := len(a[0])

	dp := make([]int, n)

	for i := range n {
		if i > 0 {
			dp[i] = dp[i-1]
		}
		if a[0][i] == '.' && a[1][i] == '.' {
			if i == 0 || a[0][i-1] == 'x' && a[1][i-1] == 'x' {
				dp[i]++
			}
		} else if a[0][i] == '.' {
			// a[1][i] == 'x'
			if i == 0 || a[0][i-1] == 'x' {
				dp[i]++
			}
		} else if a[1][i] == '.' {
			if i == 0 || a[1][i-1] == 'x' {
				dp[i]++
			}
		}
	}

	if dp[n-1] > 4 {
		return 0
	}
	var res int

	get := func(i int) int {
		if i < 0 {
			return 0
		}
		return dp[i]
	}

	var fp int
	for i := n - 1; i >= 0; i-- {
		if a[0][i] == '.' && a[1][i] == '.' {
			// 在a[1][i]处放置
			for r := range 2 {
				if (i == n-1 || a[r][i+1] == 'x') && (i == 0 || a[r][i-1] == 'x') && get(i-1)+fp+1 == 3 {
					res++
				}
			}
		} else if a[0][i] == '.' || a[1][i] == '.' {
			if get(i-1)+fp == 3 {
				res++
			}
		}

		if a[0][i] == '.' && a[1][i] == '.' {
			if i == n-1 || a[0][i+1] == 'x' && a[1][i+1] == 'x' {
				fp++
			}
		} else if a[0][i] == '.' {
			if i == n-1 || a[0][i+1] == 'x' {
				fp++
			}
		} else if a[1][i] == '.' {
			if i == n-1 || a[1][i+1] == 'x' {
				fp++
			}
		}
	}

	return res
}
