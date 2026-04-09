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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var w, f, n int
	fmt.Fscan(reader, &w, &f, &n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	return solve(w, f, s)
}

func solve(w int, f int, s []int) int {
	// n := len(s)
	var sum int
	for _, x := range s {
		sum += x
	}

	dp := make([]bool, sum+1)
	dp[0] = true
	sum = 0
	for _, v := range s {
		sum += v
		for w := sum; w >= v; w-- {
			if dp[w-v] {
				dp[w] = true
			}
		}
	}

	best := (sum + w - 1) / w

	for v := range sum + 1 {
		if dp[v] {
			best = min(best, max((v+w-1)/w, (sum-v+f-1)/f))
		}
	}

	return best
}
