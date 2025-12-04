package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
}

func solve(n int) int {

	var ds []int
	for i := n; i > 0; i /= 10 {
		ds = append(ds, i%10)
	}

	// 好像搞的有点复杂了
	slices.Reverse(ds)

	dp := make([][2]int, 1<<10)
	ndp := make([][2]int, 1<<10)

	dp[0][1] = 1

	update := func(s int, e int, v int, x int) {
		if v == 0 {
			return
		}
		for j := range 10 {
			if e == 1 && j > x {
				break
			}
			var ns int
			if s == 0 && j == 0 {
				ns = 0
			} else {
				ns = s | (1 << j)
			}
			if bits.OnesCount(uint(ns)) > 2 {
				continue
			}
			ne := e
			if j < x {
				ne = 0
			}
			ndp[ns][ne] += v
		}
	}

	for _, x := range ds {
		// state = 0, or only 1 digits
		// i > 0
		for e := range 2 {
			update(0, e, dp[0][e], x)
			for j1 := range 10 {
				for j2 := range j1 + 1 {
					s := (1 << j1) | (1 << j2)
					update(s, e, dp[s][e], x)
				}
			}
		}

		for j1 := range 10 {
			for j2 := range j1 + 1 {
				for e := range 2 {
					s := (1 << j1) | (1 << j2)
					dp[s][e] = ndp[s][e]
					ndp[s][e] = 0
				}
			}
		}
		dp[0][0] = 1
		dp[0][1] = 0
		ndp[0][0] = 0
		ndp[0][1] = 0
	}

	var ans int
	for e := range 2 {
		for j1 := range 10 {
			for j2 := range j1 {
				ans += dp[1<<j1|1<<j2][e]
			}
			ans += dp[1<<j1][e]
		}
	}

	return ans
}
