package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := drive(reader)
	var buf bytes.Buffer
	for _, s := range ans {
		buf.WriteString(s)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) []string {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	intervals := make([][2]int, m)
	for i := range m {
		fmt.Fscan(reader, &intervals[i][0], &intervals[i][1])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][2]int, q)
	for i := range q {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(n, intervals, queries)
}

func solve(n int, intervals [][2]int, queries [][2]int) []string {
	// cover两种情况，一种是由(i, j)各cover前后两端，且它们*相连*
	// 第二种是i cover整段，但是有一段在中间
	pending := make([][]int, n+1)
	for i, cur := range queries {
		r := cur[1]
		pending[r] = append(pending[r], i)
	}

	todo := make([][]int, n+1)
	for _, cur := range intervals {
		l, r := cur[0], cur[1]
		todo[r] = append(todo[r], l)
	}

	ans := make([]string, len(queries))

	// dp[i]表示在i处开始，已经结束的最大的cover距离
	dp := make([]int, n+1)
	var nearest int

	check := func(l int, r int) string {

		i1 := sort.SearchInts(todo[r], l)
		i2 := sort.SearchInts(todo[r], l+1)
		if i2-i1 > 1 {
			// 有至少两个(l..r)的区间
			return "Yes"
		}
		if i2-i1 == 1 && (nearest >= l || i1+1 < len(todo[r])) {
			// 有一个[l...r]的区间, 且包含一个小区间,
			return "Yes"
		}

		if dp[l] > 0 {
			r1 := l + dp[l]
			i2 := sort.SearchInts(todo[r], r1+1) - 1
			// todo[r][i2] <= r1
			if i2 >= 0 && l < todo[r][i2] {
				return "Yes"
			}
		}

		return "No"
	}

	for r := 1; r <= n; r++ {
		slices.Sort(todo[r])

		for _, id := range pending[r] {
			l := queries[id][0]
			ans[id] = check(l, r)
		}
		for _, l := range todo[r] {
			dp[l] = r - l + 1
			nearest = max(nearest, l)
		}
	}

	return ans
}
