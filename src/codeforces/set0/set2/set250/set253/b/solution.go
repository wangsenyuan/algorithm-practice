package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()

	reader := bufio.NewReader(r)
	fmt.Fprintln(w, drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	return solve(c)
}

func solve(c []int) int {
	slices.Sort(c)
	n := len(c)
	best := n - 1

	for l, r := 0, 0; l < n; l++ {
		for r < n && c[r] <= 2*c[l] {
			r++
		}
		best = min(best, l+n-r)
		if r == n {
			break
		}
	}

	return best
}
