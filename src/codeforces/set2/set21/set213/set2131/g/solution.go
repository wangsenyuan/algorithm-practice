package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	return solve(s, k)
}

const mod = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

var dp [31]int

func init() {
	for i := range 31 {
		dp[i] = i + 1
		for j := range i {
			dp[i] = mul(dp[i], dp[j])
		}
	}
}

func solve(s []int, k int) int {
	slices.Sort(s)

	for i := range s {
		s[i]--
	}

	score := 1

	var f func(v int, k int) int

	f = func(v int, k int) int {
		if k == 0 {
			return 1
		}
		ans := v + 1
		v = min(v-1, 30)
		k--
		for j := range v + 1 {
			if k >= 1<<j {
				ans = mul(ans, dp[j])
				k -= 1 << j
			} else {
				ans = mul(ans, f(j, k))
				break
			}
		}
		return ans
	}

	for _, v := range s {
		if v <= 30 && k >= 1<<v {
			k -= 1 << v
			score = mul(score, dp[v])
		} else {
			// 1 << v > k
			score = mul(score, f(v, k))
			break
		}
	}

	return score
}
