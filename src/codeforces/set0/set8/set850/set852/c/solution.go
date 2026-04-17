package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for i, x := range res {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, x)
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	return solve(b)
}

func solve(b []int) []int {
	n := len(b)

	type pair struct {
		diff int
		idx  int
	}

	items := make([]pair, n)
	for i := 0; i < n; i++ {
		items[i] = pair{b[i] + b[(i+1)%n], i}
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].diff != items[j].diff {
			return items[i].diff < items[j].diff
		}
		return items[i].idx < items[j].idx
	})

	res := make([]int, n)
	for rank, it := range items {
		res[it.idx] = rank
	}
	return res
}
