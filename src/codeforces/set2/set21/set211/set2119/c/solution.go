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

func drive(reader *bufio.Reader) int {
	var n, l, r, k int
	fmt.Fscan(reader, &n, &l, &r, &k)
	return solve(n, l, r, k)
}

func solve(n int, l int, r int, k int) int {
	if n == 1 {
		return l
	}
	if n == 2 {
		// a & b = a ^ b => -1
		return -1
	}

	if n&1 == 1 {
		// 奇数个 l xor = l
		return l
	}

	h := bits.Len(uint(l))
	a := 1 << h
	// 1 << h <= r 是必须的
	if a > r {
		return -1
	}
	if k <= n-2 {
		return l
	}
	return a
}
