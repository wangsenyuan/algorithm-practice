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
		res := drive(reader)
		for _, x := range res {
			fmt.Fprintln(writer, x)
		}
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, m)
	for i := range m {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		queries[i] = []int{l, r}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	// poby 先操作的结果
	dp := make([]int, n+1)
	// rekkles 先操作的结果
	dp2 := make([]int, n+1)

	for i, v := range a {
		dp[i+1] = play(v) + dp[i]
		dp2[i+1] = play(v+1) + dp2[i]
	}

	res := make([]int, len(queries))
	for i, cur := range queries {
		l, r := cur[0]-1, cur[1]-1
		res[i] = dp2[r+1] - dp2[l]
		diff := (dp2[r+1] - dp2[l]) - (dp[r+1] - dp[l])
		diff = (diff + 1) / 2
		res[i] -= diff
	}
	return res
}

func play(x int) int {
	var res int
	for x > 1 {
		res++
		x >>= 1
		if x == 1 {
			break
		}
		x++
	}
	return res
}
