package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, l, res := drive(reader)
	fmt.Println(l, len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (a []int, m int, l int, res []int) {
	var n int

	fmt.Fscan(reader, &n, &m)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	l, res = solve(m, a)
	return
}

func solve(m int, a []int) (int, []int) {
	// n := len(a)
	dp := make([]int, m+1)
	for _, v := range a {
		if v <= m {
			dp[v]++
		}
	}
	arr := slices.Clone(a)
	slices.Sort(arr)
	arr = slices.Compact(arr)

	cnt := slices.Clone(dp)

	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		for u := 2 * v; u <= m; u += v {
			dp[u] += cnt[v]
		}
	}

	best := slices.Max(dp)

	if best == 0 {
		return 1, nil
	}
	var num int
	for i, v := range dp {
		if v == best {
			num = i
			break
		}
	}
	var res []int

	for i, v := range a {
		if num%v == 0 {
			res = append(res, i+1)
		}
	}

	return num, res
}
