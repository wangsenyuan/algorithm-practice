package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(reader, &s)
	res := solve(s)
	fmt.Println(res)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(s string) int {

	dp := make([]int, 3)
	ndp := make([]int, 3)
	for _, x := range s {
		d := int(x - 'a')

		clear(ndp)

		for d1 := range 3 {
			if d != d1 {
				ndp[d] = add(ndp[d], dp[d1])
			}
		}
		for d1 := range 3 {
			dp[d1] = add(dp[d1], ndp[d1])
		}
		dp[d] = add(dp[d], 1)
	}

	return add(dp[0], add(dp[1], dp[2]))
}
