package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s, x int
	fmt.Fscan(reader, &s, &x)
	res := solve(s, x)
	fmt.Println(res)
}

func solve(s int, x int) int {
	dp := make([]int, 2)
	dp[0] = 1
	ndp := make([]int, 2)

	for i := 0; i < 62; i++ {
		u := (s >> i) & 1
		v := (x >> i) & 1
		clear(ndp)
		for c := 0; c < 2; c++ {
			// a[i] ^ b[i] = v
			// (a[i] + b[i] + c) % 2 = u
			for d1 := range 2 {
				d2 := v ^ d1
				if (d1+d2+c)%2 == u {
					nc := (d1 + d2 + c) / 2
					ndp[nc] += dp[c]
				}
			}
		}
		copy(dp, ndp)
	}

	res := dp[0]

	if s == x {
		// (a, b) = (0, x) or (x, 0)
		res -= 2
	}

	return res
}
