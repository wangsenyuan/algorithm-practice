package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if len(res) == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Println(res[0], res[1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	grid := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &grid[i])
	}
	_ = m
	return solve(grid)
}

func solve(grid []string) []int {
	// 找到不在一行的两个, 那么它们必须在同一列
	n := len(grid)
	m := len(grid[0])

	row := make([]int, n)
	col := make([]int, m)
	var sum int
	for i := range n {
		for j := range m {
			if grid[i][j] == '*' {
				row[i]++
				col[j]++
				sum++
			}
		}
	}
	if sum == 0 {
		return []int{1, 1}
	}

	for i := range n {
		for j := range m {
			w := row[i] + col[j]
			if grid[i][j] == '*' {
				w--
			}
			if w == sum {
				return []int{i + 1, j + 1}
			}
		}
	}

	return nil
}
