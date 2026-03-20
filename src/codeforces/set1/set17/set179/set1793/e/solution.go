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

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([]int, q)
	for i := range q {
		fmt.Fscan(reader, &queries[i])
	}
	return solve(a, queries)
}

func solve(a []int, queries []int) []int {
	slices.Sort(a)
	n := len(a)
	f := make([]int, n+1)

	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		f[i] = f[i-1]
		if a[i-1] <= i {
			f[i] = max(f[i], f[i-a[i-1]]+1)
			// 分配a[i]个人和第i人一组看同一本书，前面i-a[i]个人看f[i-a[i]]本书，
			// 其他的n-i个人，每个人看一本书，所以，一共可以看 f[i-a[i]] + 1 + n - i 本书
			book := f[i-a[i-1]] + 1 + n - i
			ans[book] = max(ans[book], i)
		} else {
			// 从1..到 a[i-1]这些人（超过了i)一起看一本书，他们肯定都满意
			book := n - a[i-1] + 1
			ans[book] = max(ans[book], i)
		}
	}

	for i := n - 1; i > 0; i-- {
		ans[i] = max(ans[i], ans[i+1])
	}

	res := make([]int, len(queries))
	for i, k := range queries {
		res[i] = ans[k]
	}
	return res
}
