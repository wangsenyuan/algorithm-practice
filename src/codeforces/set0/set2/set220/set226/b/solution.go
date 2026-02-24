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
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([]int, q)
	for i := range q {
		fmt.Fscan(reader, &queries[i])
	}
	return solve(a, queries)
}

func solve(a []int, queries []int) []int {
	n := len(a)

	slices.Sort(a)
	slices.Reverse(a)

	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}

	res := make([]int, n)

	for i := range n {
		res[i] = -1
	}

	play := func(k int) int {
		if k >= n {
			k = n - 1
		}
		if res[k] >= 0 {
			return res[k]
		}
		var tmp int
		sz := 1
		var d int
		for i := 0; i < n; {
			j := min(n, i+sz)
			tmp += (sum[j] - sum[i]) * d
			d++
			i = j
			sz *= k
		}
		res[k] = tmp
		return res[k]
	}

	ans := make([]int, len(queries))

	for i, k := range queries {
		ans[i] = play(k)
	}

	return ans
}
