package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	b := make([]int, m)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a []int, b []int) int {
	sort.Ints(a)
	slices.Reverse(a)

	sort.Ints(b)

	n := len(a)
	m := len(b)
	dp := make([]int, n)

	for i := range n {
		dp[i] = abs(a[i] - b[i])
		if i > 0 {
			dp[i] += dp[i-1]
		}
	}
	res := dp[n-1]

	var sum int
	for i := n - 1; i >= 0; i-- {
		j := m - (n - i)
		sum += abs(b[j] - a[i])
		if i > 0 {
			res = max(res, sum+dp[i-1])
		} else {
			res = max(res, sum)
		}
	}
	return res
}

func abs(a int) int {
	return max(a, -a)
}
