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
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Fprintln(writer, s[1:len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(k, a, b)
}

func solve(k int, a, b []int) []int {
	// TODO
	_ = k
	_ = a
	_ = b
	return make([]int, len(a))
}
