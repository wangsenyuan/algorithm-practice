package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	people := make([][]int, n)
	for i := 0; i < n; i++ {
		people[i] = make([]int, 2)
		fmt.Fscan(reader, &people[i][0], &people[i][1])
	}
	return solve(people)
}

const inf = 1 << 60

func solve(people [][]int) int {
	// n := len(people)
	var max_height int
	for _, cur := range people {
		max_height = max(max_height, cur[0], cur[1])
	}
	dp := make([]int, max_height+1)
	for i := range max_height + 1 {
		dp[i] = inf
	}
	dp[0] = 0
	ndp := make([]int, max_height+1)
	// 这样子，复杂性，就变成了 1e6 * 1e3, 肯定不行
	// 高度不会超过1e3
	for _, cur := range people {
		w, h := cur[0], cur[1]
		for j := range max_height + 1 {
			ndp[j] = inf
		}
		for j := range max_height + 1 {
			h1 := max(h, j)
			ndp[h1] = min(ndp[h1], dp[j]+w)
			h2 := max(w, j)
			ndp[h2] = min(ndp[h2], dp[j]+h)
		}
		copy(dp, ndp)
	}
	best := inf
	for h := range max_height + 1 {
		if h > 0 && dp[h] <= best/h {
			best = min(best, dp[h]*h)
		}
	}
	return best
}
