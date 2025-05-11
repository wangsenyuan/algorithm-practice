package p3548

func canPartitionGrid(grid [][]int) bool {
	n := len(grid)
	m := len(grid[0])
	row := make([]int, n)
	col := make([]int, m)
	var sum int
	for i := range n {
		for j := range m {
			row[i] += grid[i][j]
			col[j] += grid[i][j]
			sum += grid[i][j]
		}
	}
	if sum%2 == 0 {
		var tmp int
		for i := range n {
			tmp += row[i]
			if tmp*2 == sum {
				return true
			}
		}
		tmp = 0
		for i := range m {
			tmp += col[i]
			if tmp*2 == sum {
				return true
			}
		}
	}

	for range 4 {
		if solve(grid, sum) {
			return true
		}
		buf := make([][]int, len(grid[0]))

		for i := range len(buf) {
			buf[i] = make([]int, len(grid))
		}

		for i := range len(grid) {
			for j := range len(grid[i]) {
				buf[j][len(grid)-1-i] = grid[i][j]
			}
		}

		grid = buf
	}

	return false
}

func solve(grid [][]int, sum int) bool {
	// 只会从上到下处理
	// (sum - x) % 2 = 0
	n := len(grid)
	m := len(grid[0])
	row := make([]int, n)
	var row_sum int
	for i := range n {
		for j := range m {
			row[i] += grid[i][j]
		}
		row_sum += row[i]
	}

	if m > 1 {
		col := make(map[int]int)
		for i := 0; i+1 < n; i++ {
			row_sum -= row[i]
			x := sum - 2*row_sum
			if i == 0 {
				if grid[i][0] == x || grid[i][m-1] == x {
					return true
				}
			}
			for j := range m {
				col[grid[i][j]] = j
			}
			if i > 0 {
				if _, ok := col[x]; ok {
					return true
				}
			}
		}
	} else {
		for i := 0; i+1 < n; i++ {
			row_sum -= row[i]
			x := sum - 2*row_sum
			if grid[i][0] == x {
				return true
			}
		}
		var tmp int
		for i := 1; i+1 < n; i++ {
			tmp += row[i]
			if tmp*2+row[0] == sum {
				return true
			}
		}
	}

	return false
}
