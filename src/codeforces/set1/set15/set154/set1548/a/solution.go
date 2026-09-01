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

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	friends := make([][]int, m)
	for i := range m {
		friends[i] = make([]int, 2)
		fmt.Fscan(reader, &friends[i][0], &friends[i][1])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		var typ int
		fmt.Fscan(reader, &typ)
		if typ == 3 {
			queries[i] = []int{3}
			continue
		}
		var u, v int
		fmt.Fscan(reader, &u, &v)
		queries[i] = []int{typ, u, v}
	}
	return solve(n, friends, queries)
}

func solve(n int, friends [][]int, queries [][]int) []int {
	// TODO
	_ = n
	_ = friends
	_ = queries
	return nil
}
