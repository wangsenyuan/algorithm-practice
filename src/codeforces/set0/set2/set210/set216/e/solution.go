package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var k, b, n int
	fmt.Fscan(reader, &k, &b, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, b, a)
}

func solve(k int, b int, a []int) int {

	var res0 int

	var cnt0 int
	for _, v := range a {
		if v > 0 {
			res0 += cnt0 * (cnt0 + 1) / 2
			cnt0 = 0
		} else {
			cnt0++
		}
	}
	res0 += cnt0 * (cnt0 + 1) / 2
	if b == 0 {
		return res0
	}
	var res int

	dp := make(map[int]int)
	dp[0] = 1
	k--
	var sum int
	for _, v := range a {
		sum = (sum + v) % k
		res += dp[(sum-b+k)%k]
		dp[sum]++
	}

	if b == k {
		res -= res0
	}

	return res
}
