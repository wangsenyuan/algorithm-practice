package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Println(drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

const mod = 998244353

func add(a, b int) int {
	return (a + b) % mod
}

func mul(a, b int) int {
	return (a * b) % mod
}

func solve(s string) int {
	n := len(s)
	dp := make([][][][2]int, n+1)
	ndp := make([][][][2]int, n+1)
	for i := range n + 1 {
		dp[i] = make([][][2]int, n+1)
		ndp[i] = make([][][2]int, n+1)
		for j := range n + 1 {
			dp[i][j] = make([][2]int, 3)
			ndp[i][j] = make([][2]int, 3)
		}
	}

	dp[0][0][0][0] = 1

	var res int

	// 0表示 (, 1表示 )(, 2表示 有两个	 )(
	for _, c := range s {
		if c == '(' {
			for i := range n {
				for j := range i + 1 {
					for prev := range 2 {
						if prev == 0 {
							for k := range 3 {
								ndp[i+1][j+1][k][0] = add(ndp[i+1][j+1][k][0], dp[i][j][k][prev])
							}
						} else {
							for k := range 3 {
								nk := min(2, k+1)
								ndp[i+1][j+1][nk][1] = add(ndp[i+1][j+1][nk][1], dp[i][j][k][prev])
							}
						}
					}
				}
			}
		} else {
			// level down
			for i := 1; i < n; i++ {
				for j := i; j > 0; j-- {
					for k := range 3 {
						for prev := range 2 {
							ndp[i+1][j-1][k][1] = add(ndp[i+1][j-1][k][1], dp[i][j][k][prev])
						}
					}
				}
			}
		}
		for i := range n + 1 {
			for j := range i + 1 {
				for k := range 3 {
					for prev := range 2 {
						dp[i][j][k][prev] = add(dp[i][j][k][prev], ndp[i][j][k][prev])
						if prev == 1 && k == 2 && j == 0 {
							res = add(res, mul(ndp[i][j][k][prev], i-2))
						}
						ndp[i][j][k][prev] = 0
					}
				}
			}
		}
	}

	return res
}
