package p3905

var dd = []int{-1, 0, 1, 0, -1}

func colorGrid(n int, m int, sources [][]int) [][]int {
	res := make([][]int, n)
	marked := make([][]int, n)
	for i := range n {
		res[i] = make([]int, m)
		marked[i] = make([]int, m)
	}
	var que [][]int
	for _, cur := range sources {
		r, c, w := cur[0], cur[1], cur[2]
		res[r][c] = w
		que = append(que, []int{r, c})
		marked[r][c] = 2
	}
	for len(que) > 0 {
		var next [][]int
		for _, cur := range que {
			r, c := cur[0], cur[1]
			w := res[r][c]
			for i := range 4 {
				nr, nc := r+dd[i], c+dd[i+1]
				if nr >= 0 && nr < n && nc >= 0 && nc < m && marked[nr][nc] < 2 {
					res[nr][nc] = max(res[nr][nc], w)
					if marked[nr][nc] == 0 {
						marked[nr][nc] = 1
						next = append(next, []int{nr, nc})
					}
				}
			}
		}

		for _, cur := range next {
			r, c := cur[0], cur[1]
			marked[r][c] = 2
		}

		que = next
	}
	return res
}
