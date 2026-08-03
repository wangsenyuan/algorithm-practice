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
	if solve(a, b) {
		return "YES"
	}
	return "NO"
}

func solve(a, b []int) bool {
	n := len(a)
	var add int
	for i := n - 1; i >= 0; i-- {
		v := a[i] + add
		if v > b[i] {
			// need to sub
			add = v - b[i]
		} else {
			add = 0
		}
	}
	return add == 0
}
