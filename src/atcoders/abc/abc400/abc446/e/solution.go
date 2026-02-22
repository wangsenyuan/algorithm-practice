package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var m, a, b int
	fmt.Fscan(reader, &m, &a, &b)
	res := solve(m, a, b)
	fmt.Println(res)
}

func solve(m int, a int, b int) int {
	// 0 <= x, y < m
	// x和y不是m的倍数
	// bfs

	dp := make([][]int, m)
	marked := make([][]bool, m)
	for i := range m {
		dp[i] = make([]int, m)
		marked[i] = make([]bool, m)
	}

	var f func(s1 int, s2 int) int

	f = func(s1 int, s2 int) (res int) {
		if s1 == 0 || s2 == 0 {
			return 1
		}

		if dp[s1][s2] != 0 {
			return dp[s1][s2]
		}

		defer func() {
			dp[s1][s2] = res
		}()

		if marked[s1][s2] {
			res = -1
		} else {
			marked[s1][s2] = true
			s3 := (b*s1 + a*s2) % m
			res = f(s2, s3)
			// marked[s1][s2] = false
		}

		return
	}

	var res int

	for x := range m {
		for y := range m {
			if f(x, y) == -1 {
				res++
			}
		}
	}

	return res
}
