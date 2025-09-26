package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, v := range res {
		buf.WriteString(fmt.Sprintf("%d\n", v))
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, 1<<n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	queries := make([]int, m)
	for i := range queries {
		fmt.Fscan(reader, &queries[i])
	}
	return solve(n, a, queries)
}

func solve(n int, a []int, queries []int) []int {
	dp := make([][2]int, n+1)
	// dp[i][0] 表示这层，和一开始一致时的inversion数
	// dp[i][1] 表示这层，swap后的inversions的数量

	check := func(l []int, r []int) int {
		var res int
		for i, j := 0, 0; i < len(l) || j < len(r); {
			if j == len(r) || i < len(l) && l[i] <= r[j] {
				i++
			} else {
				// l[i] > r[j]
				res += len(l) - i
				j++
			}
		}
		return res
	}

	buf := make([]int, 1<<n)

	merge := func(l int, mid int, r int) {
		p := l
		for i, j := l, mid; i < mid || j < r; {
			if j == r || i < mid && a[i] <= a[j] {
				buf[p] = a[i]
				i++
			} else {
				buf[p] = a[j]
				j++
			}
			p++
		}
		copy(a[l:r], buf[l:r])
	}

	play := func(d int, l int, mid int, r int) {
		dp[d][0] += check(a[l:mid], a[mid:r])
		dp[d][1] += check(a[mid:r], a[l:mid])
		merge(l, mid, r)
	}

	var dfs func(d int, l int)

	dfs = func(d int, l int) {
		r := l + 1<<d
		if l+1 == r {
			return
		}
		mid := (l + r) / 2
		dfs(d-1, l)
		dfs(d-1, mid)
		play(d, l, mid, r)
	}

	dfs(n, 0)
	var sum int
	for i := range n + 1 {
		sum += dp[i][0]
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		for j := q; j > 0; j-- {
			sum -= dp[j][0]
			dp[j][0], dp[j][1] = dp[j][1], dp[j][0]
			sum += dp[j][0]
		}
		ans[i] = sum
	}

	return ans
}
