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

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a []int, b []int) string {
	play := func(arr []int) []int {
		var sum int
		for _, v := range arr {
			sum ^= v
		}
		var res []int
		for _, v := range arr {
			res = append(res, v^sum)
		}
		return append(res, sum)
	}
	x := play(a)
	y := play(b)
	slices.Sort(x)
	slices.Sort(y)
	if slices.Equal(x, y) {
		return "YES"
	}
	return "NO"
}
