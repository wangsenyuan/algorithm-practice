package p3552

func minMoves(matrix []string) int {
	// A, B, C.. 这些要一起处理
	n := len(matrix)
	m := len(matrix[0])
	pos := make([][]int, 26)
	dist := make([][]int, n)
	for i, row := range matrix {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = -1
			if row[j] >= 'A' && row[j] <= 'Z' {
				x := int(row[j] - 'A')
				pos[x] = append(pos[x], i*m+j)
			}
		}
	}
	que := make([]int, n*m)
	var head, tail int

	if matrix[0][0] >= 'A' && matrix[0][0] <= 'Z' {
		x := int(matrix[0][0] - 'A')
		for _, cur := range pos[x] {
			i, j := cur/m, cur%m
			dist[i][j] = 0
			que[head] = cur
			head++
		}
	} else {
		dist[0][0] = 0
		que[head] = 0
		head++
	}

	var dd = []int{-1, 0, 1, 0, -1}

	for tail < head {
		cur := que[tail]
		tail++
		r, c := cur/m, cur%m
		if r == n-1 && c == m-1 {
			break
		}
		for i := 0; i < 4; i++ {
			u, v := r+dd[i], c+dd[i+1]
			if u >= 0 && u < n && v >= 0 && v < m && dist[u][v] < 0 && matrix[u][v] != '#' {
				if matrix[u][v] >= 'A' && matrix[u][v] <= 'Z' {
					x := int(matrix[u][v] - 'A')
					for _, w := range pos[x] {
						p, q := w/m, w%m
						que[head] = w
						head++
						dist[p][q] = dist[r][c] + 1
					}
				} else {
					que[head] = u*m + v
					head++
					dist[u][v] = dist[r][c] + 1
				}
			}
		}
	}

	return dist[n-1][m-1]
}
