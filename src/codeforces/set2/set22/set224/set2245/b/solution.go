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
	var n, c int
	fmt.Fscan(reader, &n, &c)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	return solve(c, a)
}

func solve(c int, a []int) int {
	slices.Sort(a)

	n := len(a)

	var res int
	for l, r := 0, n-1; l <= r; {
		if l < r && a[l] < c {
			res += a[r] - c
			l++
			r--
		} else {
			// l <= r || a[l] >= c
			for i := l; i <= r; i++ {
				res += a[i] - c
			}
			break
		}
	}
	return res
}
