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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)

	f := func(num int) int {
		// num的最大的odd factor
		for num%2 == 0 {
			num /= 2
		}
		return num
	}

	w := make([]int, n)
	v := make([]int, n)
	for i := range n {
		w[i] = f(a[i])
		v[i] = bits.TrailingZeros(uint(a[i]))
	}

	check := func(i int, j int) bool {
		if w[i]%w[j] != 0 {
			return false
		}
		l := v[i]
		r := v[j]
		return l+j-i == r || r <= j-i-1
	}

	dp := make([]int, n)
	dp[0] = 1
	for j := 1; j < n; j++ {
		dp[j] = 1
		for i := j - 1; i >= 0; i-- {
			if check(i, j) {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}
	return n - slices.Max(dp)
}
