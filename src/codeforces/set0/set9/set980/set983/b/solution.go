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
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		queries[i] = []int{l, r}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	// n <= 5000
	dp := make([][]int, n)
	f := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		f[i] = make([]int, n)
	}

	for r := range n {
		for l := r; l >= 0; l-- {
			if l == r {
				dp[l][r] = a[l]
				f[l][r] = a[l]
			} else {
				f[l][r] = f[l][r-1] ^ f[l+1][r]
				dp[l][r] = max(dp[l][r-1], dp[l+1][r], f[l][r])
			}
		}
	}
	ans := make([]int, len(queries))
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		ans[i] = dp[l-1][r-1]
	}
	return ans
}
