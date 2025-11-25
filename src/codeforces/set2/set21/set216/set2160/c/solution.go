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
	var t int
	fmt.Fscan(reader, &t)
	for range t {
		var n int
		fmt.Fscan(reader, &n)
		res := solve(n)
		if res {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func solve(n int) bool {
	if n == 0 {
		return true
	}
	lo := bits.TrailingZeros(uint(n))
	n >>= lo
	hi := bits.Len(uint(n))
	if hi&1 == 1 && (n>>(hi/2))&1 == 1 {
		return false
	}
	m := bits.Reverse(uint(n))
	m >>= bits.TrailingZeros(m)
	return m == uint(n)
}
