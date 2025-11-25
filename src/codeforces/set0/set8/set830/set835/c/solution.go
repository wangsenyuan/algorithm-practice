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
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q, c int
	fmt.Fscan(reader, &n, &q, &c)
	stars := make([][]int, n)
	for i := 0; i < n; i++ {
		var x, y, s int
		fmt.Fscan(reader, &x, &y, &s)
		stars[i] = []int{x, y, s}
	}
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		var t, x1, y1, x2, y2 int
		fmt.Fscan(reader, &t, &x1, &y1, &x2, &y2)
		queries[i] = []int{t, x1, y1, x2, y2}
	}
	return solve(c, stars, queries)
}

const H = 100
const W = 100

func solve(c int, stars [][]int, queries [][]int) []int {
	dp := make([][H + 1][W + 1]int, c+1)

	for _, star := range stars {
		x, y, s := star[0], star[1], star[2]
		for i := range c + 1 {
			dp[i][x][y] += int(s+i) % (c + 1)
		}
	}

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			for d := range c + 1 {
				dp[d][i][j] += dp[d][i][j-1] + dp[d][i-1][j] - dp[d][i-1][j-1]
			}
		}
	}

	get := func(t int, x1 int, y1 int, x2 int, y2 int) int {
		return dp[t][x2][y2] - dp[t][x1-1][y2] - dp[t][x2][y1-1] + dp[t][x1-1][y1-1]
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		t, x1, y1, x2, y2 := cur[0], cur[1], cur[2], cur[3], cur[4]
		t %= (c + 1)
		ans[i] = get(t, x1, y1, x2, y2)
	}
	return ans
}
