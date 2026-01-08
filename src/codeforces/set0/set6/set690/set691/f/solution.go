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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	q := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &q[i])
	}
	return solve(a, q)
}

func solve(a []int, q []int) []int {
	slices.Sort(a)

	mx := max(slices.Max(a), slices.Max(q))
	freq := make([]int, mx+2)
	n := len(a)

	var arr []int

	for i := 0; i < n; {
		arr = append(arr, a[i])
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		freq[a[j]] = i - j
	}

	dp := make([]int, mx+2)

	for _, v := range arr {
		for w := v; w <= mx; w += v {
			// dp[w] = freq[v] * freq[v/w]
			if w != v*v {
				dp[w] += freq[v] * freq[w/v]
			} else {
				dp[w] += freq[v] * (freq[v] - 1)
			}
		}
	}

	for i := 1; i <= mx; i++ {
		dp[i] += dp[i-1]
	}

	ans := make([]int, len(q))
	for i, k := range q {
		ans[i] = n*(n-1) - dp[k-1]
	}

	return ans
}
