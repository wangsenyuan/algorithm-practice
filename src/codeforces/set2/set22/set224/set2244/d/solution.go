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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a []int, b []int) int {
	n := len(a)
	slices.Sort(b)

	var res int
	for i, j := 0, 0; j < len(b); j++ {
		var cur int
		for i < b[j] {
			cur += a[i]
			i++
		}
		res += abs(cur)
	}

	for i := b[len(b)-1]; i < n; i++ {
		res += a[i]
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
