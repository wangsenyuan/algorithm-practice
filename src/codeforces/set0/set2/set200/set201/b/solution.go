package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res.cost)
	fmt.Println(res.li, res.lj)
}

type answer struct {
	cost     int
	li, lj   int
}

func drive(reader *bufio.Reader) answer {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	grid := make([][]int, n)
	for i := range n {
		grid[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &grid[i][j])
		}
	}
	return solve(grid)
}

func solve(grid [][]int) answer {
	// TODO: solve by hand first.
	return answer{}
}
