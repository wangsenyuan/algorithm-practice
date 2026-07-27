package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	ans := make([]int, t)
	for i := range t {
		var n, c int
		var s string
		fmt.Fscan(reader, &n, &c, &s)
		ans[i] = solve(n, c, s)
	}
	return ans
}

const mod = 1e9 + 7

func mul(a, b int) int {
	return a * b % mod
}
func solve(n, c int, s string) int {
	if s[0] != '1' || s[n-1] != '1' {
		return -1
	}

	ans := 1
	for i := 1; i < n; i++ {
		if s[i] == '1' {
			ans = mul(ans, 2)
			c /= gcd(c, 2)
		} else {
			ans = mul(ans, i)
			c /= gcd(c, i)
		}
	}

	if c == 1 {
		return -1
	}

	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
