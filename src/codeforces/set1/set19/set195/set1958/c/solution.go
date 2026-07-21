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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]int, t)
	for i := range t {
		var n int
		var k int
		fmt.Fscan(reader, &n, &k)
		res[i] = solve(n, k)
	}
	return res
}

func solve(n int, k int) int {
	return n - bits.TrailingZeros(uint(k))
}
