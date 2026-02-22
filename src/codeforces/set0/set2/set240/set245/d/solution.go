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
	var n int
	fmt.Fscan(reader, &n)
	b := make([][]int, n)
	for i := range n {
		b[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &b[i][j])
		}
	}
	a := solve(b)
	s := fmt.Sprintf("%v", a)
	fmt.Fprintln(writer, s[1:len(s)-1])
}

func solve(b [][]int) []int {
	n := len(b)
	a := make([]int, n)
	for i := range n {
		for j := range n {
			if i != j {
				a[i] |= b[i][j]
			}
		}
	}
	return a
}
