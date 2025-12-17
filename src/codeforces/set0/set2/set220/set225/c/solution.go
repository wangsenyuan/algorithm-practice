package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, x, y int
	fmt.Fscan(reader, &n, &m, &x, &y)
	a := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, x, y)
}

func solve(a []string, x int, y int) int {
	n := len(a)
	m := len(a[0])

	// dp[j] 表示把这列都变成#的cost
	dp := make([][]int, m)
	for j := range m {
		dp[j] = make([]int, 2)
	}
	for i := range n {
		for j := range m {
			if a[i][j] == '.' {
				dp[j][1]++
			} else {
				dp[j][0]++
			}
		}
	}
	// fp[i][j] 表示到i为止，宽度为j的cost
	fp := make([][]int, y+1)
	nfp := make([][]int, y+1)

	for j := range y + 1 {
		fp[j] = make([]int, 2)
		nfp[j] = make([]int, 2)
		for c := range 2 {
			fp[j][c] = inf
			nfp[j][c] = inf
		}
	}

	fp[1][0] = dp[0][0]
	fp[1][1] = dp[0][1]

	for i := 1; i < m; i++ {
		for c := range 2 {
			for j := range y + 1 {
				if j+1 <= y {
					nfp[j+1][c] = min(nfp[j+1][c], fp[j][c]+dp[i][c])
				}
				if j >= x {
					nfp[1][c^1] = min(nfp[1][c^1], fp[j][c]+dp[i][c^1])
				}
			}
		}
		for c := range 2 {
			for j := range y + 1 {
				fp[j][c] = nfp[j][c]
				nfp[j][c] = inf
			}
		}

	}

	res := inf
	for j := x; j <= y; j++ {
		res = min(res, fp[j][0], fp[j][1])
	}

	return res
}

const inf = 1 << 60
