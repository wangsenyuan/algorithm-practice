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
	var n, d int
	fmt.Fscan(reader, &n, &d)
	p := make([]int, n)
	q := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	for i := range n {
		fmt.Fscan(reader, &q[i])
	}
	return solve(n, d, p, q)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func solve(n int, d int, p []int, q []int) int {
	dp := make([][]int, d+1)
	for i := range d + 1 {
		dp[i] = make([]int, d+1)
	}
	dp[0][0] = 1

	for t := range n {
		s := abs(p[t] - q[t])

		nxt := make([][]int, d+1)
		dp2 := make([][]int, d+1)
		dp3 := make([][]int, d+1)
		for i := range d + 1 {
			nxt[i] = make([]int, d+1)
			dp2[i] = make([]int, d+1)
			dp3[i] = make([]int, d+1)
		}

		for i := range d + 1 {
			for j := range d + 1 {
				dp2[i][j] = dp[i][j]
				if i != 0 && j != d {
					dp2[i][j] = add(dp2[i][j], dp2[i-1][j+1])
				}
			}
		}

		for i := range d + 1 {
			for j := range d + 1 {
				si := i
				sj := j - s
				if sj < 0 {
					si += sj
					sj = 0
				}
				if 0 <= si && si <= d && 0 <= sj && sj <= d {
					nxt[i][j] = add(nxt[i][j], dp2[si][sj])
				}
				ti := i - (s + 1)
				tj := j + 1
				if 0 <= ti && ti <= d && 0 <= tj && tj <= d {
					nxt[i][j] = sub(nxt[i][j], dp2[ti][tj])
				}
			}
		}

		for i := range d + 1 {
			for j := range d + 1 {
				dp3[i][j] = dp[i][j]
				if i != 0 && j != 0 {
					dp3[i][j] = add(dp3[i][j], dp3[i-1][j-1])
				}
				if i+1 <= d && j+s+1 <= d {
					nxt[i+1][j+s+1] = add(nxt[i+1][j+s+1], dp3[i][j])
				}
				if i+s+1 <= d && j+1 <= d {
					nxt[i+s+1][j+1] = add(nxt[i+s+1][j+1], dp3[i][j])
				}
			}
		}
		dp = nxt
	}

	var res int
	for _, cur := range dp {
		for _, v := range cur {
			res = add(res, v)
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
