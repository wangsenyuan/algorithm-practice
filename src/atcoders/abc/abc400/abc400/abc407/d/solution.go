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
	var h, w int
	fmt.Fscan(reader, &h, &w)
	a := make([][]int, h)
	for i := range h {
		a[i] = make([]int, w)
		for j := range w {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

func solve(a [][]int) int {
	if len(a) < len(a[0]) {
		a = transpose(a)
	}
	// n >= m => m <= 4
	n, m := len(a), len(a[0])
	// n * m <= 20
	M := 1 << m
	dp := make([]map[int]int, M)
	ndp := make([]map[int]int, M)
	for state := range M {
		// 计算被cover备份的 xor 值
		dp[state] = make(map[int]int)
	}
	dp[M-1][0] = 1

	var f func(s1 int, s2 int, r int, c int, v int)

	f = func(s1 int, s2 int, r int, c int, v int) {
		if c == m {
			if len(ndp[s2]) == 0 {
				ndp[s2] = make(map[int]int)
			}
			ndp[s2][v]++
			return
		}

		// 不覆盖
		f(s1, s2, r, c+1, v)

		if (s1>>c)&1 == 0 {
			// 和上一行匹配
			f(s1, s2|(1<<c), r, c+1, v^a[r-1][c]^a[r][c])
		}
		if c+1 < m {
			// 连续两列
			f(s1, s2|(1<<c)|(1<<(c+1)), r, c+2, v^a[r][c]^a[r][c+1])
		}
	}
	for r := range n {
		for s, w := range dp {
			for v := range w {
				f(s, 0, r, 0, v)
			}
		}
		copy(dp, ndp)
		clear(ndp)
	}

	var sum int
	for _, row := range a {
		for _, v := range row {
			sum ^= v
		}
	}

	var best int
	for _, w := range dp {
		for v := range w {
			best = max(best, v^sum)
		}
	}

	return best
}

func transpose(a [][]int) [][]int {
	m := len(a)
	n := len(a[0])
	b := make([][]int, n)
	for i := range n {
		b[i] = make([]int, m)
		for j := range m {
			b[i][j] = a[j][i]
		}
	}
	return b
}
