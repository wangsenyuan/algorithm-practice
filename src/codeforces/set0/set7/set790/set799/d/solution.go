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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var a, b, h, w, n int
	fmt.Fscan(reader, &a, &b, &h, &w, &n)
	extensions := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &extensions[i])
	}
	return solve(a, b, h, w, extensions)
}

func solve(a int, b int, h int, w int, extensions []int) int {
	a, b = min(a, b), max(a, b)
	h, w = min(h, w), max(h, w)
	if a <= h && b <= w {
		return 0
	}
	slices.Sort(extensions)
	slices.Reverse(extensions)

	if len(extensions) > 34 {
		extensions = extensions[:34]
	}

	// n := len(extensions)

	// dp[i][h] 表示前i个exentions达到高度为h时的最大宽度
	dp := make([]int, b+1)
	dp[min(h, a)] = w

	ndp := make([]int, b+1)

	for i, v := range extensions {
		for j := 1; j <= b; j++ {
			u := min(b, j*v)
			ndp[u] = max(ndp[u], dp[j])
			ndp[j] = max(ndp[j], dp[j]*v)
		}

		for j := 1; j <= b; j++ {
			x := ndp[j]
			if a <= min(j, x) && b <= max(j, x) {
				return i + 1
			}
			dp[j] = x
			ndp[j] = 0
		}
	}

	return -1
}
