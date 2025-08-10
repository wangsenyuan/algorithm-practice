package p3643

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	for i := x; i < x+k/2; i++ {
		d1 := i - x
		for j := y; j < y+k; j++ {
			grid[i][j], grid[x+k-1-d1][j] = grid[x+k-1-d1][j], grid[i][j]
		}
	}

	return grid
}
