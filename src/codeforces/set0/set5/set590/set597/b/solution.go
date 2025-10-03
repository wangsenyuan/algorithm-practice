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
	var n int
	fmt.Fscan(reader, &n)
	orders := make([][]int, n)
	for i := 0; i < n; i++ {
		orders[i] = make([]int, 2)
		fmt.Fscan(reader, &orders[i][0], &orders[i][1])
	}
	return solve(orders)
}

func solve(orders [][]int) int {
	slices.SortFunc(orders, func(a, b []int) int {
		return a[1] - b[1]
	})

	var res int
	last := -1
	for _, cur := range orders {
		if cur[0] > last {
			res++
			last = cur[1]
		}
	}
	return res
}
