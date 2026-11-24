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

	s, ops := drive(reader)
	fmt.Fprintln(writer, s, len(ops))
	for _, op := range ops {
		fmt.Fprintln(writer, op[0], op[1])
	}
}

func drive(reader *bufio.Reader) (int, [][2]int) {
	var n, k int
	fmt.Fscan(reader, &n, &k)

	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	return solve(k, a)
}

type node struct {
	id  int
	val int
}

func solve(k int, a []int) (int, [][2]int) {
	n := len(a)
	arr := make([]node, n)
	for i, v := range a {
		arr[i] = node{i, v}
	}

	// n <= 100, k <= 1000
	var ops [][2]int

	for k > 0 {
		slices.SortFunc(arr, func(a node, b node) int {
			return a.val - b.val
		})
		first := arr[0]
		last := arr[n-1]
		if first.val == last.val {
			break
		}
		k--
		arr[n-1].val--
		arr[0].val++
		ops = append(ops, [2]int{last.id + 1, first.id + 1})
	}
	slices.SortFunc(arr, func(a node, b node) int {
		return a.val - b.val
	})

	diff := arr[n-1].val - arr[0].val

	return diff, ops
}
