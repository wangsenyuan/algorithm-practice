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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const inf = 1 << 60

func solve(a []int) int {
	slices.Sort(a)

	dp := make([][]int, 3)
	ndp := make([][]int, 3)

	for i := range 3 {
		dp[i] = make([]int, 3)
		ndp[i] = make([]int, 3)
		for j := range 3 {
			dp[i][j] = -inf
			ndp[i][j] = -inf
		}
	}

	dp[0][0] = 0

	n := len(a)
	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}

		cnt := i - j

		if j > 0 && a[j-1] == a[j]-1 {
			for d0 := range 3 {
				for d1 := range 3 {
					// (?, d0, ?) and (d1, ?, ?)
					if d0+d1 <= cnt {
						for d2 := 0; d2 < 3 && d0+d1+d2 <= cnt; d2++ {
							// d2是作为(x,y,z)的开始的数量
							ndp[d1][d2] = max(ndp[d1][d2], dp[d0][d1]+d0+(cnt-d0-d1-d2)/3)
						}
					}
				}
			}
		}

		var tmp int
		for d0 := range 3 {
			for d1 := range 3 {
				tmp = max(tmp, dp[d0][d1])
			}
		}
		for d1 := 0; d1 < 3 && d1 <= cnt; d1++ {
			ndp[0][d1] = max(ndp[0][d1], tmp+(cnt-d1)/3)
		}

		for d0 := range 3 {
			for d1 := range 3 {
				dp[d0][d1] = ndp[d0][d1]
				ndp[d0][d1] = -inf
			}
		}
	}

	var res int
	for d0 := range 3 {
		for d1 := range 3 {
			res = max(res, dp[d0][d1])
		}
	}
	return res
}
