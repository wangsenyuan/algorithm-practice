package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for _, v := range drive(reader) {
		fmt.Println(v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	items := make([][]int, n)
	for i := range n {
		items[i] = make([]int, 2)
		fmt.Fscan(reader, &items[i][0], &items[i][1])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(items, queries)
}

func solve(items [][]int, queries [][]int) []int {
	type query struct {
		l int
		r int
		c int
	}

	n := len(items)
	qs := make([]query, len(queries))
	ids := make([]int, len(queries))
	for i, cur := range queries {
		qs[i] = query{cur[0] - 1, cur[1], cur[2]}
		ids[i] = i
	}

	dp := make([]data, n+1)
	ans := make([]int, len(queries))

	var process func(int, int, []int)
	process = func(left int, right int, queryIDs []int) {
		if len(queryIDs) == 0 {
			return
		}
		if left+1 == right {
			weight, value := items[left][0], items[left][1]
			for _, id := range queryIDs {
				if weight <= qs[id].c {
					ans[id] = value
				}
			}
			return
		}

		mid := (left + right) / 2
		leftIDs := make([]int, 0, len(queryIDs)/2)
		rightIDs := make([]int, 0, len(queryIDs)/2)
		crossingIDs := make([]int, 0, len(queryIDs))
		maxCapacity := 0

		for _, id := range queryIDs {
			cur := qs[id]
			if cur.r <= mid {
				leftIDs = append(leftIDs, id)
			} else if cur.l >= mid {
				rightIDs = append(rightIDs, id)
			} else {
				crossingIDs = append(crossingIDs, id)
				maxCapacity = max(maxCapacity, cur.c)
			}
		}

		if len(crossingIDs) > 0 {
			clear(dp[mid][:maxCapacity+1])

			for i := mid - 1; i >= left; i-- {
				weight, value := items[i][0], items[i][1]
				previous, current := &dp[i+1], &dp[i]
				copy(current[:min(weight, maxCapacity+1)], previous[:min(weight, maxCapacity+1)])
				for capacity := weight; capacity <= maxCapacity; capacity++ {
					current[capacity] = max(previous[capacity], previous[capacity-weight]+value)
				}
			}

			for i := mid + 1; i <= right; i++ {
				weight, value := items[i-1][0], items[i-1][1]
				previous, current := &dp[i-1], &dp[i]
				copy(current[:min(weight, maxCapacity+1)], previous[:min(weight, maxCapacity+1)])
				for capacity := weight; capacity <= maxCapacity; capacity++ {
					current[capacity] = max(previous[capacity], previous[capacity-weight]+value)
				}
			}

			for _, id := range crossingIDs {
				cur := qs[id]
				for leftCapacity := 0; leftCapacity <= cur.c; leftCapacity++ {
					ans[id] = max(ans[id], dp[cur.l][leftCapacity]+dp[cur.r][cur.c-leftCapacity])
				}
			}
		}

		process(left, mid, leftIDs)
		process(mid, right, rightIDs)
	}

	process(0, n, ids)
	return ans
}

const W = 501

type data [W]int
