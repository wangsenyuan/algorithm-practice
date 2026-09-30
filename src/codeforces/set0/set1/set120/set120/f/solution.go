package main

import (
	"bufio"
	"fmt"
	"os"
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
	spiders := make([][][2]int, n)
	for i := range n {
		var beads int
		fmt.Fscan(reader, &beads)
		edges := make([][2]int, beads-1)
		for j := range edges {
			fmt.Fscan(reader, &edges[j][0], &edges[j][1])
		}
		spiders[i] = edges
	}
	return solve(spiders)
}

func solve(spiders [][][2]int) int {
	_ = spiders
	// TODO: solve by hand first.
	return 0
}
