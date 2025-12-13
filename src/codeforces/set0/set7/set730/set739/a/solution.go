package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	mex, ans := drive(reader)
	fmt.Println(mex)
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (mex int, ans []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	queries := make([][]int, m)
	for i := range m {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		queries[i] = []int{l, r}
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) (int, []int) {
	// 长度为x时，mex为x+1
	mex := n
	for _, cur := range queries {
		l, r := cur[0], cur[1]
		mex = min(mex, r-l+1)
	}

	ans := make([]int, n)
	for i := range n {
		ans[i] = i % mex
	}
	return mex, ans
}
