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
		a, b := drive(reader)
		fmt.Fprintln(writer, a, b)
	}
}

func drive(reader *bufio.Reader) (int, int) {
	var x, y int
	fmt.Fscan(reader, &x, &y)
	return solve(x, y)
}

func solve(x, y int) (int, int) {
	_ = x
	_ = y
	// TODO: solve by hand first.
	return 0, 0
}
