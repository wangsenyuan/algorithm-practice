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

	s := fmt.Sprintf("%v", drive(reader))
	fmt.Fprintln(writer, s[1:len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	return solve(n, k)
}

func solve(n int, k int) []int {
	// 2k ≤ n, so k pairs of -1 and n-k pairs of +1 give
	// n - |2k - n| = 2k.
	a := make([]int, 2*n)
	for i := range n {
		left, right := 2*i+1, 2*i+2
		if i < k {
			left, right = right, left
		}
		a[2*i], a[2*i+1] = left, right
	}
	return a
}
