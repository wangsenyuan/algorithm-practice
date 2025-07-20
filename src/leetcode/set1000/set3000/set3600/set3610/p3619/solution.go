package p3619

func countIslands(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	marked := make([][]bool, n)
	for i := range n {
		marked[i] = make([]bool, m)
	}

	que := make([]int, n*m)

	var dd = []int{-1, 0, 1, 0, -1}

	visit := func(i int, j int) int {
		var head, tail int
		que[head] = i*m + j
		head++
		marked[i][j] = true
		sum := grid[i][j]
		for tail < head {
			cur := que[tail]
			tail++
			r, c := cur/m, cur%m
			for i := range 4 {
				x, y := r+dd[i], c+dd[i+1]
				if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] > 0 && !marked[x][y] {
					marked[x][y] = true
					sum += grid[x][y]
					que[head] = x*m + y
					head++
				}
			}
		}
		return sum
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 || marked[i][j] {
				continue
			}
			tmp := visit(i, j)
			if tmp%k == 0 {
				res++
			}
		}
	}
	return res
}
