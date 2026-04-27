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
		var n, x int
		fmt.Fscan(reader, &n, &x)
		res := solve(n, x)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func solve(n int, x int) []int {
	if n == 1 {
		return []int{x}
	}

	var h int
	for (x>>h)&1 == 1 {
		h++
	}
	// mex(arr) = m
	m := 1 << h
	res := make([]int, n)
	var or int
	for i := 0; i < n-1 && i < m; i++ {
		res[i] = i
		or |= i
	}

	if m > n-1 && or|(n-1) == x {
		res[n-1] = n - 1
	} else {
		res[n-1] = x
	}

	for i := m; i < n; i++ {
		res[i] = x
	}
	return res
}
