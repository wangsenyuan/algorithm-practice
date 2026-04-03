package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	if res == nil {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	for i, x := range res {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(x)
	}
	fmt.Println()
}

func drive(reader *bufio.Reader) (arr []int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	res = solve(arr)
	return
}

func solve(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}

	const incColor = 0
	const decColor = 1

	// dp[i][0]: after processing [0..i], if arr[i] is put into increasing,
	// the largest possible last value of the decreasing subsequence.
	// dp[i][1]: after processing [0..i], if arr[i] is put into decreasing,
	// the smallest possible last value of the increasing subsequence.
	dp := make([][2]int, n)
	from := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i][0] = -1
		dp[i][1] = inf
		from[i][0] = -1
		from[i][1] = -1
	}

	dp[0][incColor] = inf
	dp[0][decColor] = -1

	for i := 0; i+1 < n; i++ {
		x, y := arr[i], arr[i+1]

		if dp[i][incColor] >= 0 {
			decLast := dp[i][incColor]

			if y > x && decLast > dp[i+1][incColor] {
				dp[i+1][incColor] = decLast
				from[i+1][incColor] = incColor
			}
			if y < decLast && x < dp[i+1][decColor] {
				dp[i+1][decColor] = x
				from[i+1][decColor] = incColor
			}
		}

		if dp[i][decColor] < inf {
			incLast := dp[i][decColor]

			if y < x && incLast < dp[i+1][decColor] {
				dp[i+1][decColor] = incLast
				from[i+1][decColor] = decColor
			}
			if y > incLast && x > dp[i+1][incColor] {
				dp[i+1][incColor] = x
				from[i+1][incColor] = decColor
			}
		}
	}

	lastColor := -1
	if dp[n-1][incColor] >= 0 {
		lastColor = incColor
	} else if dp[n-1][decColor] < inf {
		lastColor = decColor
	} else {
		return nil
	}

	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		res[i] = lastColor
		if i > 0 {
			lastColor = from[i][lastColor]
		}
	}

	return res
}

const inf = 1 << 60
