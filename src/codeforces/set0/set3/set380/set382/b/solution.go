package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var a, b, w, x, c int
	fmt.Fscan(reader, &a, &b, &w, &x, &c)
	return solve(a, b, w, x, c)
}

type data struct {
	val int
	sec int
}

func solve(a int, b int, w int, x int, c int) int {
	if c <= a {
		return 0
	}

	y := w - x
	d := c*w - a*w - b

	check := func(k int) bool {
		if (b+x-1)/x >= k && b >= k*x {
			return c-k <= a
		}
		// b < k * x
		// c - k <= a - (k * x - b) / w
		// (c - k) * w <= a * w - (k * x - b)
		// c * w - k * w <= a * w - k * x + b
		// c * w - a * w - b <= k * (w - x)
		// d <= k * (w - x)
		if (d+y-1)/y >= k && d > k*y {
			return false
		}
		return true
	}

	return sort.Search(1e15, check)
}
