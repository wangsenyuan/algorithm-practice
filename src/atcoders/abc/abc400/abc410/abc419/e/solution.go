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
	var n, m, l int
	fmt.Fscan(reader, &n, &m, &l)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(m, l, a)
}

func solve(m int, l int, a []int) int {
	n := len(a)

	// dp[i][j] = 让位置 pos % l = i, 且使的a[pos] % m = j 的最少操作数
	dp := make([][]int, l)
	for i := range l {
		dp[i] = make([]int, m)
	}

	for pos := range n {
		i := pos % l
		for j := range m {
			// a[pos] % m = j 的操作数
			w := a[pos] % m
			if w == 0 {
				w = m
			}
			if j >= w {
				dp[i][j] += j - w
			} else {
				// j < w
				dp[i][j] += m - w + j
			}
		}
	}
	fp := make([]int, m)
	nfp := make([]int, m)
	for i := range m {
		fp[i] = 1 << 60
		nfp[i] = 1 << 60
	}
	fp[0] = 0
	for i := range l {
		for j := range m {
			for r := range m {
				nfp[(j+r)%m] = min(nfp[(j+r)%m], dp[i][j]+fp[r])
			}
		}
		copy(fp, nfp)
		for j := range m {
			nfp[j] = 1 << 60
		}
	}

	return fp[0]
}
