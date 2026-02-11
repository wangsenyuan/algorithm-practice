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
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	queries := make([]int, q)
	for i := range q {
		fmt.Fscan(reader, &queries[i])
	}
	return solve(n, queries)
}

func solve(n int, queries []int) []int {

	var f func(p int) int

	f = func(p int) int {
		if p&1 == 1 {
			// 在奇数位置上的数字
			return (p + 1) / 2
		}
		// p是偶数
		return f(p + (n - p/2))
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = f(q)
	}
	return ans
}
