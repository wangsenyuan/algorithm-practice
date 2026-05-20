package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, len(res))
	for _, x := range res {
		fmt.Fprintln(writer, x[0], x[1])
	}
}

func drive(reader *bufio.Reader) (c []int, res [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	c = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	res = solve(slices.Clone(c))
	return
}

func solve(c []int) [][]int {
	n := len(c)
	var res [][]int

	a := make([]int, n)
	pos := make([]int, n)
	pos1 := make([]int, n)
	for i := range n {
		a[i] = i
		pos[i] = i
		c[i]--
		pos1[c[i]] = i
	}

	swap := func(u int, v int) {
		res = append(res, []int{u + 1, v + 1})
		pos[u], pos[v] = pos[v], pos[u]
		a[pos[u]] = u
		a[pos[v]] = v
	}

	for i := range n {
		for j := i - 1; j >= 0; j-- {
			swap(i, a[j])
		}
		for j := 1; j <= i; j++ {
			if pos1[a[j]] < pos1[i] {
				swap(a[j], i)
			}
		}
	}

	return res
}
