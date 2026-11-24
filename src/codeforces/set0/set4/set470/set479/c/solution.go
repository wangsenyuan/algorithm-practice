package main

import (
	"bufio"
	"cmp"
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
	var n int
	fmt.Fscan(reader, &n)

	exams := make([][2]int, n)
	for i := range n {
		fmt.Fscan(reader, &exams[i][0], &exams[i][1])
	}

	return solve(exams)
}

func solve(exams [][2]int) int {
	slices.SortFunc(exams, func(a [2]int, b [2]int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	var ans int

	for _, cur := range exams {
		if ans <= cur[1] {
			ans = cur[1]
		} else {
			ans = cur[0]
		}
	}

	return ans
}
