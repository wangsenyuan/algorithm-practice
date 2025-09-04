package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var t, k int
	fmt.Fscan(reader, &t, &k)
	queries := make([][]int, t)
	for i := range t {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(k, queries)
}

const mod = 1000000007

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

func pow(a, b int) int {
	var res = 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func solve(k int, queries [][]int) []int {
	B := queries[0][1]
	for _, cur := range queries {
		B = max(B, cur[1])
	}

	dp := make([]int, B+1)
	fp := make([]int, B+1)
	dp[0] = 1
	fp[0] = 1
	for i := 1; i <= B; i++ {
		dp[i] = dp[i-1]
		if i >= k {
			dp[i] = add(dp[i], dp[i-k])
		}
		fp[i] = add(fp[i-1], dp[i])
	}

	res := make([]int, len(queries))
	for i, cur := range queries {
		a, b := cur[0], cur[1]
		res[i] = add(fp[b], mod-fp[a-1])
	}

	return res
}
