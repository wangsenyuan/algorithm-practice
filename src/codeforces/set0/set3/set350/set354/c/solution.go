package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

func solve(a []int, k int) int {
	m := slices.Min(a)
	if m <= k+1 {
		return m
	}

	mx := slices.Max(a)
	freq := make([]int, mx+1)
	for _, v := range a {
		freq[v]++
	}
	for i := 1; i <= mx; i++ {
		freq[i] += freq[i-1]
	}

	n := len(a)

	check := func(d int) bool {
		var cnt int
		for i := 1; i*d <= mx; i++ {
			r := min(mx, i*d+k)
			cnt += freq[r] - freq[i*d-1]
		}
		return cnt == n
	}

	// m > k + 1
	ans := k

	for d := k + 1; d <= m; d++ {
		if check(d) {
			ans = d
		}
	}

	return ans
}
