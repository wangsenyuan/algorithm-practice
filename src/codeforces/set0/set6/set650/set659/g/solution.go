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
	fmt.Fprintln(writer, drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	h := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &h[i])
	}
	return solve(h)
}

func solve(h []int) int {
	// TODO
	_ = h
	return 0
}
