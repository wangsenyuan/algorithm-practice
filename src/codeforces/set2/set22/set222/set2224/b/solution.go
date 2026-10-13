package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)

	mx := slices.Max(a)

	if mx == 0 {
		return n
	}
	freq := make(map[int]int)
	for _, v := range a {
		freq[v]++
	}

	// 0, 1, 2, 3, ... mex - 1
	// 假设最大值是x
	// 那么x放在第一个位置, n * x + ... sum(1, ... mex) + (n - mex - 1) * mex

	// 先放置最大值, 或者在0, 1, ... mex 后放置最大值
	var mex int
	for freq[mex] > 0 {
		mex++
	}

	if mex < mx {
		return mx*n + mex*(mex-1)/2 + (n-mex)*mex
	}
	return mx*n + mx*(mx-1)/2 + (n-mex+1)*mex
}
