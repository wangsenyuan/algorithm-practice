package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if len(res) == 1 {
		fmt.Println(res[0])
	} else if len(res) == 2 {
		fmt.Println(res[0], res[1])
	}
}

func drive(reader *bufio.Reader) []int {
	var h, w int
	fmt.Fscan(reader, &h, &w)

	grid := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(reader, &grid[i])
	}

	return solve(h, w, grid)
}

func solve(h int, w int, grid []string) []int {
	vis := make([][]bool, h)
	for i := range h {
		vis[i] = make([]bool, w)
	}

	var x, y int

	for {
		if vis[x][y] {
			return []int{-1}
		}
		vis[x][y] = true
		switch grid[x][y] {
		case 'U':
			if x == 0 {
				return []int{x + 1, y + 1}
			}
			x--
		case 'D':
			if x == h-1 {
				return []int{x + 1, y + 1}
			}
			x++
		case 'L':
			if y == 0 {
				return []int{x + 1, y + 1}
			}
			y--
		case 'R':
			if y == w-1 {
				return []int{x + 1, y + 1}
			}
			y++
		}
	}
}
