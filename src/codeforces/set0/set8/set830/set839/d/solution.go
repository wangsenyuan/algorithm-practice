package main

import (
	"bufio"
	"fmt"
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

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func inv(a int) int {
	return pow(a, mod-2)
}

func solve(a []int) int {
	x := slices.Max(a)
	freq := make([]int, x+1)
	for _, num := range a {
		freq[num]++
	}
	// sum i * C(n, i) = n * pow(2, n - 1)
	for i := 1; i <= x; i++ {
		for j := i * 2; j <= x; j += i {
			freq[i] += freq[j]
		}
	}

	dp := make([]int, x+1)
	var ans int
	for i := x; i > 1; i-- {
		if freq[i] == 0 {
			continue
		}
		dp[i] = mul(freq[i], pow(2, freq[i]-1))
		for j := 2 * i; j <= x; j += i {
			dp[i] = sub(dp[i], dp[j])
		}
		// dp[i] = mul(dp[i], i)
		ans = add(ans, mul(dp[i], i))
	}

	return ans
}
