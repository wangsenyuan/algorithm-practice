package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	fmt.Println(res.cost)
	fmt.Println(res.li, res.lj)
}

type answer struct {
	cost   int
	li, lj int
}

func drive(reader *bufio.Reader) ([][]int, answer) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	grid := make([][]int, n)
	for i := range n {
		grid[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &grid[i][j])
		}
	}
	return grid, solve(grid)
}

func square(x int) int {
	return x * x
}

func solve(grid [][]int) answer {
	n := len(grid)
	m := len(grid[0])
	// (i - r) ** 2 = i ** 2 + r ** 2 - 2 * i * r
	row := make([]int, n)
	col := make([]int, m)
	for i := range n {
		for j := range m {
			row[i] += grid[i][j]
			col[j] += grid[i][j]
		}
	}

	var r0 int
	var best0 = 1 << 60
	for r := range n + 1 {
		var sum int
		for i := range n {
			// 每个格子的高度是4
			sum += row[i] * square(r*4-(i*4+2))
		}

		if sum < best0 {
			best0 = sum
			r0 = r
		}
	}

	var c0 int
	var best1 = 1 << 60
	for c := range m + 1 {
		var sum int
		for j := range m {
			sum += col[j] * square(c*4-(j*4+2))
		}

		if sum < best1 {
			best1 = sum
			c0 = c
		}
	}

	return answer{cost: best0 + best1, li: r0, lj: c0}
}
