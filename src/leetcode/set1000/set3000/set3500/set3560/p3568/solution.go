package p3568

func minMoves(classroom []string, energy int) int {
	var sr, sc int
	var pos [][]int
	n := len(classroom)
	m := len(classroom[0])

	id := make([][]int, n)

	for r, row := range classroom {
		id[r] = make([]int, m)
		for c, ch := range []byte(row) {
			if ch == 'S' {
				sr, sc = r, c
			} else if ch == 'L' {
				id[r][c] = len(pos)
				pos = append(pos, []int{r, c})
			}
		}
	}
	k := len(pos)
	K := 1 << k

	dp := make([][][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([][]int, energy+1)
			for e := range energy + 1 {
				dp[i][j][e] = make([]int, K)
				for x := range K {
					dp[i][j][e][x] = -1
				}
			}
		}
	}

	type data struct {
		r int
		c int
		e int
		s int
	}

	dp[sr][sc][energy][0] = 0

	var que []data
	que = append(que, data{sr, sc, energy, 0})
	var tail int

	var dd = []int{-1, 0, 1, 0, -1}

	for tail < len(que) {
		cur := que[tail]
		tail++
		if cur.s == K-1 {
			return dp[cur.r][cur.c][cur.e][cur.s]
		}

		if cur.e == 0 {
			// not able to
			continue
		}

		for i := 0; i < 4; i++ {
			nr, nc := cur.r+dd[i], cur.c+dd[i+1]
			if nr < 0 || nr >= n || nc < 0 || nc >= m || classroom[nr][nc] == 'X' {
				continue
			}
			ne := cur.e - 1
			if classroom[nr][nc] == 'R' {
				ne = energy
			}
			ns := cur.s
			if classroom[nr][nc] == 'L' && (cur.s>>id[nr][nc])&1 == 0 {
				ns |= 1 << id[nr][nc]
			}
			if dp[nr][nc][ne][ns] == -1 {
				dp[nr][nc][ne][ns] = dp[cur.r][cur.c][cur.e][cur.s] + 1
				que = append(que, data{nr, nc, ne, ns})
			}
		}
	}

	return -1
}
