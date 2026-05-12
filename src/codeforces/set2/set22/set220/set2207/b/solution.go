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
	var n, m, l int
	fmt.Fscan(reader, &n, &m, &l)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(m, l, a)
}

func solve(m int, l int, a []int) int {
	n := len(a)
	left := n
	k := min(m, left+1)
	level := make([]int, k)

	for t := 1; t <= l; t++ {
		level[len(level)-1]++
		slices.SortFunc(level, func(a, b int) int {
			return b - a
		})

		if left > 0 && a[n-left] == t {
			level = level[1:]
			left--
			want := min(m, left+1)
			if len(level) < want {
				level = append(level, 0)
			}
		}
	}

	if len(level) == 0 {
		return 0
	}
	return level[0]
}
