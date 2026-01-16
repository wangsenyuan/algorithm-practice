package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	b := make([]int, m)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, b []int) int {
	var arr []pair
	for _, v := range a {
		arr = append(arr, pair{v, 0})
	}
	for _, v := range b {
		arr = append(arr, pair{v, 1})
	}

	slices.SortFunc(arr, func(x, y pair) int {
		return x.first - y.first
	})
	n := len(arr)

	// dp[i] = 把a增加到arr[i]的值
	dp := make([]int, n)
	var sum int
	var cnt int
	for i, cur := range arr {
		if cur.second == 0 {
			sum += cur.first
			cnt++
		}
		dp[i] = cnt*arr[i].first - sum
	}

	best := 1 << 60

	sum = 0
	cnt = 0
	for i := n - 1; i >= 0; i-- {
		cur := arr[i]
		if cur.second == 1 {
			sum += cur.first
			cnt++
		}
		fp := sum - cnt*arr[i].first
		best = min(best, fp+dp[i])
	}

	return best
}
