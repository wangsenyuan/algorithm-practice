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

const inf = 1 << 60

func solve(a []int) int {
	n := len(a)

	if n == 2 {
		return abs(a[0] - a[1])
	}

	calc := func(arr []int) int {
		// len(arr) = 2 or 3
		return slices.Max(arr) - slices.Min(arr)
	}

	if n == 3 {
		return calc(a)
	}
	// 2个数，或者3个数一组，但是起点可以是n-2, n-1, 0

	dp := make([]int, n)

	play := func(w int) int {
		for i := range n {
			dp[i] = inf
		}
		switch w {
		case 3:
			dp[0] = calc([]int{a[n-2], a[n-1], a[0]})
			dp[1] = abs(a[n-2]-a[n-1]) + abs(a[0]-a[1])
			dp[2] = min(dp[0]+abs(a[1]-a[2]), abs(a[n-2]-a[n-1])+calc(a[:3]))
		case 2:
			dp[0] = abs(a[n-1] - a[0])
			dp[1] = calc([]int{a[n-1], a[0], a[1]})
			dp[2] = dp[0] + abs(a[1]-a[2])
		default:
			// w == 1
			dp[1] = abs(a[0] - a[1])
			dp[2] = calc(a[:3])
		}

		for i := 1; i < n-(w-1); i++ {
			for d := 2; d <= 3 && i+d <= n-(w-1); d++ {
				dp[i+d-1] = min(dp[i+d-1], dp[i-1]+calc(a[i:i+d]))
			}
		}
		return dp[n-w]
	}

	return min(play(1), play(2), play(3))
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
