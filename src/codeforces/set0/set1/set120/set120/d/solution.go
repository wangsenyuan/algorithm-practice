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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	c := make([][]int, n)
	for i := range n {
		c[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &c[i][j])
		}
	}
	var a, b, cc int
	fmt.Fscan(reader, &a, &b, &cc)
	return solve(c, a, b, cc)
}

func solve(c [][]int, a, b, cc int) int {
	// TODO
	_ = c
	_ = a
	_ = b
	_ = cc
	return 0
}
