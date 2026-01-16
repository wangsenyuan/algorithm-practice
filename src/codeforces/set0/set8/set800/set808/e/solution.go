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
	w := make([]int, n)
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &w[i], &c[i])
	}
	return solve(m, w, c)
}

type data struct {
	cost   int
	weight int
}

func solve(m int, w []int, c []int) int {
	n := len(w)
	arr := make([]data, n)
	for i := range n {
		arr[i] = data{c[i], w[i]}
	}

	slices.SortFunc(arr, func(a data, b data) int {
		// 根据单位cost进行排序
		return b.cost*a.weight - a.cost*b.weight
	})

	dp := make([]int, m+1)
	var sum int
	for _, cur := range arr {
		sum = min(m, sum+cur.weight)
		for j := sum; j >= max(cur.weight, sum-3); j-- {
			dp[j] = max(dp[j], dp[j-cur.weight]+cur.cost)
		}
	}
	return slices.Max(dp)
}
