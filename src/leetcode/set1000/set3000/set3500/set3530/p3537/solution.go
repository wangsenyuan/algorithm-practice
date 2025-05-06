package p3537

func specialGrid(n int) [][]int {
	N := 1 << n
	grid := make([][]int, N)
	for i := range N {
		grid[i] = make([]int, N)
	}

	var dfs func(x0 int, y0 int, x1 int, y1 int, lo int, hi int)

	dfs = func(x0 int, y0 int, x1 int, y1 int, lo int, hi int) {
		if x0+1 == x1 && y0+1 == y1 {
			// lo +1 == hi
			grid[x0][y0] = lo
			return
		}
		x2 := (x0 + x1) / 2
		y2 := (y0 + y1) / 2
		cnt := (hi - lo) / 4
		dfs(x0, y2, x2, y1, lo, lo+cnt)
		dfs(x2, y2, x1, y1, lo+cnt, lo+2*cnt)
		dfs(x2, y0, x1, y2, lo+2*cnt, lo+3*cnt)
		dfs(x0, y0, x2, y2, lo+3*cnt, hi)
	}

	dfs(0, 0, N, N, 0, 1<<(2*n))

	return grid
}
