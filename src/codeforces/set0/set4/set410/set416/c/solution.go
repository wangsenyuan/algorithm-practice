package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, best, accepted := drive(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d %d\n", len(accepted), best))
	for _, cur := range accepted {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) (requests [][]int, tables []int, best int, accepted [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	requests = make([][]int, n)
	for i := range n {
		var c, p int
		fmt.Fscan(reader, &c, &p)
		requests[i] = []int{c, p}
	}
	var m int
	fmt.Fscan(reader, &m)
	tables = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &tables[i])
	}
	best, accepted = solve(requests, tables)
	return
}

type request struct {
	id int
	c  int
	p  int
}

type table struct {
	id int
	r  int
}

type data struct {
	val  int
	take int
}

func solve(requests [][]int, tables []int) (best int, accepted [][]int) {
	n := len(requests)
	rs := make([]request, n)
	for i := range n {
		rs[i] = request{i + 1, requests[i][0], requests[i][1]}
	}
	m := len(tables)
	ts := make([]table, m)
	for i := range m {
		ts[i] = table{i + 1, tables[i]}
	}

	slices.SortFunc(rs, func(a, b request) int {
		return b.c - a.c
	})

	slices.SortFunc(ts, func(a, b table) int {
		return b.r - a.r
	})

	dp := make([][]data, m+1)
	for i := range m + 1 {
		dp[i] = make([]data, n+1)
	}

	for i := range m {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = dp[i+1][j]
			if dp[i+1][j+1].val < dp[i][j+1].val {
				dp[i+1][j+1] = dp[i][j+1]
			}
			if ts[i].r >= rs[j].c && dp[i][j].val+rs[j].p >= dp[i+1][j+1].val {
				dp[i+1][j+1].val = dp[i][j].val + rs[j].p
				dp[i+1][j+1].take = j
			}
		}
	}

	best = dp[m][n].val

	for i, j := m, n; i > 0; i-- {
		if dp[i][j].val == dp[i-1][j].val {
			continue
		}
		// dp[i][j].val > dp[i-1][j].val
		j = dp[i][j].take
		accepted = append(accepted, []int{rs[j].id, ts[i-1].id})
	}

	return
}
