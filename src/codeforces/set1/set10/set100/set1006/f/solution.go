package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(k, a)
}

func solve(k int, a [][]int) int {
	n := len(a)
	m := len(a[0])

	dp := make([][]map[int]int, n)
	fp := make([][]map[int]int, n)
	for i := range n {
		dp[i] = make([]map[int]int, m)
		fp[i] = make([]map[int]int, m)
		for j := range m {
			dp[i][j] = make(map[int]int)
			fp[i][j] = make(map[int]int)
		}
	}
	dp[0][0][a[0][0]] = 1
	h1 := (n + m - 1) / 2
	for i := range n {
		for j := range m {
			if i+j > h1 {
				break
			}
			if i+j == 0 {
				continue
			}
			if i > 0 {
				for x, y := range dp[i-1][j] {
					dp[i][j][x^a[i][j]] += y
				}
			}
			if j > 0 {
				for x, y := range dp[i][j-1] {
					dp[i][j][x^a[i][j]] += y
				}
			}
		}
	}
	h2 := n + m - 1 - h1
	fp[n-1][m-1][a[n-1][m-1]] = 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if n-1-i+m-1-j > h2 {
				break
			}
			if i+1 < n {
				for x, y := range fp[i+1][j] {
					fp[i][j][x^a[i][j]] += y
				}
			}
			if j+1 < m {
				for x, y := range fp[i][j+1] {
					fp[i][j][x^a[i][j]] += y
				}
			}
		}
	}

	var res int
	for i := range n {
		for j := range m {
			if i+j == h1 {
				if i == n-1 && j == m-1 {
					res += dp[i][j][k]
					continue
				}
				for x, y := range dp[i][j] {
					if i+1 < n {
						res += y * fp[i+1][j][x^k]
					}
					if j+1 < m {
						res += y * fp[i][j+1][x^k]
					}
				}
			}
		}
	}
	return res
}
