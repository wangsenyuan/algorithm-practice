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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int {
	// slices.Sort(a)
	// 只能买一次?
	var res int
	n := len(a)
	sum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sum[i] = sum[i+1] + a[i]
		if i <= n-k {
			w := (n - i) / k
			free := sum[i] - sum[i+w]
			res = max(res, free)
		}
	}

	return res
}
