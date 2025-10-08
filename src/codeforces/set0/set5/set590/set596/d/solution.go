package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, h int
	var p float64
	fmt.Fscan(reader, &n, &h, &p)
	x := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	return solve(n, h, p, x)
}

func solve(n int, h int, p float64, x []int) float64 {
	x = append(x, 1<<30)
	x = append(x, -(1 << 30))
	sort.Ints(x)
	n += 2
	next := make([]int, n+1)
	next[n] = n
	for i := n - 1; i >= 0; i-- {
		next[i] = i
		if i+1 < n && x[i+1]-x[i] < h {
			next[i] = next[i+1]
		}
	}
	prev := make([]int, n)
	for i := range n {
		prev[i] = i
		if i > 0 && x[i]-x[i-1] < h {
			prev[i] = prev[i-1]
		}
	}

	dp := make([][][2][2]float64, n)

	for i := range n {
		dp[i] = make([][2][2]float64, n)
		for j := range n {
			for u := range 2 {
				for v := range 2 {
					dp[i][j][u][v] = -1
				}
			}
		}
	}

	var f func(l int, r int, b1 int, b2 int) float64

	f = func(l int, r int, b1 int, b2 int) (res float64) {
		if l > r {
			return 0
		}
		if dp[l][r][b1][b2] != -1 {
			return dp[l][r][b1][b2]
		}
		defer func() {
			dp[l][r][b1][b2] = res
		}()
		j := min(next[l], r)
		if j == r {
			if b2 == 0 {
				res += (1 - p) * float64(min(h, x[r+1]-x[r]-h)+x[r]-x[l])
			} else {
				res += (1 - p) * float64(min(h, x[r+1]-x[r])+x[r]-x[l])
			}
		} else {
			res += (1 - p) * f(j+1, r, 1, b2)
			res += (1 - p) * float64(x[j]+h-x[l])
		}

		if b1 == 1 {
			res += p * float64(min(h, x[l]-x[l-1]-h))
		} else {
			// b1 == 0
			res += p * float64(min(h, x[l]-x[l-1]))
		}

		res += p * f(l+1, r, 0, b2)

		i := max(prev[r], l)

		if i == l {
			if b1 == 1 {
				res += p * float64(min(h, x[l]-x[l-1]-h)+x[r]-x[l])
			} else {
				res += p * float64(min(h, x[l]-x[l-1])+x[r]-x[l])
			}
		} else {
			res += p * f(l, i-1, b1, 0)
			res += p * float64(x[r]-x[i]+h)
		}

		if b2 == 0 {
			res += (1 - p) * float64(min(h, x[r+1]-x[r]-h))
		} else {
			res += (1 - p) * float64(min(h, x[r+1]-x[r]))
		}

		res += (1 - p) * f(l, r-1, b1, 1)

		res /= 2

		return
	}

	return f(1, n-2, 0, 1)
}
