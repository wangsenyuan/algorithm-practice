package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)
	points := make([][2]int, n)
	for i := range points {
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	return solve(points)
}

func solve(points [][2]int) int64 {
	n := len(points)
	pref := make([]int, n+1)
	hi := make([]int, n+1)
	lo := make([]int, n+1)
	for i := range n + 1 {
		hi[i] = -1
		lo[i] = n + 1
	}
	for _, cur := range points {
		pref[cur[1]] = 1
		hi[cur[0]] = max(hi[cur[0]], cur[1])
		lo[cur[0]] = min(lo[cur[0]], cur[1])
	}
	// 可以快速的计算出 y1...y2中间有多少个点
	for i := 1; i <= n; i++ {
		pref[i] += pref[i-1]
	}

	hi_suf := make([]int, n+2)
	lo_suf := make([]int, n+2)
	for i := range n + 2 {
		hi_suf[i] = -1
		lo_suf[i] = n + 1
	}
	for i := n; i > 0; i-- {
		hi_suf[i] = max(hi_suf[i+1], hi[i])
		lo_suf[i] = min(lo_suf[i+1], lo[i])
	}

	next := make([]int, n+1)
	for i := n; i >= 0; i-- {
		next[i] = i + 1
		if i+1 <= n && hi[i+1] == -1 {
			next[i] = next[i+1]
		}
	}

	var ans int

	hi_pre := -1
	lo_pre := n + 1

	for i := next[0]; i < n; i = next[i] {
		// 如果在i..i+1中间分
		i1 := next[i]
		hi_pre = max(hi_pre, hi[i])
		lo_pre = min(lo_pre, lo[i])
		y1 := max(lo_pre, lo_suf[i1])
		y2 := min(hi_pre, hi_suf[i1])
		if y1 < y2 {
			cnt := pref[y2] - pref[y1-1]
			ans += max(0, cnt-1)
		}
	}

	return int64(ans)
}
