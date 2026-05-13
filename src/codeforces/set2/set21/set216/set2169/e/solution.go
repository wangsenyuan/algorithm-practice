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
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	points := make([][]int, n)
	for i := range n {
		points[i] = make([]int, 3)
	}
	for d := range 3 {
		for i := range n {
			fmt.Fscan(reader, &points[i][d])
		}
	}
	return solve(points)
}

const inf = 1 << 60

func solve(points [][]int) int {

	dp := make([]int, 1<<4)
	ndp := make([]int, 1<<4)
	for mask := range 1 << 4 {
		dp[mask] = -inf
		ndp[mask] = -inf
	}

	dp[0] = 0
	fp := make([]int, 1<<4)

	for _, cur := range points {
		clear(fp)
		for mask := range 1 << 4 {
			for d := range 4 {
				if (mask>>d)&1 == 1 {
					v := cur[d&1]
					if d < 2 {
						v *= -1
					}
					fp[mask] += v
				}
			}
			ndp[mask] = -inf
		}

		for mask := range 1 << 4 {
			if dp[mask] > -inf {
				ndp[mask] = max(ndp[mask], dp[mask]+cur[2])

				rev := ((1 << 4) - 1) ^ mask

				for sub := rev; sub > 0; sub = (sub - 1) & rev {
					ndp[mask|sub] = max(ndp[mask|sub], dp[mask]+2*fp[sub])
				}
			}
		}
		copy(dp, ndp)
	}

	return dp[(1<<4)-1]
}
