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

	var n int
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		if solve(a, b) {
			fmt.Fprintln(writer, "Yes")
		} else {
			fmt.Fprintln(writer, "No")
		}
	}
}

func solve(a int, b int) bool {
	prod := int64(a) * int64(b)
	c := cubeRoot(prod)
	if c*c*c != prod {
		return false
	}
	return a%int(c) == 0 && b%int(c) == 0
}

func cubeRoot(x int64) int64 {
	var lo, hi int64 = 0, 1000000
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if mid*mid*mid <= x {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}
