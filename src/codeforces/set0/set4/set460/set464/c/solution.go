package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	s := readString(reader)
	n1 := readString(reader)
	n, _ := strconv.Atoi(n1)
	queries := make([]string, n)
	for i := range n {
		queries[i] = readString(reader)
	}
	return solve(s, queries)
}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(nums ...int) int {

	res := 1
	for _, num := range nums {
		res = res * num % mod
	}

	return res
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

func inverse(a int) int {
	return pow(a, mod-2)
}

const X = 1000010

func solve(s string, queries []string) int {
	// dp[x][y] 表示当更新x的时候，对y的影响
	dp := make([][]int, 10)
	for i := range 10 {
		dp[i] = []int{i, 10}
	}
	ndp := make([][]int, 10)
	for i := range 10 {
		ndp[i] = make([]int, 2)
	}

	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		d := int(q[0] - '0')
		q = q[3:]
		for j := range 10 {
			copy(ndp[j], dp[j])
		}
		val := 0
		base := 1
		for j := len(q) - 1; j >= 0; j-- {
			c := int(q[j] - '0')
			val = add(val, mul(dp[c][0], base))
			base = mul(base, dp[c][1])
		}
		ndp[d][0] = val
		ndp[d][1] = base
		for j := range 10 {
			copy(dp[j], ndp[j])
		}
	}

	var ans int
	base := 1

	for i := len(s) - 1; i >= 0; i-- {
		c := int(s[i] - '0')
		ans = add(ans, mul(dp[c][0], base))
		base = mul(base, dp[c][1])
	}

	return ans
}
