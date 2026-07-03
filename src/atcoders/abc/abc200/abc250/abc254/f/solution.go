package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for _, v := range drive(reader) {
		fmt.Println(v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 4)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2], &queries[i][3])
	}
	return solve(a, b, queries)
}

func solve(a, b []int, queries [][]int) []int {
	n := len(a)
	h := bits.Len(uint(n))
	dp := make([][]int, h)
	fp := make([][]int, h)

	for i := range h {
		dp[i] = make([]int, n+1)
		fp[i] = make([]int, n+1)
	}

	for i := range n - 1 {
		dp[0][i+1] = abs(a[i+1] - a[i])
		fp[0][i+1] = abs(b[i+1] - b[i])
	}

	for d := 1; d < h; d++ {
		for i := 1; i < n; i++ {
			dp[d][i] = dp[d-1][i]
			fp[d][i] = fp[d-1][i]
			if i+(1<<(d-1)) < n {
				dp[d][i] = gcd(dp[d][i], dp[d-1][i+(1<<(d-1))])
				fp[d][i] = gcd(fp[d][i], fp[d-1][i+(1<<(d-1))])
			}
		}
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		r1, r2, c1, c2 := cur[0]-1, cur[1]-1, cur[2]-1, cur[3]-1
		var x, y int
		if r1 < r2 {
			d1 := bits.Len(uint(r2-r1)) - 1
			x1 := dp[d1][r1+1]
			x2 := dp[d1][r2+1-(1<<d1)]
			x = gcd(x1, x2)
		}
		if c1 < c2 {
			d2 := bits.Len(uint(c2-c1)) - 1
			y1 := fp[d2][c1+1]
			y2 := fp[d2][c2+1-(1<<d2)]
			y = gcd(y1, y2)
		}

		ans[i] = gcd(a[r1]+b[c1], gcd(x, y))
	}

	return ans
}

func solve1(a, b []int, queries [][]int) []int {
	t1 := make(SegTree, 2*len(a))

	for i := 1; i < len(a); i++ {
		t1.Update(i, abs(a[i]-a[i-1]))
	}

	t2 := make(SegTree, 2*len(b))
	for i := 1; i < len(b); i++ {
		t2.Update(i, abs(b[i]-b[i-1]))
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		r1, r2, c1, c2 := cur[0]-1, cur[1]-1, cur[2]-1, cur[3]-1
		x := t1.Query(r1+1, r2+1)
		y := t2.Query(c1+1, c2+1)
		ans[i] = gcd(a[r1]+b[c1], gcd(x, y))
	}

	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func abs(num int) int {
	return max(num, -num)
}

type SegTree []int

func (s SegTree) Update(p int, v int) {
	n := len(s) / 2
	p += n
	s[p] = v
	for p > 1 {
		s[p>>1] = gcd(s[p], s[p^1])
		p >>= 1
	}
}

func (s SegTree) Query(l, r int) int {
	if l > r {
		return 0
	}
	n := len(s) / 2
	l += n
	r += n
	ans := 0
	for l < r {
		if l&1 == 1 {
			ans = gcd(ans, s[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = gcd(ans, s[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}
