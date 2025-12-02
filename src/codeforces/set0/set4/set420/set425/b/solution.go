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
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(n, m, k, a)
}

const inf = 1 << 60

func solve(n int, m int, k int, a [][]int) int {

	calc := func(x []int) int {
		// x是行坐标
		var cnt int
		for j := range m {
			var tmp int
			for i := range n {
				// 如果y[j] = 0
				tmp += x[i] ^ a[i][j]
			}
			cnt += min(tmp, n-tmp)
			if cnt > k {
				return inf
			}
		}

		return cnt
	}

	if n <= k {
		x := make([]int, n)
		res := inf
		for state := range 1 << n {
			for i := range n {
				x[i] = (state >> i) & 1
			}
			res = min(res, calc(x))
		}
		if res <= k {
			return res
		}
		return -1
	}

	calc2 := func(y []int) int {
		var cnt int
		for i := range n {
			var tmp int
			for j := range m {
				tmp += y[j] ^ a[i][j]
			}
			cnt += min(tmp, m-tmp)
			if cnt > k {
				return inf
			}
		}
		return cnt
	}
	y := make([]int, m)
	// n > k
	res := inf
	for i := range n {
		// 如果这一行不变
		// a[i][j] ^ x[i] = y[j]
		// 如果x[i] = 0
		for j := range m {
			y[j] = a[i][j]
		}
		res = min(res, calc2(y))
		// 如果x[i] = 1
		for j := range m {
			y[j] = a[i][j] ^ 1
		}
		res = min(res, calc2(y))
	}

	if res <= k {
		return res
	}
	return -1
}
