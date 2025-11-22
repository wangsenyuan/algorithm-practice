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
	var t int
	fmt.Fscan(reader, &t)
	for range t {
		var L, R int
		fmt.Fscan(reader, &L, &R)
		fmt.Fprintln(writer, solve(L, R))
	}

}

func solve(L int, R int) int {
	var dr []int
	for i := R; i > 0; i >>= 1 {
		dr = append(dr, i&1)
	}
	var dl []int
	for i := L; i > 0; i >>= 1 {
		dl = append(dl, i&1)
	}
	for len(dl) < len(dr) {
		dl = append(dl, 0)
	}
	slices.Reverse(dl)
	slices.Reverse(dr)

	n := len(dr)
	dp := make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, 4)
		for j := range 4 {
			dp[i][j] = make([]int, 4)
			for k := range 4 {
				dp[i][j][k] = -1
			}
		}
	}

	var f func(i int, d1 int, d2 int) int
	f = func(i int, d1 int, d2 int) (ans int) {
		if i == n {
			return 1
		}
		if dp[i][d1][d2] >= 0 {
			return dp[i][d1][d2]
		}
		defer func() {
			dp[i][d1][d2] = ans
		}()

		for v1 := range 2 {
			if d1&2 == 0 && v1 > dr[i] || d1&1 == 0 && v1 < dl[i] {
				continue
			}
			for v2 := range 2 {
				if v1+v2 == 2 {
					continue
				}
				if d2&2 == 0 && v2 > dr[i] || d2&1 == 0 && v2 < dl[i] {
					continue
				}
				nd1 := d1
				if v1 > dl[i] {
					nd1 |= 1
				}
				if v1 < dr[i] {
					nd1 |= 2
				}
				nd2 := d2
				if v2 > dl[i] {
					nd2 |= 1
				}
				if v2 < dr[i] {
					nd2 |= 2
				}
				ans += f(i+1, nd1, nd2)
			}
		}

		return
	}

	return f(0, 0, 0)
}
