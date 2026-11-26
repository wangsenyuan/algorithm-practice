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

	if drive(reader) {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)

	first := make([]string, n)
	last := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &first[i], &last[i])
	}

	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}

	return solve(first, last, p)
}

func solve(first, last []string, p []int) bool {
	var prev string

	for _, i := range p {
		i--
		u, v := first[i], last[i]
		if u > v {
			u, v = v, u
		}
		// u < v
		if len(prev) == 0 || prev < u {
			prev = u
			continue
		}
		if len(prev) == 0 || prev < v {
			prev = v
			continue
		}
		return false
	}
	return true
}
