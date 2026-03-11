package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)
	fmt.Println(solve(n, m))
}

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

func mul(a, b int) int {
	return a * b % mod
}

func buildPrefix(a [][]int) [][]int {
	m := len(a)
	p := make([][]int, m+1)
	for i := range m + 1 {
		p[i] = make([]int, m+1)
	}
	for i := range m {
		row := 0
		for j := range m {
			row = add(row, a[i][j])
			p[i+1][j+1] = add(p[i][j+1], row)
		}
	}
	return p
}

func rect(p [][]int, x1, x2, y1, y2 int) int {
	if x1 > x2 || y1 > y2 {
		return 0
	}
	res := p[x2+1][y2+1]
	res = sub(res, p[x1][y2+1])
	res = sub(res, p[x2+1][y1])
	res = add(res, p[x1][y1])
	return res
}

func sumAll(a [][]int) int {
	var res int
	for i := range a {
		for j := range a[i] {
			res = add(res, a[i][j])
		}
	}
	return res
}

func solve(n int, m int) int {
	if n < m {
		n, m = m, n
	}

	makeGrid := func() [][]int {
		res := make([][]int, m)
		for i := range m {
			res[i] = make([]int, m)
		}
		return res
	}

	// A painted figure is exactly a sequence of contiguous row intervals.
	// For a valid figure, left endpoints can only change direction once
	// (non-increasing, then non-decreasing), and right endpoints similarly
	// (non-decreasing, then non-increasing).
	//
	// State bits:
	// 00: left not turned,  right not turned
	// 01: left not turned,  right turned
	// 10: left turned,      right not turned
	// 11: left turned,      right turned
	s00 := makeGrid()
	s01 := makeGrid()
	s10 := makeGrid()
	s11 := makeGrid()

	for l := 0; l < m; l++ {
		for r := l; r < m; r++ {
			s00[l][r] = 1
		}
	}

	oneRow := sumAll(s00)
	ans := mul(oneRow, n)

	for h := 2; h <= n; h++ {
		p00 := buildPrefix(s00)
		p01 := buildPrefix(s01)
		p10 := buildPrefix(s10)
		p11 := buildPrefix(s11)

		t00 := makeGrid()
		t01 := makeGrid()
		t10 := makeGrid()
		t11 := makeGrid()

		for l := 0; l < m; l++ {
			for r := l; r < m; r++ {
				// l <= l1, r1 <= r
				t00[l][r] = rect(p00, l, m-1, 0, r)
				// l <= l1, r <= r1 (r1 - 1)
				t01[l][r] = add(
					rect(p00, l, r, r+1, m-1),
					rect(p01, l, r, r, m-1),
				)
				// l1 <= l, r1 <= r
				t10[l][r] = add(
					rect(p00, 0, l-1, l, r),
					rect(p10, 0, l, l, r),
				)
				t11[l][r] = rect(p00, 0, l-1, r+1, m-1)
				t11[l][r] = add(t11[l][r], rect(p01, 0, l-1, r, m-1))
				t11[l][r] = add(t11[l][r], rect(p10, 0, l, r+1, m-1))
				t11[l][r] = add(t11[l][r], rect(p11, 0, l, r, m-1))
			}
		}

		s00, s01, s10, s11 = t00, t01, t10, t11

		total := 0
		total = add(total, sumAll(s00))
		total = add(total, sumAll(s01))
		total = add(total, sumAll(s10))
		total = add(total, sumAll(s11))

		// 有placements个起始位置，放置这个图形
		placements := n - h + 1
		ans = add(ans, mul(total, placements))
	}

	return ans
}
