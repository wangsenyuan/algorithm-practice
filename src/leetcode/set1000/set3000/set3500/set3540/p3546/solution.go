package p3546

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
	return false
}
