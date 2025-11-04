package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	coins := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &coins[i])
	}
	return solve(k, coins)
}

func solve(k int, coins []int) []int {
	dp := make([][]bool, k+1)
	ndp := make([][]bool, k+1)
	for i := range k + 1 {
		dp[i] = make([]bool, k+1)
		ndp[i] = make([]bool, k+1)
	}
	dp[0][0] = true

	for _, c := range coins {
		for i := 0; i <= k; i++ {
			for j := 0; j <= k; j++ {
				if dp[i][j] {
					ndp[i][j] = true
					if i+c <= k {
						ndp[i+c][j] = true
						if j+c <= k {
							ndp[i+c][j+c] = true
						}
					}
				}
			}
		}
		for i := range k + 1 {
			for j := range k + 1 {
				dp[i][j] = ndp[i][j]
				ndp[i][j] = false
			}
		}
	}
	var res []int
	for j := range k + 1 {
		if dp[k][j] {
			res = append(res, j)
		}
	}
	return res
}
