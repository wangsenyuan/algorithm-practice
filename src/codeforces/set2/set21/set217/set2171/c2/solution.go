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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a, b []int) string {
	var x int
	for _, v := range a {
		x ^= v
	}
	for _, v := range b {
		x ^= v
	}
	if x == 0 {
		return "Tie"
	}
	h := bits.Len(uint(x)) - 1

	for i := len(a) - 1; i >= 0; i-- {
		if (a[i]^b[i])&(1<<h) > 0 {
			if i&1 == 0 {
				return "Ajisai"
			}
			return "Mai"
		}
	}

	return "Tie"
}
