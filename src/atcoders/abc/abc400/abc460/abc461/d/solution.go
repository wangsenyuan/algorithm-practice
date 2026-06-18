package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var h, w, k int
	fmt.Fscan(reader, &h, &w, &k)
	grid := make([]string, h)
	for i := range h {
		fmt.Fscan(reader, &grid[i])
	}
	return solve(k, grid)
}

func solve(k int, grid []string) int {
	if len(grid) > len(grid[0]) {
		grid = rotate(grid)
	}

	// n <= m, n * m <= 500 => n <= 22
	n := len(grid)
	m := len(grid[0])
	// 在确定r1, r2的情况下, 迭代c
	col := make([]int, m)

	var res int

	for r1 := range n {
		clear(col)
		for r2 := r1; r2 < n; r2++ {
			var todo []int
			var sum int
			todo = append(todo, 0)
			var pos1, pos2 int
			for c := range m {
				if grid[r2][c] == '1' {
					col[c]++
				}
				sum += col[c]
				// sum - s1 >= k
				for pos1 < len(todo) && sum-todo[pos1] > k {
					pos1++
				}
				for pos2 < len(todo) && sum-todo[pos2] >= k {
					pos2++
				}
				res += pos2 - pos1
				todo = append(todo, sum)
			}
		}
	}

	return res
}

func rotate(a []string) []string {
	n := len(a)
	m := len(a[0])
	res := make([]string, m)

	buf := make([]byte, n)
	for j := range m {
		for i := range n {
			buf[i] = a[i][j]
		}
		res[j] = string(buf)
	}

	return res
}
