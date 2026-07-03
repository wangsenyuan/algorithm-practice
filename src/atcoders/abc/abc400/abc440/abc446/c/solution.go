package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Println(drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n, d int
	fmt.Fscan(reader, &n, &d)
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(n, d, a, b)
}

func solve(n, d int, a, b []int) int {
	var stock int
	for l, r := 0, 0; r < n; r++ {
		stock += a[r]

		// b[r]
		var cnt int
		for l <= r && cnt+a[l] <= b[r] {
			cnt += a[l]
			l++
		}
		if l <= r {
			a[l] -= b[r] - cnt
			stock -= b[r]
		} else {
			stock -= cnt
		}

		if l == r-d {
			stock -= a[l]
			l++
		}
	}

	return stock
}
