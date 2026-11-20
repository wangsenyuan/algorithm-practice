package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var n int
	var k int
	fmt.Fscan(reader, &n, &k)
	days := make([][]int, n)
	for i := range n {
		days[i] = make([]int, 3)
		fmt.Fscan(reader, &days[i][0], &days[i][1], &days[i][2])
	}
	return solve(k, days)
}

func solve(k int, days [][]int) int64 {
	sum := make([]int, 2)
	var cost int
	for _, cur := range days {
		sum[0] += cur[0]
		sum[1] += cur[1]
		cost += cur[0] * cur[2]
	}

	if sum[0] > k || sum[1] < k {
		return -1
	}
	k -= sum[0]
	// 还需要吃k个单位的
	slices.SortFunc(days, func(a []int, b []int) int {
		return a[2] - b[2]
	})
	for _, cur := range days {
		if k == 0 {
			break
		}
		w := min(k, cur[1]-cur[0])
		k -= w
		cost += w * cur[2]
	}

	return int64(cost)
}
