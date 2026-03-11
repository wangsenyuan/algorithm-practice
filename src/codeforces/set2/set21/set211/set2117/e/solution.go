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
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a []int, b []int) int {
	n := len(a)

	pos := make([][]int, 3)
	for i := range 3 {
		pos[i] = make([]int, n+1)
	}

	play := func(a []int, b []int) int {
		for i := range 3 {
			for j := range n + 1 {
				pos[i][j] = -1
			}
		}

		var best int
		for i := n - 1; i >= 0; i-- {
			if pos[i&1][a[i]] < 0 {
				pos[i&1][a[i]] = i
			}
			v := b[i]
			d := i & 1
			if pos[2][v] > 0 || pos[d][v] >= i || pos[d^1][v] > i+1 {
				best = max(best, i+1)
			}
			if pos[2][v] < 0 {
				pos[2][v] = i
			}
		}
		return best
	}

	return max(play(a, b), play(b, a))
}
