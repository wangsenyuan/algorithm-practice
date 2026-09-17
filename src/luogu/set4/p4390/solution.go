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

	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var op, w int
	fmt.Fscan(reader, &op, &w)
	var ops [][]int
	for {
		fmt.Fscan(reader, &op)
		if op == 3 {
			break
		}
		if op == 1 {
			var x, y, a int
			fmt.Fscan(reader, &x, &y, &a)
			ops = append(ops, []int{1, x, y, a})
			continue
		}
		var x1, y1, x2, y2 int
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		ops = append(ops, []int{2, x1, y1, x2, y2})
	}
	return solve(w, ops)
}

func solve(w int, ops [][]int) []int {
	// TODO
	_ = w
	nq := 0
	for _, op := range ops {
		if op[0] == 2 {
			nq++
		}
	}
	return make([]int, nq)
}
