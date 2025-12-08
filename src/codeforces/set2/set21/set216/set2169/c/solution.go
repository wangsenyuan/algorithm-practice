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
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	// 1 * 1 - 0
	var best, sum int
	n := len(a)
	var ans int
	for i := 1; i <= n; i++ {
		best = min(best, i*i-i-sum)
		sum += a[i-1]
		tmp := i*i + i - sum
		ans = max(ans, tmp-best)
	}

	return ans + sum
}
