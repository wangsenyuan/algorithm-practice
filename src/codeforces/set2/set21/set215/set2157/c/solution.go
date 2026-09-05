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
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n, k, q int
	fmt.Fscan(reader, &n, &k, &q)
	queries := make([][3]int, q)
	for i := range queries {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(n, k, queries)
}

func solve(n, k int, queries [][3]int) []int {
	// TODO
	_ = k
	_ = queries
	return make([]int, n)
}
