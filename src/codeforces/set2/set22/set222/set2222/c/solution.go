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
		fmt.Fprintln(writer, drive(reader))
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

const inf = 1 << 60

func solve(a []int) int {
	n := len(a)
	arr := slices.Clone(a)
	slices.Sort(arr)
	median := arr[n/2]

	dp := make([]int, n)

	for i := range n {
		dp[i] = -inf
		var cnt [2]int
		for j := i; j >= 0; j-- {
			if a[j] <= median {
				cnt[0]++
			}
			if a[j] >= median {
				cnt[1]++
			}
			if (i-j+1)%2 == 1 && cnt[0] > (i-j+1)/2 && cnt[1] > (i-j+1)/2 {
				if j > 0 {
					dp[i] = max(dp[i], dp[j-1]+1)
				} else {
					dp[i] = max(dp[i], 1)
				}
			}
		}
	}

	return dp[n-1]
}
