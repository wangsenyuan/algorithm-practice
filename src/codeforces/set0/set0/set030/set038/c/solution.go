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

	fmt.Fprintln(writer, drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, l int
	fmt.Fscan(reader, &n, &l)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, l, a)
}

func solve(n, l int, a []int) int {
	var best int

	x := slices.Max(a)

	for d := l; d <= x; d++ {
		var tmp int
		for _, v := range a {
			tmp += v / d * d
		}
		best = max(best, tmp)
	}

	return best
}
