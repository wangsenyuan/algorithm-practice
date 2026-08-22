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
		for _, line := range drive(reader) {
			fmt.Fprintln(writer, line)
		}
	}
}

func drive(reader *bufio.Reader) []string {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	var s string
	fmt.Fscan(reader, &s)
	queries := readQueries(reader, q)
	_ = n
	return solve(s, queries)
}

func readQueries(reader *bufio.Reader, q int) [][3]int {
	queries := make([][3]int, q)
	for i := range queries {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return queries
}

func solve(s string, queries [][3]int) []string {
	// TODO: solve by hand first.
	_ = s
	return make([]string, len(queries))
}
