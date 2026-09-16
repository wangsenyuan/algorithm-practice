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

func drive(reader *bufio.Reader) []int64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	del := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &del[i])
	}
	return solve(a, del)
}

func solve(a []int, del []int) []int64 {
	// TODO
	_ = a
	return make([]int64, len(del))
}
