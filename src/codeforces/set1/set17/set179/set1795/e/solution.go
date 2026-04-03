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
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	dp := play(a)
	slices.Reverse(a)
	fp := play(a)
	slices.Reverse(fp)
	slices.Reverse(a)
	var best int
	var sum int
	for i, v := range a {
		sum += v
		best = max(best, dp[i]+fp[i]-2*a[i])
	}
	return sum - best
}

func play(a []int) []int {
	n := len(a)
	stack := make([]int, n)
	var top int

	get := func(i int) int {
		return a[i] - i
	}

	dp := make([]int, n)

	for i := range n {
		v := a[i] - i
		for top > 0 && get(stack[top-1]) > v {
			top--
		}
		j := max(-1, i-a[i])
		if top > 0 {
			j = stack[top-1]
			dp[i] = dp[j]
		}

		dp[i] += (a[i] - (i - j - 1) + a[i]) * (i - j) / 2
		stack[top] = i
		top++
	}
	return dp
}
