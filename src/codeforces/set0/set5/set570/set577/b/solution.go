package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, m)
}

func solve(a []int, m int) bool {
	// m <= 1000
	// 循环的～
	n := len(a)
	if n > m {
		return true
	}

	dp := make([]bool, m)
	ndp := make([]bool, m)
	for _, v := range a {
		v %= m
		if v == 0 {
			return true
		}

		copy(ndp, dp)

		for w := range m {
			if ndp[w] {
				dp[(w+v)%m] = true
			}
		}

		dp[v] = true

		if dp[0] {
			return true
		}
	}

	return false
}
