package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, v := range res {
		if v {
			fmt.Fprintln(writer, "yes")
		} else {
			fmt.Fprintln(writer, "no")
		}
	}
}

func drive(reader *bufio.Reader) []bool {
	var n, m, x int
	fmt.Fscan(reader, &n, &m, &x)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(x, a, queries)
}

func solve(x int, a []int, queries [][]int) []bool {
	n := len(a)
	qs := make([][]int, n)
	for i, cur := range queries {
		r := cur[1] - 1
		qs[r] = append(qs[r], i)
	}

	mx := slices.Max(a)
	pos := make([]int, mx+1)
	for i := range mx + 1 {
		pos[i] = -1
	}

	ans := make([]bool, len(queries))

	last := -1
	for i, v := range a {
		w := v ^ x
		if w <= mx {
			last = max(last, pos[w])
		}
		pos[v] = i
		for _, j := range qs[i] {
			l := queries[j][0] - 1
			ans[j] = l <= last
		}
	}
	return ans
}
