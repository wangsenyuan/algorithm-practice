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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}
	return solve(p)
}

func solve(p []int) int {
	n := len(p)
	res := 1
	dp := 1
	// dp表示[?:i]
	for i := 1; i < n; i++ {
		if p[i-1] > p[i] {
			// 这里怎么更新dp?
			// dp = [?...i-1] 的最大lds的sum
			// = sum(i-1 - j) where p[j] ... p[i-1]是递减序列
			// = sum(i - j) = sum(i - (i - 1)) + dp = cur + dp
			dp += i
		}
		dp++
		res += dp
	}

	return res
}
