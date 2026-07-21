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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	grid := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &grid[i])
	}
	return solve(grid)
}

const mod = 1000003

func solve(grid []string) int {
	// n := len(grid)
	m := len(grid[0])
	col := make([][2]int, m)

	res := 1

	for i := 0; i < m; i++ {
		col[i][0] = 1
		col[i][1] = 1
	}

	for i, row := range grid {
		var dp [2]int
		// dp[0] 表示左端开口, dp[1]表示右端开口
		dp[0] = 1
		dp[1] = 1
		for j := range m {
			if row[j] == '1' {
				dp[(j&1)^1] = 0
				col[j][(i&1)^1] = 0
			} else if row[j] == '2' {
				dp[(j&1)^1] = 0
				col[j][i&1] = 0
			} else if row[j] == '3' {
				dp[j&1] = 0
				col[j][i&1] = 0
			} else if row[j] == '4' {
				dp[j&1] = 0
				col[j][(i&1)^1] = 0
			}
		}
		if dp[0] == 0 && dp[1] == 0 {
			return 0
		}
		if dp[0] == 1 && dp[1] == 1 {
			res = (res * 2) % mod
		}
	}

	for i := range m {
		if col[i][0]+col[i][1] == 0 {
			return 0
		}
		if col[i][0]+col[i][1] == 2 {
			res = (res * 2) % mod
		}
	}

	return res
}
