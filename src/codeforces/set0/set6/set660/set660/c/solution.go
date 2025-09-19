package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, best, res := drive(reader)
	fmt.Println(best)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (k int, a []int, best int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &k)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	best, res = solve(a, k)
	return
}

func solve(a []int, k int) (max_length int, res []int) {
	n := len(a)
	sum := make([]int, n+1)
	var j int
	var best int
	dp := make([]int, n)
	for i, v := range a {
		sum[i+1] = sum[i] + v
		// i - j + 1 是这个区间的总数
		// sum[i+1] - sum[j] 是这个区间1的个数
		// i - j + 1 - ()
		for j < n && (i-j+1)-(sum[i+1]-sum[j]) > k {
			j++
		}
		dp[i] = j
		if i-j+1 > best-dp[best]+1 {
			best = i
		}
	}

	max_length = best - dp[best] + 1

	res = slices.Clone(a)

	r := best
	for k > 0 && r >= 0 {
		if res[r] == 0 {
			res[r] = 1
			k--
		}
		r--
	}

	return
}
