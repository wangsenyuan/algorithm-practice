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
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	sn := readString(reader)
	n, _ := strconv.Atoi(sn)
	s := readString(reader)
	return solve(n, s)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(n int, s string) int {
	if len(s) > 2*n {
		return 0
	}
	// dp[i][level] = 包含s[:i], 且level时的方案数
	dp := make([][]int, len(s)+1)
	ndp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		ndp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	get := func(c byte) int {
		if c == '(' {
			return 0
		}
		return 1
	}

	p := kmp(s)

	for range 2 * n {
		for i := range len(s) + 1 {
			for level := range n + 1 {
				if dp[i][level] == 0 {
					continue
				}
				for c := range 2 {
					nextLevel := level
					if c == 0 {
						nextLevel++
					} else {
						nextLevel--
					}
					if nextLevel < 0 || nextLevel > n {
						continue
					}
					i1 := i
					if i1 < len(s) {
						for i1 > 0 && get(s[i1]) != c {
							i1 = p[i1-1]
						}

						if get(s[i1]) == c {
							i1++
						}
					}

					ndp[i1][nextLevel] = add(ndp[i1][nextLevel], dp[i][level])
				}
			}
		}

		for i := range len(s) + 1 {
			for level := range n + 1 {
				dp[i][level] = ndp[i][level]
				ndp[i][level] = 0
			}
		}
	}

	return dp[len(s)][0]
}

func kmp(s string) []int {
	n := len(s)
	res := make([]int, n)
	for i := 1; i < n; i++ {
		j := res[i-1]
		for j > 0 && s[i] != s[j] {
			j = res[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		res[i] = j
	}
	return res
}
