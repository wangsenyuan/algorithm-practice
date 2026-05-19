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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

func solve(a []int, k int) int {
	slices.Sort(a)
	slices.Reverse(a)
	n := len(a)

	var res int
	for i := 0; i+1 < n; i += 2 {
		res += a[i] - a[i+1]
	}

	res = max(0, res-k)
	if n&1 == 1 {
		res += a[n-1]
	}

	return res
}
