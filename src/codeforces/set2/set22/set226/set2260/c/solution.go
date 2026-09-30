package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

func solve(x, y int) (res int, cnt int) {
	res = x + y
	w := bits.Len(uint(x &^ res))
	mask := 1<<w - 1
	cnt = x&mask - res&mask
	return
}
