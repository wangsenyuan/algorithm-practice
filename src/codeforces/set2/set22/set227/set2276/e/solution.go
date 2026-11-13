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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		for _, v := range drive(reader) {
			fmt.Fprint(writer, v, " ")
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) []int64 {
	var n int
	var s string
	fmt.Fscan(reader, &n, &s)
	return solve(s)
}

func solve(s string) []int64 {
	n := len(s)
	// next[i][d] is the first index >= i where digit d occurs, or n.
	next := make([][10]int, n)
	var at [10]int
	for d := range at {
		at[d] = n
	}
	for i := n - 1; i >= 0; i-- {
		at[s[i]-'0'] = i
		next[i] = at
	}

	// A substring is bad when it has 2025 as a subsequence and not 2026.
	// For a fixed start l those bad substrings end on [end25, end26).
	// diff and slope are the second difference of how many bad substrings cover each index.
	diff := make([]int64, n+1)
	slope := make([]int64, n+1)
	endOf := func(l int, pat string) int {
		r := l
		for j := range len(pat) {
			if r >= n {
				return n
			}
			r = next[r][pat[j]-'0']
			if r >= n {
				return n
			}
			r++
		}
		return r - 1
	}
	for l := range n {
		end25 := endOf(l, "2025")
		end26 := endOf(l, "2026")
		if end25 >= end26 {
			continue
		}
		diff[l] += int64(end26 - end25)
		slope[end25+1]--
		if end26 < n {
			slope[end26+1]++
		}
	}

	ans := make([]int64, n)
	var bad, step int64
	for i := range n {
		bad += diff[i]
		step += slope[i]
		bad += step
		// (i+1)*(n-i) substrings contain index i (0-based).
		ans[i] = int64(i+1)*int64(n-i) - bad
	}
	return ans
}
