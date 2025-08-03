package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	var buf bytes.Buffer
	for _, x := range ans {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func process(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	// 只需要关注第一个s的高度
	// n := len(a)
	first := a[0]

	ans := make([]int, len(queries))

	for i, cur := range queries {
		w, h := cur[0], cur[1]
		// 0.... w - 1
		ans[i] = max(first, a[w-1])
		first = ans[i] + h
	}
	return ans
}
