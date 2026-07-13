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

	for _, f := range drive(reader) {
		fmt.Fprintln(writer, len(f))
		if len(f) > 0 {
			for i, v := range f {
				if i > 0 {
					fmt.Fprint(writer, " ")
				}
				fmt.Fprint(writer, v)
			}
			fmt.Fprintln(writer)
		}
	}
}

func drive(reader *bufio.Reader) [][]int {
	var tc int
	fmt.Fscan(reader, &tc)
	res := make([][]int, tc)
	for i := range tc {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		a := make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[j])
		}
		b := make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &b[j])
		}
		res[i] = solve(a, b)
	}
	return res
}

func solve(a, b []int) []int {
	dp1 := play(a)
	dp2 := play(b)

	n := len(a)
	m := len(b)

	get := func(x int, y int) int {
		return dp1[x] + dp2[y]
	}

	find := func(k int) int {
		// 假设分别是 x, y
		// 2 * x + y <= n
		// x + 2 * y <= m
		// x + y = k
		// 2 * x + k - x <= n => x <= n - k
		// x + 2 * (k - x) <= m => x >= 2 * k - m
		l := max(0, 2*k-m)
		r := min(k, n-k)
		for r-l+1 > 3 {
			m1 := l + (r-l+1)/3
			m2 := r - (r-l+1)/3
			w1 := get(m1, k-m1)
			w2 := get(m2, k-m2)
			if w1 <= w2 {
				l = m1
			} else {
				r = m2
			}
		}
		var res int
		for i := l; i <= r; i++ {
			res = max(res, get(i, k-i))
		}
		return res
	}

	var ans []int

	kmax := min(n, m, (len(a)+len(b))/3)

	for k := 1; k <= kmax; k++ {
		ans = append(ans, find(k))
	}

	return ans
}

func play(a []int) []int {
	slices.Sort(a)
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	dp := make([]int, n+1)

	for i := 1; i*2 <= n; i++ {
		left := sum[i]
		right := sum[n] - sum[n-i]
		dp[i] = right - left
	}

	return dp[:n/2+1]
}
