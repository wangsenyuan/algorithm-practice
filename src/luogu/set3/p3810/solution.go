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

	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([][3]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i][0], &a[i][1], &a[i][2])
	}
	return solve(k, a)
}

func solve(k int, a [][3]int) []int {
	// TODO
	_ = k
	return make([]int, len(a))
}
