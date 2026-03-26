package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	fmt.Println(res)
}

func solve1(n int) int {
	var dp [6][4]int
	dp[5][3] = 1
	var ndp [6][4]int

	// 把k个东西分给n个人，不限制每个人得到多少个
	// = C(n + k - 1, k)
	for range n {
		for x := range 6 {
			for y := range 4 {
				for x1 := range x + 1 {
					for y1 := range y + 1 {
						ndp[x-x1][y-y1] += dp[x][y]
					}
				}
			}
		}
		for x := range 6 {
			for y := range 4 {
				dp[x][y] = ndp[x][y]
				ndp[x][y] = 0
			}
		}
	}
	return dp[0][0]
}

func solve(n int) int {

	calc := func(n int, k int) int {
		res := 1
		for i := range k {
			res *= (n - i)
			res /= (i + 1)
		}
		return res
	}

	x := calc(n+4, 5)
	y := calc(n+2, 3)

	return x * y
}
