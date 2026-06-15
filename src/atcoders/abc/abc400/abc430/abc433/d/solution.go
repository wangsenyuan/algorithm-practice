package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, M int
	fmt.Fscan(reader, &n, &M)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, M)
}

func solve(a []int, M int) int {
	dp := make([]map[int]int, 11)
	base := make([]int, 11)
	for i := range 11 {
		dp[i] = make(map[int]int)
		if i == 0 {
			base[i] = 1
		} else {
			base[i] = 10 * base[i-1] % M
		}
	}

	for _, v := range a {
		s := strconv.Itoa(v)
		dp[len(s)][v%M]++
	}

	var res int
	for _, v := range a {
		if v == 0 {
			continue
		}
		for d := 1; d <= 10; d++ {
			// v % M * base[d] % M
			w := v % M * base[d] % M
			if w == 0 {
				res += dp[d][0]
			} else {
				res += dp[d][M-w]
			}
		}
	}

	return res
}
