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
	tc := readNNums(reader)[0]
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}
func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	res := make([]int, len(ss))
	for i, v := range ss {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func drive(reader *bufio.Reader) string {
	dim := readNNums(reader)
	n, c := dim[0], dim[1]
	s := make([]string, n)
	for i := range n {
		s[i] = readString(reader)
	}
	return solve(c, s)
}

func solve(c int, s []string) string {
	n := len(s)

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
	}
	c--
	for i := range n {
		dp[i][c] = 1
	}

	wall := make([]int, n)
	for i := range n {
		wall[i] = -1
	}
	for i := range n {
		for j := range n {
			if s[i][j] == '#' {
				wall[j] = i
			}
		}
	}

	for r := n - 2; r >= 0; r-- {
		for j := range n {
			if dp[r][j] == 0 {
				// only process when it is not set yet
				if s[r][j] == '.' {
					for _, d := range []int{-1, 0, 1} {
						nc := j + d
						if nc >= 0 && nc < n && dp[r+1][nc] == 1 {
							dp[r][j] = 1
							break
						}
					}
				} else if wall[j] == r {
					// it is the lowest wall
					for _, d := range []int{-1, 0, 1} {
						nc := j + d
						if nc >= 0 && nc < n && dp[r+1][nc] == 1 {
							dp[r][j] = 1
							break
						}
					}
					if dp[r][j] == 1 {
						// good
						for i := r - 1; i >= 0; i-- {
							dp[i][j] = 1
						}
					}
				}
			}
		}
	}

	buf := make([]byte, n)
	for i := range n {
		if dp[0][i] == 1 {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}
	}

	return string(buf)
}

func abs(a int) int {
	return max(a, -a)
}
