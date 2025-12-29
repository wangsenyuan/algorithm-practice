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
		res := drive(reader)
		output(writer, res)
	}
}

func output(writer *bufio.Writer, a []int) {
	for i := range a {
		fmt.Fprint(writer, a[i], " ")
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

func solve(a []int, k int) []int {
	for i := range a {
		a[i] += (a[i]) % (k + 1) * k
	}

	return a
}