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
		res := drive(reader)
		fmt.Fprintln(writer, len(res))
		for i := range res {
			fmt.Fprint(writer, res[i], " ")
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) []int {
	var x, y int
	fmt.Fscan(reader, &x, &y)
	return solve(x, y)
}

func solve(x int, y int) []int {
	// 一个最大值，一个最小值
	n := 2 * (x - y)

	res := make([]int, n)
	z := y
	delta := 1
	for i := range n {
		res[i] = z
		z += delta
		if z > x {
			z = x - 1
			delta = -1
		}
	}

	return res
}
